#!/bin/bash

# Leverage the default env variables as described in:
# https://docs.github.com/en/actions/reference/environment-variables#default-environment-variables
if [[ $GITHUB_ACTIONS != "true" ]]
then
  /usr/bin/pike "$@"
  exit $?
fi

flags=""

echo "running command:"
echo pike scan -d "$INPUT_DIRECTORY" "$flags"

/usr/bin/pike scan -d "$INPUT_DIRECTORY" "$flags"
export pike_EXIT_CODE=$?
