#!/bin/bash

# Leverage the default env variables as described in:
# https://docs.github.com/en/actions/reference/environment-variables#default-environment-variables
if [[ $GITHUB_ACTIONS != "true" ]]
then
  /usr/bin/pike "$@"
  exit $?
fi

echo "running command:"

args=()

if [ -n "$INPUT_DIRECTORY" ]
then
  args+=(--directory "$INPUT_DIRECTORY")
fi

if [ -n "$INPUT_FILE" ]
then
  args+=(--file "$INPUT_FILE")
fi

if [ -n "$INPUT_AUTO" ]
then
  args+=(--auto)
fi

if [ -n "$INPUT_WRITE" ] && [ "$INPUT_WRITE" = "true" ]
then
  args+=(--write)
fi

if [ -n "$INPUT_INIT" ] && [ "$INPUT_INIT" = "true" ]
then
  args+=(--init)
fi

if [ -n "$INPUT_OUTPUT" ]
then
  args+=(--output "$INPUT_OUTPUT")
fi

echo pike "$INPUT_VERB" "${args[@]}"
/usr/bin/pike "$INPUT_VERB" "${args[@]}"

export pike_EXIT_CODE=$?
