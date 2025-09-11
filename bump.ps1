$version = $( git describe --tags --abbrev = 0 )
$splitter = $version.split(".")
$build = [int]($splitter[2]) + 1
$newVersion = $splitter[0] + "." + $splitter[1] + "." + $build

write-host $newVersion
git tag -a $newVersion -m "new release"
git push origin $newVersion
