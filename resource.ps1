#!/usr/bin/env pwsh
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

$provider = $resource.Split("_")[0]
$baseMapping = path $PSScriptRoot "src" "mapping" $provider
$mapping = path $baseMapping $type

write-host "Copying to $baseMapping"

if (test-path($baseMapping))
{
    $source = path $mapping "template.json"
    $destination = path $mapping "$resource.json"
    Copy-Item  -path $source -destination $destination
}
else
{
    write-host "$baseMapping not found"
}

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

$tffile=path $PSScriptRoot $tffile
new-item $tffile -value $content
