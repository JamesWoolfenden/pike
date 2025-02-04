#!/bin/bash

# Set working directory
ROOT_DIR="$(pwd)"
SCHEMA_URL="https://schema.cloudformation.us-east-1.amazonaws.com/CloudformationSchema.zip"
FILEPATH="${ROOT_DIR}/CloudformationSchema.zip"

# Function to cleanup on exit
cleanup() {
    rm -f "${FILEPATH}"
}

# Error handling
set -e
trap cleanup EXIT

# Remove existing JSON files
rm -f ./*.json

# Download schema
echo "Downloading schema..."
curl -L -o "${FILEPATH}" "${SCHEMA_URL}" || {
    echo "Failed to download schema"
    exit 1
}

# Extract zip file
echo "Extracting schema..."
unzip -o "${FILEPATH}" || {
    echo "Failed to extract schema"
    exit 1
}
