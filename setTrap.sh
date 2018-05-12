#!/usr/bin/env bash

# trap 'history 2|head -n 1|cut -c 8-' DEBUG

# trap 'command=$(echo !!); ./application.exe "$command"' DEBUG

function prompt {
  command=$(history 1|head -n 1|cut -c 8-)
  # echo "last command: $command"
  ./application.exe "$command"
}

PROMPT_COMMAND=prompt
