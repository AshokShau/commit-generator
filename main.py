import datetime
import subprocess

file_name = "commit_file.txt"
start_date = "2024-09-24"
end_date = "2024-09-28"
commit_message = "Test-Automation-Commit"
commits_per_day = 2
author_name = "AshokShau"
author_email = "114943948+AshokShau@users.noreply.github.com"

def log_info(message):
    print(f"[INFO] {message}")

def log_error(message):
    print(f"[ERROR] {message}")

def append_to_file(filename, content):
    with open(filename, 'a') as f:
        f.write(content)

def git_add(filename):
    subprocess.run(["git", "add", filename], check=True)

def git_commit(current_date, commit_index):
    env = {
        "GIT_COMMITTER_DATE": current_date.strftime("%Y-%m-%d %H:%M:%S"),
        "GIT_AUTHOR_DATE": current_date.strftime("%Y-%m-%d %H:%M:%S"),
        "GIT_AUTHOR_NAME": author_name,
        "GIT_AUTHOR_EMAIL": author_email,
        "GIT_COMMITTER_NAME": author_name,
        "GIT_COMMITTER_EMAIL": author_email
    }
    subprocess.run(["git", "commit", "--date", current_date.strftime("%Y-%m-%d %H:%M:%S"),
                    "-m", f"{commit_message} - {commit_index + 1}"], env=env, check=True)

def git_push():
    subprocess.run(["git", "push", "origin", "master"], check=True)

def main():
    try:
        start = datetime.datetime.strptime(start_date, "%Y-%m-%d")
        end = datetime.datetime.strptime(end_date, "%Y-%m-%d")
    except ValueError as e:
        log_error(f"Error parsing date: {e}")
        return

    with open(file_name, 'w') as f:
        f.write("Commits:\n")

    current_date = start
    while current_date <= end:
        for i in range(commits_per_day):
            content = f"This is commit number {i + 1} for {current_date.strftime('%Y-%m-%d')}.\n"
            append_to_file(file_name, content)

            git_add(file_name)
            git_commit(current_date, i)

        current_date += datetime.timedelta(days=1)

    git_push()
    log_info("Changes pushed to repository successfully.")

if __name__ == "__main__":
    main()
