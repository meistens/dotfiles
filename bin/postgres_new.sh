#!/bin/bash

# Script to copy all contents from the current directory to the project root

# Set your project root path
PROJECT_ROOT="/home/jpoxyl/Projects/ScriptStuff"

# Safety check: prevent copying if already in the project root
if [ "$PWD" = "$PROJECT_ROOT" ]; then
  echo "You are already in the project root. No files copied."
  exit 1
fi

# define specific files and folder to copy
FILES_AND_FOLDERS=(
    "db_init/"
    ".env.example"
    "compose_general.yml"
)

# for loop
for item in "${FILES_AND_FOLDERS[@]}"; do
    source_path="$PROJECT_ROOT/$item"
    if [ -e "$source_path" ]; then
        cp -a "$source_path" .
        echo "Copied: $item" from "$PROJECT_ROOT"
    else
        echo "Warning: $item not found, skipping..."
    fi
done


echo "Specified files and directories have been copied from $PROJECT_ROOT to $(pwd)"
