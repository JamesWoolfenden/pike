#!/bin/sh
latestTag=$(git describe --tags)
echo "Updating version file with new tag: $latestTag"
echo "package pike" > src/version.go
echo "" >> src/version.go
echo "const Version = \"$latestTag\"" >> src/version.go
