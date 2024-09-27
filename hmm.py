    import subprocess
    from datetime import datetime, timedelta

    start_date = datetime(2023, 1, 1)
    end_date = datetime(2024, 9, 27)
    commit_message = "Test"
    commits_per_day = 10  # Number of commits per day

    current_date = start_date
    while current_date <= end_date:
        for i in range(commits_per_day):
            filename = f"file_{current_date.strftime('%Y-%m-%d')}_{i}.txt"
            with open(filename, 'w') as f:
                f.write(f"This is commit number {i + 1} for {current_date.strftime('%Y-%m-%d')}.\n")


            subprocess.run(["git", "add", filename], check=True)
            env = {
                "GIT_COMMITTER_DATE": current_date.strftime("%Y-%m-%d %H:%M:%S"),
                "GIT_AUTHOR_DATE": current_date.strftime("%Y-%m-%d %H:%M:%S"),
                "GIT_AUTHOR_NAME": "AshokShau",
                "GIT_AUTHOR_EMAIL": "114943948+AshokShau@users.noreply.github.com",
                "GIT_COMMITTER_NAME": "AshokShau",
                "GIT_COMMITTER_EMAIL": "114943948+AshokShau@users.noreply.github.com"
            }

            print(f"Committing for date: {current_date.strftime('%Y-%m-%d %H:%M:%S')} - Commit {i + 1}")
            result = subprocess.run(
                ["git", "commit", "--date", current_date.strftime("%Y-%m-%d %H:%M:%S"), "-m", f"{commit_message} - {i + 1}"],
                env=env,
                capture_output=True
            )

            if result.returncode != 0:
                print(f"Error committing: {result.stderr.decode()}")
                break
        current_date += timedelta(days=1)

    subprocess.run(["git", "push", "origin", "master"], check=True)
