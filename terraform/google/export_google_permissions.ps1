#!/usr/bin/env pwsh
# Export all Google Cloud IAM permissions from GCP IAM API to JSON
# This script uses gcloud CLI to query all testable permissions and organizes
# them by service and resource type for use with Pike resource.ps1

param(
    [string]$Project = $env:GCP_PROJECT,
    [string]$OutputFile = "google_permissions.json"
)

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

Write-Host "Pike Google Cloud Permissions Exporter" -ForegroundColor Cyan
Write-Host ("-" * 60) -ForegroundColor Gray

# Check if gcloud is installed
$gcloudPath = Get-Command gcloud -ErrorAction SilentlyContinue
if (-not $gcloudPath)
{
    Write-Host "ERROR: gcloud CLI not found!" -ForegroundColor Red
    Write-Host "Please install it from: https://cloud.google.com/sdk/docs/install" -ForegroundColor Yellow
    exit 1
}

Write-Host "✓ Found gcloud CLI: $($gcloudPath.Source)" -ForegroundColor Green

# Check if project is set
if (-not $Project)
{
    Write-Host "ERROR: GCP_PROJECT environment variable not set!" -ForegroundColor Red
    Write-Host "Please set it with: export GCP_PROJECT=your-project-id" -ForegroundColor Yellow
    Write-Host "Or pass it as parameter: -Project your-project-id" -ForegroundColor Yellow
    exit 1
}

Write-Host "✓ Using GCP Project: $Project" -ForegroundColor Green

# Query all testable permissions
Write-Host "`nQuerying all testable permissions from Google Cloud..." -ForegroundColor Cyan
Write-Host "This may take 30-60 seconds..." -ForegroundColor Yellow

$resourceUri = "//cloudresourcemanager.googleapis.com/projects/$Project"

try
{
    $permissionsJson = gcloud iam list-testable-permissions $resourceUri --format=json 2>&1

    if ($LASTEXITCODE -ne 0)
    {
        Write-Host "ERROR: Failed to query permissions!" -ForegroundColor Red
        Write-Host $permissionsJson -ForegroundColor Red
        Write-Host "`nMake sure you're authenticated with: gcloud auth login" -ForegroundColor Yellow
        exit 1
    }

    $permissions = $permissionsJson | ConvertFrom-Json

    Write-Host "✓ Retrieved $($permissions.Count) permissions" -ForegroundColor Green

}
catch
{
    Write-Host "ERROR: Failed to query permissions: $_" -ForegroundColor Red
    exit 1
}

# Organize permissions by service and resource type
# Structure: { "service.resource": { "permissions": [...], "terraform_resources": [...] } }
Write-Host "`nOrganizing permissions by service and resource type..." -ForegroundColor Cyan

$permissionsDb = @{}

foreach ($perm in $permissions)
{
    $permName = $perm.name

    # Parse permission: "service.resource.action" -> "service.resource"
    $parts = $permName.Split('.')

    if ($parts.Count -ge 2)
    {
        $service = $parts[0]
        $resource = $parts[1]
        $key = "$service.$resource"

        if (-not $permissionsDb.ContainsKey($key))
        {
            $permissionsDb[$key] = @{
                service = $service
                resource = $resource
                permissions = @()
                stage_counts = @{
                    GA = 0
                    BETA = 0
                    ALPHA = 0
                    DEPRECATED = 0
                }
            }
        }

        # Add permission details
        $permissionsDb[$key].permissions += @{
            name = $permName
            stage = $perm.stage
            customRolesSupportLevel = $perm.customRolesSupportLevel
        }

        # Count by stage
        if ($perm.stage)
        {
            $permissionsDb[$key].stage_counts[$perm.stage]++
        }
    }
}

Write-Host "✓ Organized into $($permissionsDb.Count) service.resource groups" -ForegroundColor Green

# Enhance with Terraform resource name suggestions
Write-Host "`nMapping to Terraform resource names..." -ForegroundColor Cyan

foreach ($key in $permissionsDb.Keys)
{
    $entry = $permissionsDb[$key]
    $service = $entry.service
    $resource = $entry.resource

    # Generate likely Terraform resource names
    # Examples:
    # - compute.instances -> google_compute_instance
    # - storage.buckets -> google_storage_bucket
    # - resourcemanager.projects -> google_project

    $terraformNames = @()

    # Singular form (most common)
    $singular = $resource -replace 's$', ''
    $terraformNames += "google_$( $service )_$singular"

    # Plural form (if different)
    if ($singular -ne $resource)
    {
        $terraformNames += "google_$( $service )_$resource"
    }

    # Handle special cases
    # resourcemanager -> project, folder, organization
    if ($service -eq "resourcemanager")
    {
        $terraformNames += "google_$resource"
        $terraformNames += "google_$( $resource -replace 's$', '' )"
    }

    # Unique terraform names
    $entry.terraform_resources = @($terraformNames | Select-Object -Unique)
}

Write-Host "✓ Added Terraform resource name suggestions" -ForegroundColor Green

# Save to JSON
Write-Host "`nWriting to $OutputFile..." -ForegroundColor Cyan
$outputPath = path $PSScriptRoot $OutputFile
$json = $permissionsDb | ConvertTo-Json -Depth 10
$json | Out-File -FilePath $outputPath -Encoding UTF8

Write-Host "✓ Successfully exported to: $outputPath" -ForegroundColor Green

# Statistics
Write-Host "`nStatistics:" -ForegroundColor Cyan
Write-Host "  Total service.resource groups: $($permissionsDb.Count)" -ForegroundColor Gray
Write-Host "  Total permissions:             $($permissions.Count)" -ForegroundColor Gray

$gaCount = ($permissions | Where-Object { $_.stage -eq "GA" }).Count
$betaCount = ($permissions | Where-Object { $_.stage -eq "BETA" }).Count
$alphaCount = ($permissions | Where-Object { $_.stage -eq "ALPHA" }).Count

Write-Host "  GA permissions:                $gaCount" -ForegroundColor Gray
Write-Host "  BETA permissions:              $betaCount" -ForegroundColor Gray
Write-Host "  ALPHA permissions:             $alphaCount" -ForegroundColor Gray

# Show top 10 services by permission count
Write-Host "`nTop 10 services by number of resource types:" -ForegroundColor Cyan
$permissionsDb.Values | Group-Object service |
    Sort-Object Count -Descending |
    Select-Object -First 10 |
    ForEach-Object {
        Write-Host "  $($_.Name): $($_.Count) resource types" -ForegroundColor Gray
    }

Write-Host "`n✓ Done! The permissions database is ready for use with resource.ps1" -ForegroundColor Green
