# Export all Azure Resource Provider Operations to JSON
# This script uses the Azure PowerShell module to export all available
# Azure resource provider operations/permissions for use with Pike

# Requires: Az.Resources module
# Install with: Install-Module -Name Az.Resources -Scope CurrentUser

param(
    [string]$OutputFile = "azure_permissions_full.json"
)

Write-Host "Checking for Az.Resources module..." -ForegroundColor Cyan

# Check if Az.Resources module is installed
if (-not (Get-Module -ListAvailable -Name Az.Resources)) {
    Write-Host "ERROR: Az.Resources module not found!" -ForegroundColor Red
    Write-Host "Please install it with: Install-Module -Name Az.Resources -Scope CurrentUser" -ForegroundColor Yellow
    exit 1
}

# Import the module
Import-Module Az.Resources -ErrorAction Stop

Write-Host "Fetching all Azure provider operations..." -ForegroundColor Cyan
Write-Host "This may take a minute..." -ForegroundColor Yellow

try {
    # Get all provider operations
    $operations = Get-AzProviderOperation *

    if ($null -eq $operations -or $operations.Count -eq 0) {
        Write-Host "ERROR: No operations retrieved!" -ForegroundColor Red
        Write-Host "This usually means you're not logged into Azure." -ForegroundColor Yellow
        Write-Host "Please run: Connect-AzAccount" -ForegroundColor Yellow
        exit 1
    }

    Write-Host "Found $($operations.Count) operations" -ForegroundColor Green

    # Convert to JSON and save
    $json = $operations | ConvertTo-Json -Depth 10
    $json | Out-File -FilePath $OutputFile -Encoding UTF8

    Write-Host "Successfully exported to: $OutputFile" -ForegroundColor Green

    # Show some statistics
    $providers = $operations | Select-Object -ExpandProperty Operation | ForEach-Object {
        $_.Split('/')[0]
    } | Select-Object -Unique | Sort-Object

    Write-Host "`nFound $($providers.Count) unique resource providers:" -ForegroundColor Cyan
    $providers | ForEach-Object { Write-Host "  - $_" }

} catch {
    Write-Host "ERROR: Failed to fetch operations" -ForegroundColor Red
    Write-Host $_.Exception.Message -ForegroundColor Red
    exit 1
}

Write-Host "`nNext steps:" -ForegroundColor Cyan
Write-Host "1. Run: python3 terraform/azurerm/parse_azure_permissions.py" -ForegroundColor Yellow
Write-Host "2. This will create azure_permissions.json in the required format" -ForegroundColor Yellow
