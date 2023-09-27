#!/bin/bash

# Check if a commit message is provided
if [ -z "$1" ]; then
  echo "Usage: $0 'Your commit message'"
  exit 1
fi

# Calculate the date for one year ago
# Use -d for Linux and -v for macOS
if date --help >/dev/null 2>&1; then
  # Linux
  OLD_DATE=$(date -d "1 year ago" +"%Y-%m-%d %H:%M:%S")
else
  # macOS
  OLD_DATE=$(date -v -1y +"%Y-%m-%d %H:%M:%S")
fi

# Change to the directory of your Git repository (optional)
# cd /path/to/your/repo

# Stage all changes
git add .

# Commit with the calculated date
GIT_COMMITTER_DATE="$OLD_DATE" git commit --date="$OLD_DATE" -m "$1"

# Push changes to the remote repository (defaulting to master)
git push origin master
