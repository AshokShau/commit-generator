#!/bin/bash

# Check if a commit message is provided
if [ -z "$1" ]; then
  echo "Usage: $0 'Your commit message'"
  exit 1
fi

# Set the date to a specific date in 2000
OLD_DATE="2000-01-01 12:00:00"  # Change the date and time as needed

# Change to the directory of your Git repository (optional)
# cd /path/to/your/repo

# Stage all changes
git add .

# Commit with the specified date
GIT_COMMITTER_DATE="$OLD_DATE" git commit --date="$OLD_DATE" -m "$1"

# Push changes to the remote repository (defaulting to master)
git push origin master
