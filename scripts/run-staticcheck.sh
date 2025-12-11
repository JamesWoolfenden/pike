#!/bin/bash
# Run staticcheck excluding generated files

# Run staticcheck and capture output
output=$(staticcheck ./... 2>&1)

# Filter out generated files
filtered=$(echo "$output" | grep -v "generated_api.go" | grep -v "generated_client.go")

# If there's filtered output, print it and exit with error
if [ -n "$filtered" ]; then
    echo "$filtered"
    exit 1
fi

# No issues found
exit 0
