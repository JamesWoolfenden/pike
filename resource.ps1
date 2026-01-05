#!/usr/bin/env pwsh
<#
.SYNOPSIS
    Creates resource or datasource mappings for Pike (supports AWS, Azure, GCP)
.PARAMETER resource
    The full resource name (e.g., aws_s3_bucket, azurerm_storage_account, google_storage_bucket)
.PARAMETER type
    Either "resource" (default) or "data" for datasources
.EXAMPLE
    pwsh resource.ps1 -resource aws_s3_bucket
    pwsh resource.ps1 -resource azurerm_storage_account -type data
    pwsh resource.ps1 -resource google_compute_instance
#>
param([Parameter(Mandatory)][string]$resource, $type = "resource")

function path()
{
    if ($Env:OS -eq "Windows_NT")
    {
        $slash = "\"
    }
    else
    {
        $slash = "/"
    }

    $path = $args[0]
    for ($i = 1; $i -lt $args.count; $i++) {
        $path += $slash + $args[$i]
    }
    return $path
}

# Enhancement 3: Resource Name Validation
function Validate-Resource($resourceName, $provider, $type)
{
    $membersFile = path $PSScriptRoot "src" "parse" "$provider-members.json"

    if (-not (Test-Path $membersFile))
    {
        Write-Host "Warning: Cannot validate resource - $membersFile not found" -ForegroundColor Yellow
        return $true  # Allow to proceed if validation file doesn't exist
    }

    try
    {
        $members = Get-Content $membersFile -Raw | ConvertFrom-Json
        Write-Debug "Get-Content $membersFile -Raw | ConvertFrom-Json"
        Write-Debug "Type: $type"

        if ($type -eq "resource")
        {
            if ($members.resources -contains $resourceName)
            {
                Write-Host "✓ Resource '$resourceName' validated in $provider provider" -ForegroundColor Green
                return $true
            }
            else
            {
                Write-Host "✗ ERROR: Resource '$resourceName' not found in $provider provider" -ForegroundColor Red

                # Try to find similar resources
                $similar = $members.resources | Where-Object { $_ -like "*$( $resourceName.Split('_')[-1] )*" } | Select-Object -First 5

                if ($similar)
                {
                    Write-Host "`nDid you mean one of these?" -ForegroundColor Yellow
                    $similar | ForEach-Object { Write-Host "  - $_" -ForegroundColor Cyan }
                }

                return $false
            }
        }
        else
        {
            if ($members.dataSources -contains $resourceName)
            {
                Write-Host "✓ Datasource '$resourceName' validated in $provider provider" -ForegroundColor Green
                return $true
            }
            else
            {
                Write-Host "✗ ERROR: Datasource '$resourceName' not found in $provider provider" -ForegroundColor Red

                # Try to find similar resources
                $similar = $members.dataSources | Where-Object { $_ -like "*$( $resourceName.Split('_')[-1] )*" } | Select-Object -First 5

                if ($similar)
                {
                    Write-Host "`nDid you mean one of these?" -ForegroundColor Yellow
                    $similar | ForEach-Object { Write-Host "  - $_" -ForegroundColor Cyan }
                }

                return $false
            }
        }
    }
    catch
    {
        Write-Host "Warning: Error reading members file: $_" -ForegroundColor Yellow
        return $true
    }
}

# Enhancement 5: Interactive Permission Selection
function Get-AzurePermissions($resourceName, $provider, $isDataSource)
{
    if ($provider -ne "azurerm")
    {
        return $null  # Only works for Azure
    }

    $permissionsFile = path $PSScriptRoot "terraform" "azurerm" "azure_permissions.json"

    if (-not (Test-Path $permissionsFile))
    {
        Write-Host "Info: azure_permissions.json not found - skipping permission lookup" -ForegroundColor Cyan
        Write-Host "      Run export_azure_permissions.ps1 to enable this feature" -ForegroundColor Cyan
        return $null
    }

    try
    {
        $allPermissions = Get-Content $permissionsFile -Raw | ConvertFrom-Json -AsHashtable

        # Extract resource type from terraform resource name
        # Example: azurerm_storage_account -> storageAccounts
        # This is fuzzy matching - we'll present options
        $resourcePart = $resourceName -replace "^azurerm_", ""

        # Find potential matches in Azure providers
        $matches = @()

        foreach ($providerName in $allPermissions.Keys)
        {
            $provider = $allPermissions[$providerName]

            foreach ($resourceType in $provider.Keys)
            {
                $resourceTypeLower = $resourceType.ToLower()
                $resourcePartLower = $resourcePart.ToLower() -replace "_", ""

                # Check if resource type matches (fuzzy)
                if ($resourceTypeLower -like "*$resourcePartLower*" -or
                        $resourcePartLower -like "*$resourceTypeLower*")
                {

                    $perms = $provider[$resourceType]
                    $matches += @{
                        Provider = $providerName
                        ResourceType = $resourceType
                        Permissions = $perms
                        FullPath = "$providerName/$resourceType"
                    }
                }
            }
        }

        if ($matches.Count -eq 0)
        {
            Write-Host "No Azure permissions found for '$resourceName'" -ForegroundColor Yellow
            return $null
        }

        if ($matches.Count -eq 1)
        {
            Write-Host "Found permissions for: $( $matches[0].FullPath )" -ForegroundColor Green
            return $matches[0].Permissions
        }

        # Multiple matches - let user choose
        Write-Host "`nFound multiple possible Azure resource types:" -ForegroundColor Cyan
        for ($i = 0; $i -lt $matches.Count; $i++) {
            $m = $matches[$i]
            $readCount = $m.Permissions.read.Count
            $writeCount = $m.Permissions.write.Count
            $deleteCount = $m.Permissions.delete.Count
            $actionCount = $m.Permissions.action.Count

            Write-Host "  [$( $i + 1 )] $( $m.FullPath )" -ForegroundColor Yellow
            Write-Host "      Read: $readCount | Write: $writeCount | Delete: $deleteCount | Action: $actionCount" -ForegroundColor Gray
        }

        do
        {
            $choice = Read-Host "`nWhich resource type should we use? [1-$( $matches.Count ), or 0 to skip]"
            $choiceNum = [int]$choice
        } while ($choiceNum -lt 0 -or $choiceNum -gt $matches.Count)

        if ($choiceNum -eq 0)
        {
            Write-Host "Skipping permission lookup" -ForegroundColor Yellow
            return $null
        }

        $selected = $matches[$choiceNum - 1]
        Write-Host "✓ Selected: $( $selected.FullPath )" -ForegroundColor Green

        return $selected.Permissions

    }
    catch
    {
        Write-Host "Warning: Error reading permissions file: $_" -ForegroundColor Yellow
        return $null
    }
}

# Main script logic
$provider = $resource.Split("_")[0]

# Validate provider
$supportedProviders = @("aws", "azurerm", "google")
if ($provider -notin $supportedProviders)
{
    Write-Host "✗ ERROR: Unsupported provider '$provider'" -ForegroundColor Red
    Write-Host "`nSupported providers:" -ForegroundColor Yellow
    Write-Host "  - aws (e.g., aws_s3_bucket)" -ForegroundColor Cyan
    Write-Host "  - azurerm (e.g., azurerm_storage_account)" -ForegroundColor Cyan
    Write-Host "  - google (e.g., google_storage_bucket)" -ForegroundColor Cyan
    exit 1
}

Write-Host "`nPike Resource Creator" -ForegroundColor Cyan
Write-Host ("-" * 60) -ForegroundColor Gray
Write-Host "Provider: $provider" -ForegroundColor Gray
Write-Host "Resource: $resource" -ForegroundColor Gray
Write-Host "Type: $type" -ForegroundColor Gray
Write-Host ("-" * 60) -ForegroundColor Gray

# Validate resource name
Write-Host "`nValidating resource..." -ForegroundColor Cyan
if (-not (Validate-Resource $resource $provider $type))
{
    Write-Host "`nAborting: Resource validation failed" -ForegroundColor Red
    exit 1
}

# Look up Azure permissions (Azure only)
if ($provider -eq "azurerm")
{
    Write-Host "`nLooking up Azure permissions..." -ForegroundColor Cyan
    $permissions = Get-AzurePermissions $resource $provider ($type -eq "data")
}
else
{
    Write-Host "`nSkipping permission lookup (only available for Azure)" -ForegroundColor Gray
    $permissions = $null
}

if ($permissions)
{
    Write-Host "`nPermission Summary:" -ForegroundColor Cyan
    Write-Host "  Read permissions:   $( $permissions.read.Count )" -ForegroundColor Gray
    Write-Host "  Write permissions:  $( $permissions.write.Count )" -ForegroundColor Gray
    Write-Host "  Delete permissions: $( $permissions.delete.Count )" -ForegroundColor Gray
    Write-Host "  Action permissions: $( $permissions.action.Count )" -ForegroundColor Gray
}

$baseMapping = path $PSScriptRoot "src" "mapping" $provider
$mapping = path $baseMapping $type

Write-Host "`nCreating mapping files..." -ForegroundColor Cyan

if (test-path($baseMapping))
{
    $source = path $mapping "template.json"
    $destination = path $mapping "$resource.json"

    if (Test-Path $source)
    {
        Copy-Item -path $source -destination $destination
        Write-Host "✓ Copied template to: $destination" -ForegroundColor Green
    }
    else
    {
        Write-Host "⚠ Template not found: $source" -ForegroundColor Yellow
        Write-Host "  Skipping JSON creation - you'll need to create it manually" -ForegroundColor Yellow
    }
}
else
{
    Write-Host "✗ Mapping directory not found: $baseMapping" -ForegroundColor Red
}

# Create Terraform test file
if ($type -eq "data")
{
    $content = "data `"$resource`" `"pike`" {`n}`n`n"
    $output = "output `"$resource`" {`n  value = data.$resource.pike`n}"
    $content = $content + $output
    $tffile = path terraform $provider "$type.$resource.tf"
}
else
{
    $content = "resource `"$resource`" `"pike`" {}"
    $tffile = path terraform $provider "$resource.tf"
}

$tffile = path $PSScriptRoot $tffile
new-item $tffile -value $content -Force | Out-Null
Write-Host "✓ Created Terraform file: $tffile" -ForegroundColor Green

Write-Host "`n✓ Done!" -ForegroundColor Green

Write-Host "`nNext steps:" -ForegroundColor Cyan
if ($permissions)
{
    Write-Host "  1. Edit the JSON mapping file to add the permissions shown above" -ForegroundColor Yellow
    Write-Host "  2. Run: pwsh release.ps1 -Provider $provider  (to update Go embed files)" -ForegroundColor Yellow
}
else
{
    if (Test-Path $destination)
    {
        Write-Host "  1. Edit the JSON mapping file: $destination" -ForegroundColor Yellow
    }
    else
    {
        Write-Host "  1. Create the JSON mapping file with appropriate permissions" -ForegroundColor Yellow
    }
    Write-Host "  2. Run: pwsh release.ps1 -Provider $provider  (to update Go embed files)" -ForegroundColor Yellow
}
Write-Host "  3. Test with: go build" -ForegroundColor Yellow
Write-Host "  4. Test with: go test ./..." -ForegroundColor Yellow
