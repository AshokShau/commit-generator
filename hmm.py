import subprocess
from datetime import datetime, timedelta

# Define the starting date, ending date, and commit message
start_date = datetime(2024, 1, 1)
end_date = datetime(2024, 12, 30)
commit_message = "Your commit message"  # Change this to your desired commit message
commits_per_day = 10  # Number of commits per day

# Loop through each day
current_date = start_date
while current_date <= end_date:
    for i in range(commits_per_day):
        # Create a new file or modify an existing one for each commit
        filename = f"file_{current_date.strftime('%Y-%m-%d')}_{i}.txt"
        with open(filename, 'w') as f:
            f.write(f"This is commit number {i + 1} for {current_date.strftime('%Y-%m-%d')}.\n")

        # Stage the change
        subprocess.run(["git", "add", filename], check=True)

        # Set the environment variables for the commit date
        env = {
            "GIT_COMMITTER_DATE": current_date.strftime("%Y-%m-%d %H:%M:%S"),
            "GIT_AUTHOR_DATE": current_date.strftime("%Y-%m-%d %H:%M:%S"),
            "GIT_AUTHOR_NAME": "AshokShau",
            "GIT_AUTHOR_EMAIL": "114943948+AshokShau@users.noreply.github.com",
            "GIT_COMMITTER_NAME": "AshokShau",
            "GIT_COMMITTER_EMAIL": "114943948+AshokShau@users.noreply.github.com"
        }

        # Create the commit
        print(f"Committing for date: {current_date.strftime('%Y-%m-%d %H:%M:%S')} - Commit {i + 1}")
        result = subprocess.run(
            ["git", "commit", "--date", current_date.strftime("%Y-%m-%d %H:%M:%S"), "-m", f"{commit_message} - {i + 1}"],
            env=env,
            capture_output=True
        )

        # Check if the command was successful
        if result.returncode != 0:
            print(f"Error committing: {result.stderr.decode()}")
            break  # Exit loop on error

    # Increment the date by one day
    current_date += timedelta(days=1)

# Push changes to the remote repository (defaulting to master)
subprocess.run(["git", "push", "origin", "master"], check=True)
