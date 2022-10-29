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
  COMMANDS="-d $INPUT_DIRECTORY"
fi

if [ -n "$INPUT_FILE" ]
then
  COMMANDS="-f $INPUT_FILE"
fi

echo pike "$INPUT_VERB" "$COMMANDS" "$flags"
/usr/bin/pike "$INPUT_VERB" "$COMMANDS" "$flags"

export pike_EXIT_CODE=$?
