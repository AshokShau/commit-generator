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
        "GIT_AUTHOR_EMAIL": "Abishnoi69@outlook.com"
    }

    # Create the commit
    subprocess.run(
        ["git", "commit", "--date", current_date.strftime("%Y-%m-%d %H:%M:%S"), "-m", commit_message],
        env=env,
        check=True
    )

    # Increment the date by one day
    current_date += timedelta(days=1)

# Push changes to the remote repository (defaulting to master)
subprocess.run(["git", "push", "origin", "master"], check=True)
