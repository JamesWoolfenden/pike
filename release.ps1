#!/usr/bin/env pwsh
<#
.SYNOPSIS
    Updates Go embed files and lookup maps for all cloud providers
.DESCRIPTION
    This script scans provider mapping directories and automatically generates:
    - Resource and datasource embed files (files_*.go)
    - Lookup maps in provider-specific .go files

    Supports: AWS, Azure (azurerm), and GCP (google)
#>

param(
    [switch]$DryRun,
    [switch]$Verbose,
    [ValidateSet("all", "aws", "azurerm", "google")]
    [string]$Provider = "all"
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
    aws = @{
        Name = "AWS"
        MappingDir = "aws"
        ResourcePrefix = "aws"
        DataPrefix = "dataAws"
        FilesResourceGo = "files_aws.go"
        FilesDataGo = "files_datasource.go"
        ResourceLookupFile = "aws.go"
        ResourceLookupMap = "tFLookup"
        DataLookupFile = "aws_datasource.go"
        DataLookupMap = "tFLookupDataAWS"
    }
    azurerm = @{
        Name = "Azure"
        MappingDir = "azurerm"
        ResourcePrefix = "azurerm"
        DataPrefix = "dataAzurerm"
        FilesResourceGo = "files_azure.go"
        FilesDataGo = "files_azure_datasource.go"
        ResourceLookupFile = "azure.go"
        ResourceLookupMap = "tFLookupAzure"
        DataLookupFile = "azure_datasource.go"
        DataLookupMap = "TFLookupAzureData"
    }
    google = @{
        Name = "GCP"
        MappingDir = "google"
        ResourcePrefix = "google"
        DataPrefix = "dataGoogle"
        FilesResourceGo = "files_gcp.go"
        FilesDataGo = "files_gcp_datasource.go"
        ResourceLookupFile = "gcp.go"
        ResourceLookupMap = "gCPTfLookup"
        DataLookupFile = "gcp_datasource.go"
        DataLookupMap = "TFLookup"
    }
}

function Convert-ToVariableName {
    param(
        [string]$ResourceName,
        [string]$Prefix
    )

    # Remove provider prefix (aws_, azurerm_, google_)
    $name = $ResourceName -replace '^(aws|azurerm|google)_', ''

    # Split on underscore and convert to PascalCase
    $parts = $name -split '_'
    $pascalCase = ($parts | ForEach-Object {
        if ($_.Length -gt 0) {
            $_.Substring(0,1).ToUpper() + $_.Substring(1).ToLower()
        }
    }) -join ''

    # Handle special cases (common across all providers)
    $pascalCase = $pascalCase -replace 'Dns', 'DNS'
    $pascalCase = $pascalCase -replace 'Ip', 'IP'
    $pascalCase = $pascalCase -replace 'Api', 'API'
    $pascalCase = $pascalCase -replace 'Sql', 'SQL'
    $pascalCase = $pascalCase -replace 'Ssh', 'SSH'
    $pascalCase = $pascalCase -replace 'Hci', 'HCI'
    $pascalCase = $pascalCase -replace 'Vpc', 'VPC'
    $pascalCase = $pascalCase -replace 'Acm', 'ACM'
    $pascalCase = $pascalCase -replace 'Iam', 'IAM'
    $pascalCase = $pascalCase -replace 'Kms', 'KMS'
    $pascalCase = $pascalCase -replace 'Ami', 'AMI'
    $pascalCase = $pascalCase -replace 'Id', 'ID'
    $pascalCase = $pascalCase -replace 'Ids', 'IDs'
    $pascalCase = $pascalCase -replace 'V([0-9]+)', 'V$1'  # Preserve version numbers

    return "$Prefix$pascalCase"
}

function Get-MappingFiles {
    param(
        [string]$BasePath,
        [string]$MappingDir,
        [string]$Type,  # "resource" or "data"
        [string]$Prefix
    )

    $mappingPath = path $BasePath "src" "mapping" $MappingDir $Type

    if (-not (Test-Path $mappingPath)) {
        Write-Host "  Warning: Mapping path not found: $mappingPath" -ForegroundColor Yellow
        return @()
    }

    # Find all JSON files except template.json
    $files = Get-ChildItem -Path $mappingPath -Filter "*.json" -Recurse |
             Where-Object { $_.Name -ne "template.json" } |
             Sort-Object Name

    $mappings = @()

    foreach ($file in $files) {
        $relativePath = $file.FullName.Replace("$BasePath\", "").Replace("\", "/")
        $relativePath = $relativePath -replace '^src/', ''

        # Extract resource name from filename
        $resourceName = $file.BaseName

        # Get subdirectory for organization
        $subdir = Split-Path -Parent ($file.FullName.Replace($mappingPath, "").TrimStart('\').TrimStart('/'))

        $mappings += @{
            ResourceName = $resourceName
            VariableName = Convert-ToVariableName -ResourceName $resourceName -Prefix $Prefix
            FilePath = $relativePath
            SubDirectory = $subdir
        }
    }

    return $mappings
}

function Generate-FilesGo {
    param(
        [array]$Mappings,
        [string]$OutputPath
    )

    $content = @"
package pike

import (
	_ "embed" // required for embed
)

"@

    foreach ($mapping in $Mappings | Sort-Object ResourceName) {
        $content += @"
//go:embed $($mapping.FilePath)
var $($mapping.VariableName) []byte

"@
    }

    return $content
}

function Read-ExistingLookupMap {
    param(
        [string]$FilePath,
        [string]$MapName
    )

    if (-not (Test-Path $FilePath)) {
        return @{}
    }

    $content = Get-Content $FilePath -Raw

    # Extract the lookup map using regex - handle both var and inline declarations
    $pattern1 = "(?s)var $MapName = map\[string\]interface\{\}\{(.*?)\n\}"
    $pattern2 = "(?s)$MapName := map\[string\]interface\{\}\{(.*?)\n\t\}"

    $mapContent = $null
    if ($content -match $pattern1) {
        $mapContent = $Matches[1]
    }
    elseif ($content -match $pattern2) {
        $mapContent = $Matches[1]
    }

    if ($mapContent) {
        $entries = @{}
        $lines = $mapContent -split '\n' | Where-Object { $_ -match '^\s*"([^"]+)":\s*(\w+),' }

        foreach ($line in $lines) {
            if ($line -match '^\s*"([^"]+)":\s*(\w+),') {
                $entries[$Matches[1]] = $Matches[2]
            }
        }

        return $entries
    }

    return @{}
}

function Show-LookupMapUpdates {
    param(
        [hashtable]$Existing,
        [array]$Mappings,
        [string]$Type,
        [string]$ProviderName
    )

    $newEntries = @()
    $existingEntries = @()

    foreach ($mapping in $Mappings) {
        if (-not $Existing.ContainsKey($mapping.ResourceName)) {
            $newEntries += $mapping.ResourceName
        }
        else {
            $existingEntries += $mapping.ResourceName
        }
    }

    Write-Host "  $ProviderName $Type Lookup Map:" -ForegroundColor Cyan
    Write-Host "    Existing entries: $($existingEntries.Count)" -ForegroundColor Gray

    if ($newEntries.Count -gt 0) {
        Write-Host "    New entries to add: $($newEntries.Count)" -ForegroundColor Green
        if ($Verbose) {
            foreach ($entry in $newEntries | Sort-Object | Select-Object -First 10) {
                Write-Host "      + $entry" -ForegroundColor Green
            }
            if ($newEntries.Count -gt 10) {
                Write-Host "      ... and $($newEntries.Count - 10) more" -ForegroundColor Gray
            }
        }
    }
    else {
        Write-Host "    No new entries" -ForegroundColor Gray
    }

    # Check for entries in existing that are no longer in mappings
    $removedEntries = $Existing.Keys | Where-Object { $_ -notin ($Mappings | ForEach-Object { $_.ResourceName }) }
    if ($removedEntries.Count -gt 0) {
        Write-Host "    Entries no longer in mappings: $($removedEntries.Count)" -ForegroundColor Yellow
        if ($Verbose) {
            foreach ($entry in $removedEntries | Sort-Object | Select-Object -First 10) {
                Write-Host "      - $entry" -ForegroundColor Yellow
            }
            if ($removedEntries.Count -gt 10) {
                Write-Host "      ... and $($removedEntries.Count - 10) more" -ForegroundColor Gray
            }
        }
    }
}

function Update-ProviderFiles {
    param(
        [hashtable]$Config,
        [string]$BasePath,
        [bool]$IsDryRun
    )

    $providerName = $Config.Name
    Write-Host "`n$providerName Provider" -ForegroundColor Cyan
    Write-Host ("─" * 60) -ForegroundColor Gray

    # Get all resource mappings
    Write-Host "  Scanning resource mappings..." -ForegroundColor Gray
    $resourceMappings = Get-MappingFiles -BasePath $BasePath `
                                         -MappingDir $Config.MappingDir `
                                         -Type "resource" `
                                         -Prefix $Config.ResourcePrefix
    Write-Host "    Found $($resourceMappings.Count) resource mappings" -ForegroundColor Gray

    # Get all datasource mappings
    Write-Host "  Scanning datasource mappings..." -ForegroundColor Gray
    $datasourceMappings = Get-MappingFiles -BasePath $BasePath `
                                           -MappingDir $Config.MappingDir `
                                           -Type "data" `
                                           -Prefix $Config.DataPrefix
    Write-Host "    Found $($datasourceMappings.Count) datasource mappings" -ForegroundColor Gray

    # Check existing lookup maps
    $existingResourceLookup = Read-ExistingLookupMap `
        -FilePath (path $BasePath "src" $Config.ResourceLookupFile) `
        -MapName $Config.ResourceLookupMap

    $existingDataLookup = Read-ExistingLookupMap `
        -FilePath (path $BasePath "src" $Config.DataLookupFile) `
        -MapName $Config.DataLookupMap

    # Show what will be updated
    Write-Host ""
    Show-LookupMapUpdates -Existing $existingResourceLookup `
                          -Mappings $resourceMappings `
                          -Type "Resource" `
                          -ProviderName $providerName

    Show-LookupMapUpdates -Existing $existingDataLookup `
                          -Mappings $datasourceMappings `
                          -Type "Datasource" `
                          -ProviderName $providerName

    if ($IsDryRun) {
        return @{
            ResourceCount = $resourceMappings.Count
            DataCount = $datasourceMappings.Count
            Config = $Config
        }
    }

    # Generate files_*.go for resources
    Write-Host "`n  Generating src/$($Config.FilesResourceGo)..." -ForegroundColor Gray
    $filesResourceContent = Generate-FilesGo -Mappings $resourceMappings
    $filesResourcePath = path $BasePath "src" $Config.FilesResourceGo
    $filesResourceContent | Out-File -FilePath $filesResourcePath -Encoding UTF8 -NoNewline
    Write-Host "    ✓ Generated with $($resourceMappings.Count) embed statements" -ForegroundColor Green

    # Generate files_*.go for datasources
    Write-Host "  Generating src/$($Config.FilesDataGo)..." -ForegroundColor Gray
    $filesDataContent = Generate-FilesGo -Mappings $datasourceMappings
    $filesDataPath = path $BasePath "src" $Config.FilesDataGo
    $filesDataContent | Out-File -FilePath $filesDataPath -Encoding UTF8 -NoNewline
    Write-Host "    ✓ Generated with $($datasourceMappings.Count) embed statements" -ForegroundColor Green

    # Update resource lookup map
    Write-Host "  Updating src/$($Config.ResourceLookupFile) lookup map..." -ForegroundColor Gray
    $resourceGoPath = path $BasePath "src" $Config.ResourceLookupFile
    if (Test-Path $resourceGoPath) {
        $resourceContent = Get-Content $resourceGoPath -Raw

        # Generate new lookup map entries
        $lookupEntries = ($resourceMappings | Sort-Object ResourceName | ForEach-Object {
            "`t`"$($_.ResourceName)`": $($_.VariableName),"
        }) -join "`n"

        $newLookupMap = @"
var $($Config.ResourceLookupMap) = map[string]interface{}{
$lookupEntries
}
"@

        # Replace the existing lookup map
        $pattern = "(?s)var $($Config.ResourceLookupMap) = map\[string\]interface\{\}\{.*?\n\}"
        if ($resourceContent -match $pattern) {
            $resourceContent = $resourceContent -replace $pattern, $newLookupMap
            $resourceContent | Out-File -FilePath $resourceGoPath -Encoding UTF8 -NoNewline
            Write-Host "    ✓ Updated $($Config.ResourceLookupMap) map with $($resourceMappings.Count) entries" -ForegroundColor Green
        }
        else {
            Write-Host "    ⚠ Could not find $($Config.ResourceLookupMap) map in $($Config.ResourceLookupFile)" -ForegroundColor Yellow
        }
    }

    # Update datasource lookup map
    Write-Host "  Updating src/$($Config.DataLookupFile) lookup map..." -ForegroundColor Gray
    $dataGoPath = path $BasePath "src" $Config.DataLookupFile
    if (Test-Path $dataGoPath) {
        $dataContent = Get-Content $dataGoPath -Raw

        # Generate new lookup map entries
        $dataLookupEntries = ($datasourceMappings | Sort-Object ResourceName | ForEach-Object {
            "`t`"$($_.ResourceName)`": $($_.VariableName),"
        }) -join "`n"

        # Handle both var and inline declarations
        if ($Config.DataLookupMap -eq "TFLookupAzureData") {
            $newDataLookupMap = @"
	TFLookupAzureData := map[string]interface{}{
$dataLookupEntries
	}
"@
            $pattern = "(?s)TFLookupAzureData := map\[string\]interface\{\}\{.*?\n\t\}"
        }
        elseif ($Config.DataLookupMap -eq "TFLookup") {
            $newDataLookupMap = @"
	TFLookup := map[string]interface{}{
$dataLookupEntries
	}
"@
            $pattern = "(?s)TFLookup := map\[string\]interface\{\}\{.*?\n\t\}"
        }
        else {
            $newDataLookupMap = @"
var $($Config.DataLookupMap) = map[string]interface{}{
$dataLookupEntries
}
"@
            $pattern = "(?s)var $($Config.DataLookupMap) = map\[string\]interface\{\}\{.*?\n\}"
        }

        # Replace the existing lookup map
        if ($dataContent -match $pattern) {
            $dataContent = $dataContent -replace $pattern, $newDataLookupMap
            $dataContent | Out-File -FilePath $dataGoPath -Encoding UTF8 -NoNewline
            Write-Host "    ✓ Updated $($Config.DataLookupMap) map with $($datasourceMappings.Count) entries" -ForegroundColor Green
        }
        else {
            Write-Host "    ⚠ Could not find $($Config.DataLookupMap) map in $($Config.DataLookupFile)" -ForegroundColor Yellow
        }
    }

    return @{
        ResourceCount = $resourceMappings.Count
        DataCount = $datasourceMappings.Count
        Config = $Config
    }
}

# Main script
$basePath = $PSScriptRoot

Write-Host "`nPike Release - Multi-Provider Mapping Generator" -ForegroundColor Cyan
Write-Host ("═" * 60)

# Determine which providers to process
$providersToProcess = @()
if ($Provider -eq "all") {
    $providersToProcess = @("aws", "azurerm", "google")
}
else {
    $providersToProcess = @($Provider)
}

# Process each provider
$results = @()
foreach ($providerKey in $providersToProcess) {
    $config = $providerConfigs[$providerKey]
    $result = Update-ProviderFiles -Config $config -BasePath $basePath -IsDryRun $DryRun
    $results += $result
}

# Summary
Write-Host "`n" + ("═" * 60)

if ($DryRun) {
    Write-Host "[DRY RUN] Would update the following files:" -ForegroundColor Yellow
    Write-Host ""
    foreach ($result in $results) {
        $cfg = $result.Config
        Write-Host "  $($cfg.Name):" -ForegroundColor Cyan
        Write-Host "    - src/$($cfg.FilesResourceGo) ($($result.ResourceCount) embeds)" -ForegroundColor Gray
        Write-Host "    - src/$($cfg.FilesDataGo) ($($result.DataCount) embeds)" -ForegroundColor Gray
        Write-Host "    - src/$($cfg.ResourceLookupFile) ($($cfg.ResourceLookupMap) map)" -ForegroundColor Gray
        Write-Host "    - src/$($cfg.DataLookupFile) ($($cfg.DataLookupMap) map)" -ForegroundColor Gray
    }
    Write-Host "`nRun without -DryRun to apply changes" -ForegroundColor Yellow
}
else {
    Write-Host "✓ Done! Updated files for:" -ForegroundColor Green
    foreach ($result in $results) {
        Write-Host "  - $($result.Config.Name) ($($result.ResourceCount) resources, $($result.DataCount) datasources)" -ForegroundColor Gray
    }

    Write-Host "`nNext steps:" -ForegroundColor Cyan
    Write-Host "  1. Review the generated files" -ForegroundColor Yellow
    Write-Host "  2. Run: go build" -ForegroundColor Yellow
    Write-Host "  3. Run: go test ./..." -ForegroundColor Yellow
    Write-Host "  4. Commit the changes" -ForegroundColor Yellow
}

Write-Host ""
