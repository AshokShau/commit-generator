import subprocess
from datetime import datetime, timedelta

# Define the starting date and commit message
start_date = datetime(2000, 1, 1)
end_date = datetime(2000, 12, 30)
commit_message = "Your commit message"  # Change this to your desired commit message

# Stage all changes initially
subprocess.run(["git", "add", "."], check=True)

# Loop through each day and create a commit
current_date = start_date
while current_date <= end_date:
    # Set the environment variables for the commit date
    env = {
        "GIT_COMMITTER_DATE": current_date.strftime("%Y-%m-%d %H:%M:%S"),
        "GIT_AUTHOR_DATE": current_date.strftime("%Y-%m-%d %H:%M:%S"),
        "GIT_AUTHOR_NAME": "AshokShau",
        "GIT_AUTHOR_EMAIL": "114943948+AshokShau@users.noreply.github.com",  # Update this line
        "GIT_COMMITTER_NAME": "AshokShau",  # Added line
        "GIT_COMMITTER_EMAIL": "114943948+AshokShau@users.noreply.github.com"  # Added line
    }

    # Create the commit
    print(f"Committing for date: {current_date.strftime('%Y-%m-%d %H:%M:%S')}")  # Debugging output
    result = subprocess.run(
        ["git", "commit", "--date", current_date.strftime("%Y-%m-%d %H:%M:%S"), "-m", commit_message],
        env=env,
        capture_output=True  # Capture output for debugging
    )

    # Check if the command was successful
    if result.returncode != 0:
        print(f"Error committing: {result.stderr.decode()}")  # Print error message
        break  # Exit loop on error

    # Increment the date by one day
    current_date += timedelta(days=1)

# Push changes to the remote repository (defaulting to master)
subprocess.run(["git", "push", "origin", "master"], check=True)
