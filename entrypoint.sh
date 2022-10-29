#!/bin/bash

# Leverage the default env variables as described in:
# https://docs.github.com/en/actions/reference/environment-variables#default-environment-variables
if [[ $GITHUB_ACTIONS != "true" ]]
then
  /usr/bin/pike "$@"
  exit $?
fi

echo "running command:"
if [ -n "$INPUT_DIRECTORY" ]
then
  FLAG="--directory"
  COMMANDS="$INPUT_DIRECTORY"
fi

if [ -n "$INPUT_FILE" ]
then
  FLAG="--file"
  COMMANDS="$INPUT_FILE"
fi

if [ -n "$INPUT_AUTO" ];
then
  COMMANDS+=" --auto"
fi

if [ -n "$INPUT_WRITE" ] && [ "$INPUT_WRITE" = "true" ];
then
  COMMANDS+=" --write"
fi

if [ -n "$INPUT_INIT" ] && [ "$INPUT_INIT" = "true" ];
then
  COMMANDS+=" --init"
fi

if [ -n "$INPUT_OUTPUT" ]
then
  COMMANDS+=" --output "+ $INPUT_OUTPUT
fi

echo pike $INPUT_VERB $FLAG $COMMANDS
/usr/bin/pike $INPUT_VERB $FLAG $COMMANDS

export pike_EXIT_CODE=$?
