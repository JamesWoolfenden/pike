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
if [ -n "$INPUT_DIRECTORY" ]
then
  echo pike "$INPUT_VERB" -d "$INPUT_DIRECTORY" "$flags"

  /usr/bin/pike "$INPUT_VERB" -d "$INPUT_DIRECTORY" "$flags"
fi

if [ -n "$INPUT_FILE" ]
then
  echo pike "$INPUT_VERB" -f "$INPUT_FILE" "$flags"

  /usr/bin/pike "$INPUT_VERB" -f "$INPUT_FILE" "$flags"
fi

export pike_EXIT_CODE=$?
