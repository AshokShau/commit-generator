#!/bin/bash

# Check if a commit message is provided
if [ -z "$1" ]; then
  echo "Usage: $0 'Your commit message'"
  exit 1
fi

# Set the starting date to a specific date in 2000
OLD_DATE="2000-01-01 12:00:00"  # Change the date and time as needed

# Change to the directory of your Git repository (optional)
# cd /path/to/your/repo

# Stage all changes
git add .

# Commit 200 times
# shellcheck disable=SC2034
for i in {1..200}; do
  GIT_COMMITTER_DATE="$OLD_DATE" git commit --date="$OLD_DATE" -m "$1"

  # Increment the date by one day for the next commit
  OLD_DATE=$(date -d "$OLD_DATE + 1 day" +"%Y-%m-%d %H:%M:%S")

  # Optionally, stage changes again if needed
  git add .
done

# Push changes to the remote repository (defaulting to master)
git push origin master
