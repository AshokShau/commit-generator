package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

const (
	startDate     = "2024-09-27"
	endDate       = "2024-09-27"
	commitMessage = "Test"
	commitsPerDay = 20
	authorName    = "AshokShau"
	authorEmail   = "114943948+AshokShau@users.noreply.github.com"
)

func main() {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		log.Print("Error parsing start date:", err)
		return
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		log.Print("Error parsing end date:", err)
		return
	}

	filename := "commits.txt"
	fileContent := ""

	currentDate := start
	for currentDate.Before(end) || currentDate.Equal(end) {
		for i := 0; i < commitsPerDay; i++ {
			fileContent += fmt.Sprintf("This is commit number %d for %s.\n", i+1, currentDate.Format("2006-01-02"))
			if err := gitCommit(currentDate, i); err != nil {
				log.Print("Error committing:", err)
				return
			}
		}
		currentDate = currentDate.AddDate(0, 0, 1)
	}

	// Write all commit messages to the single file
	if err := os.WriteFile(filename, []byte(fileContent), 0644); err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	if err := gitAdd(filename); err != nil {
		fmt.Println("Error adding file:", err)
		return
	}

	if err := gitPush(); err != nil {
		fmt.Println("Error pushing to repository:", err)
	}
}

func gitAdd(filename string) error {
	cmd := exec.Command("git", "add", filename)
	return cmd.Run()
}

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

func gitPush() error {
	cmd := exec.Command("git", "push", "origin", "master")
	return cmd.Run()
}
