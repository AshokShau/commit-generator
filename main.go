package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

const (
	fileName      = "commit_file.txt"
	startDate     = "2024-09-24"
	endDate       = "2024-09-28"
	commitMessage = "Test-Automation-Commit"
	commitsPerDay = 30
	authorName    = "AshokShau"
	authorEmail   = "114943948+AshokShau@users.noreply.github.com"
)

// main is the entry point of the program. It generates multiple git commits
// for a specified date range and pushes them to a remote repository.
func main() {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		logError("Error parsing start date:", err)
		return
	}
	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		logError("Error parsing end date:", err)
		return
	}

	if err := os.WriteFile(fileName, []byte("Commits:\n"), 0644); err != nil {
		logError("Error writing initial file:", err)
		return
	}

	for currentDate := start; !currentDate.After(end); currentDate = currentDate.AddDate(0, 0, 1) {
		for i := 0; i < commitsPerDay; i++ {
			content := fmt.Sprintf("This is commit number %d for %s.\n", i+1, currentDate.Format("2006-01-02"))
			if err := appendToFile(fileName, content); err != nil {
				logError("Error appending to file:", err)
				return
			}
			if err := gitAdd(fileName); err != nil {
				logError("Error adding file:", err)
				return
			}
			if err := gitCommit(currentDate, i); err != nil {
				logError("Error committing:", err)
				return
			}
		}
	}

	if err := gitPush(); err != nil {
		logError("Error pushing to repository:", err)
	} else {
		logInfo("Changes pushed to repository successfully.")
	}
}

// appendToFile appends the given content to the specified file.
func appendToFile(filename, content string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	return err
}

// gitAdd stages the specified file for commit.
func gitAdd(filename string) error {
	return exec.Command("git", "add", filename).Run()
}

// gitCommit creates a git commit with the specified date and commit index.
func gitCommit(currentDate time.Time, commitIndex int) error {
	env := []string{
		fmt.Sprintf("GIT_COMMITTER_DATE=%s", currentDate.Format("2006-01-02 15:04:05")),
		fmt.Sprintf("GIT_AUTHOR_DATE=%s", currentDate.Format("2006-01-02 15:04:05")),
		fmt.Sprintf("GIT_AUTHOR_NAME=%s", authorName),
		fmt.Sprintf("GIT_AUTHOR_EMAIL=%s", authorEmail),
		fmt.Sprintf("GIT_COMMITTER_NAME=%s", authorName),
		fmt.Sprintf("GIT_COMMITTER_EMAIL=%s", authorEmail),
	}
	cmd := exec.Command("git", "commit", "--date", currentDate.Format("2006-01-02 15:04:05"), "-m", fmt.Sprintf("%s - %d", commitMessage, commitIndex+1))
	cmd.Env = env
	return cmd.Run()
}

// gitPush pushes the committed changes to the remote repository.
func gitPush() error {
	return exec.Command("git", "push", "origin", "master").Run()
}

// logInfo logs informational messages to the console.
func logInfo(v ...interface{}) {
	fmt.Println("[INFO]", fmt.Sprint(v...))
}

// logError logs error messages to the console.
func logError(v ...interface{}) {
	fmt.Println("[ERROR]", fmt.Sprint(v...))
}
