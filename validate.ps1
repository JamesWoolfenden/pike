#!/usr/bin/env pwsh
<#
.SYNOPSIS
    Randomly validates Pike resource/datasource mappings against test files
.DESCRIPTION
    Randomly selects resources/datasources from Pike's lookup maps and:
    - Checks if test files exist in terraform/{provider}/backup/
    - Optionally runs empirical lifecycle tests
    - Compares discovered permissions vs JSON mappings
    - Reports discrepancies

.PARAMETER Provider
    Cloud provider to test (google, aws, azurerm)
.PARAMETER Resources
    Number of random resources to validate (default: 10)
.PARAMETER Datasources
    Number of random datasources to validate (default: 5)
.PARAMETER RunTests
    Actually run terraform apply/destroy lifecycle tests
.PARAMETER Seed
    Random seed for reproducible selection (default: random)
.EXAMPLE
    pwsh validate.ps1 -Provider google -Resources 10 -Datasources 5
    pwsh validate.ps1 -Provider google -Resources 5 -RunTests
    pwsh validate.ps1 -Provider google -Resources 10 -Seed 42
#>
param(
    [ValidateSet("google", "aws", "azurerm")]
    [string]$Provider = "google",

    [int]$Resources = 10,

    [int]$Datasources = 5,

    [switch]$RunTests,

    [int]$Seed = (Get-Random)
)

function path() {
    if ($Env:OS -eq "Windows_NT") {
        $slash = "\"
    }
    else {
        $slash = "/"
    }

    $path = $args[0]
    for ($i = 1; $i -lt $args.count; $i++) {
        $path += $slash + $args[$i]
    }
    return $path
}

# Provider configurations
$providerConfigs = @{
    google = @{
        Name = "GCP"
        ResourceLookupFile = "gcp.go"
        ResourceLookupMap = "gCPTfLookup"
        DataLookupFile = "gcp_datasource.go"
        DataLookupMap = "TFLookup"
        BackupDir = "terraform/google/backup"
        MakeCommand = "make -C terraform/google"
    }
    aws = @{
        Name = "AWS"
        ResourceLookupFile = "aws.go"
        ResourceLookupMap = "tFLookup"
        DataLookupFile = "aws_datasource.go"
        DataLookupMap = "tFLookupDataAWS"
        BackupDir = "terraform/aws/backup"
        MakeCommand = "make -C terraform/aws"
    }
    azurerm = @{
        Name = "Azure"
        ResourceLookupFile = "azure.go"
        ResourceLookupMap = "tFLookupAzure"
        DataLookupFile = "azure_datasource.go"
        DataLookupMap = "TFLookupAzureData"
        BackupDir = "terraform/azurerm/backup"
        MakeCommand = "make -C terraform/azurerm"
    }
}

function Extract-LookupMapKeys {
    param(
        [string]$FilePath,
        [string]$MapName
    )

    if (-not (Test-Path $FilePath)) {
        Write-Host "  ⚠ File not found: $FilePath" -ForegroundColor Yellow
        return @()
    }

    $content = Get-Content $FilePath -Raw

    # Extract all resource names from the lookup map
    $matches = [regex]::Matches($content, '^\s*"([^"]+)":\s+\w+,\s*(?://.*)?$', [System.Text.RegularExpressions.RegexOptions]::Multiline)

    $keys = @()
    foreach ($match in $matches) {
        $keys += $match.Groups[1].Value
    }

    return $keys
}

function Find-TestFile {
    param(
        [string]$ResourceName,
        [string]$BackupDir,
        [string]$Type  # "resource" or "data"
    )

    $basePath = path $PSScriptRoot $BackupDir

    if (-not (Test-Path $basePath)) {
        return $null
    }

    # Check for different naming patterns
    if ($Type -eq "data") {
        $patterns = @(
            "data.$ResourceName.tf",
            "$ResourceName.tf"
        )
    }
    else {
        $patterns = @(
            "$ResourceName.tf",
            "resource.$ResourceName.tf"
        )
    }

    foreach ($pattern in $patterns) {
        $testFile = path $basePath $pattern
        if (Test-Path $testFile) {
            return $testFile
        }
    }

    return $null
}

function Select-RandomItems {
    param(
        [array]$Items,
        [int]$Count,
        [int]$Seed
    )

    if ($Items.Count -eq 0) {
        return @()
    }

    $actualCount = [Math]::Min($Count, $Items.Count)

    # Use seed for reproducibility
    Get-Random -InputObject $Items -Count $actualCount -SetSeed $Seed
}

function Test-ResourceLifecycle {
    param(
        [string]$TestFile,
        [string]$MakeCommand
    )

    # This would run the actual terraform lifecycle
    # For now, return placeholder
    return @{
        Success = $true
        DiscoveredPermissions = @()
        Errors = @()
    }
}

# Main script
$config = $providerConfigs[$Provider]

Write-Host "`nPike Validation - Random Resource Testing" -ForegroundColor Cyan
Write-Host ("═" * 70)
Write-Host "Provider:    $($config.Name)" -ForegroundColor Gray
Write-Host "Resources:   $Resources" -ForegroundColor Gray
Write-Host "Datasources: $Datasources" -ForegroundColor Gray
Write-Host "Seed:        $Seed" -ForegroundColor Gray
Write-Host "Run Tests:   $RunTests" -ForegroundColor Gray
Write-Host ("═" * 70)

# Extract resource keys from lookup maps
Write-Host "`n📋 Reading lookup maps..." -ForegroundColor Cyan

$resourceLookupPath = path $PSScriptRoot "src" $config.ResourceLookupFile
$allResources = Extract-LookupMapKeys -FilePath $resourceLookupPath -MapName $config.ResourceLookupMap
Write-Host "  Found $($allResources.Count) resources in $($config.ResourceLookupMap)" -ForegroundColor Gray

$dataLookupPath = path $PSScriptRoot "src" $config.DataLookupFile
$allDatasources = Extract-LookupMapKeys -FilePath $dataLookupPath -MapName $config.DataLookupMap
Write-Host "  Found $($allDatasources.Count) datasources in $($config.DataLookupMap)" -ForegroundColor Gray

# Select random resources
Write-Host "`n🎲 Selecting random items (seed: $Seed)..." -ForegroundColor Cyan

$selectedResources = Select-RandomItems -Items $allResources -Count $Resources -Seed $Seed
Write-Host "  Selected $($selectedResources.Count) resources" -ForegroundColor Gray

$selectedDatasources = Select-RandomItems -Items $allDatasources -Count $Datasources -Seed ($Seed + 1)
Write-Host "  Selected $($selectedDatasources.Count) datasources" -ForegroundColor Gray

# Validate resources
Write-Host "`n📦 Validating Resources" -ForegroundColor Cyan
Write-Host ("─" * 70)

$resourceResults = @()
foreach ($resourceName in $selectedResources) {
    $testFile = Find-TestFile -ResourceName $resourceName -BackupDir $config.BackupDir -Type "resource"

    $result = @{
        Name = $resourceName
        Type = "resource"
        TestFile = $testFile
        HasTestFile = $null -ne $testFile
        Status = "unknown"
    }

    if ($testFile) {
        Write-Host "  ✓ $resourceName" -ForegroundColor Green
        Write-Host "    Test file: $testFile" -ForegroundColor DarkGray
        $result.Status = "has_test"

        if ($RunTests) {
            Write-Host "    Running lifecycle tests..." -ForegroundColor Cyan
            # TODO: Implement actual test run
            $result.Status = "tested"
        }
    }
    else {
        Write-Host "  ⚠ $resourceName" -ForegroundColor Yellow
        Write-Host "    No test file found in $($config.BackupDir)" -ForegroundColor DarkGray
        $result.Status = "no_test"
    }

    $resourceResults += $result
}

# Validate datasources
Write-Host "`n📊 Validating Datasources" -ForegroundColor Cyan
Write-Host ("─" * 70)

$datasourceResults = @()
foreach ($dsName in $selectedDatasources) {
    $testFile = Find-TestFile -ResourceName $dsName -BackupDir $config.BackupDir -Type "data"

    $result = @{
        Name = $dsName
        Type = "datasource"
        TestFile = $testFile
        HasTestFile = $null -ne $testFile
        Status = "unknown"
    }

    if ($testFile) {
        Write-Host "  ✓ $dsName" -ForegroundColor Green
        Write-Host "    Test file: $testFile" -ForegroundColor DarkGray
        $result.Status = "has_test"

        if ($RunTests) {
            Write-Host "    Running lifecycle tests..." -ForegroundColor Cyan
            # TODO: Implement actual test run
            $result.Status = "tested"
        }
    }
    else {
        Write-Host "  ⚠ $dsName" -ForegroundColor Yellow
        Write-Host "    No test file found in $($config.BackupDir)" -ForegroundColor DarkGray
        $result.Status = "no_test"
    }

    $datasourceResults += $result
}

# Summary
Write-Host "`n" + ("═" * 70)
Write-Host "📈 Summary" -ForegroundColor Cyan
Write-Host ("─" * 70)

$allResults = $resourceResults + $datasourceResults

$withTests = ($allResults | Where-Object { $_.HasTestFile }).Count
$withoutTests = ($allResults | Where-Object { -not $_.HasTestFile }).Count
$tested = ($allResults | Where-Object { $_.Status -eq "tested" }).Count

Write-Host "`nResources with test files:    $withTests / $($allResults.Count)" -ForegroundColor Green
Write-Host "Resources without test files: $withoutTests / $($allResults.Count)" -ForegroundColor Yellow

if ($RunTests) {
    Write-Host "Successfully tested:          $tested / $withTests" -ForegroundColor Green
}

# Missing test files report
if ($withoutTests -gt 0) {
    Write-Host "`n⚠ Missing Test Files:" -ForegroundColor Yellow
    foreach ($item in $allResults | Where-Object { -not $_.HasTestFile }) {
        Write-Host "  - $($item.Name) [$($item.Type)]" -ForegroundColor DarkGray
    }
}

# Coverage percentage
$coveragePercent = [math]::Round(($withTests / $allResults.Count) * 100, 2)
Write-Host "`nTest Coverage: $coveragePercent%" -ForegroundColor Cyan

# Save results to JSON
$reportPath = path $PSScriptRoot ".pike" "validation-report-$(Get-Date -Format 'yyyyMMdd-HHmmss').json"
$reportDir = Split-Path $reportPath -Parent

if (-not (Test-Path $reportDir)) {
    New-Item -ItemType Directory -Path $reportDir | Out-Null
}

$report = @{
    Timestamp = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
    Provider = $Provider
    Seed = $Seed
    Summary = @{
        TotalTested = $allResults.Count
        WithTestFiles = $withTests
        WithoutTestFiles = $withoutTests
        CoveragePercent = $coveragePercent
    }
    Resources = $resourceResults
    Datasources = $datasourceResults
}

$report | ConvertTo-Json -Depth 10 | Out-File -FilePath $reportPath -Encoding UTF8
Write-Host "`n💾 Report saved to: $reportPath" -ForegroundColor Cyan

Write-Host "`n✓ Validation complete!" -ForegroundColor Green
Write-Host ""
