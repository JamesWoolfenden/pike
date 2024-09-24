#!/usr/bin/env pwsh
$files = get-childitem -path .\*.json

foreach ($file in $files)
{
    write-host $file.Name
    $filecontent = get-content $file|ConvertFrom-Json

    if ($filecontent.typeName)
    {
        write-host $filecontent.typeName exists
    }

    if (!(& "sato.exe" see -r $filecontent.typeName))
    {
        write-host $filecontent.typeName missing
    }
    else
    {
        $tf = & "sato.exe" see -r $filecontent.typeName
        write-host $tf exists
        write-host $tf create
        write-host $filecontent.handlers.create.permissions
        write-host $tf delete
        write-host $filecontent.handlers.delete.permissions
        write-host $tf read
        write-host $filecontent.handlers.read.permissions
        write-host $tf update
        write-host $filecontent.handlers.update.permissions
        Pause
    }
}
