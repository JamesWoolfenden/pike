#!/usr/bin/env pwsh
Add-Type -AssemblyName System.IO.Compression.FileSystem
function Unzip
{
    param([string]$zipfile, [string]$outpath)

    [System.IO.Compression.ZipFile]::ExtractToDirectory($zipfile, $outpath)
}

$root = "./"| Resolve-Path
$filepath = $root.Path + "\CloudformationSchema.zip"
write-host "path $filepath"
ls *.json| foreach {rm $_}

invoke-webrequest https://schema.cloudformation.us-east-1.amazonaws.com/CloudformationSchema.zip -OutFile $filepath
Unzip $filepath $root
Remove-Item $filepath
