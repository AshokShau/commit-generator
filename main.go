package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"
)

const (
<<<<<<< Updated upstream
	startDate     = "2024-09-26"
	endDate       = "2024-09-28"
=======
	fileName      = "commit_file.txt" // Single file for all commits
	startDate     = "2021-09-25"
	endDate       = "2024-09-27"
>>>>>>> Stashed changes
	commitMessage = "Test"
	commitsPerDay = 1
	authorName    = "AshokShau"
	authorEmail   = "114943948+AshokShau@users.noreply.github.com"
	repoOwner     = "AshokShau"
	repoName      = "github-bot"
	prBaseBranch  = "master" // or any other branch you want to merge into
	token         = ""       // GitHub token with repo permissions
)

func main() {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		fmt.Println("Error parsing start date:", err)
		return
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		fmt.Println("Error parsing end date:", err)
		return
	}

	currentDate := start

	// Create or open the single commit file
	if err := os.WriteFile(fileName, []byte("Commits:\n"), 0644); err != nil {
		fmt.Println("Error writing initial file:", err)
		return
	}

	for currentDate.Before(end) || currentDate.Equal(end) {
		for i := 0; i < commitsPerDay; i++ {
			// Append content to the file
			if err := appendToFile(fileName, fmt.Sprintf("This is commit number %d for %s.\n", i+1, currentDate.Format("2006-01-02"))); err != nil {
				fmt.Println("Error appending to file:", err)
				return
			}

			if err := gitAdd(fileName); err != nil {
				fmt.Println("Error adding file:", err)
				return
			}

			if err := gitCommit(currentDate, i); err != nil {
				fmt.Println("Error committing:", err)
				return
			}
		}
		currentDate = currentDate.AddDate(0, 0, 1)
	}

	if err := gitPush(); err != nil {
		fmt.Println("Error pushing to repository:", err)
		return
	}

	if err := createPullRequest(); err != nil {
		fmt.Println("Error creating pull request:", err)
	}
}

func appendToFile(filename, content string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
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
	cmd := exec.Command("git", "push", "origin", "master") // Change if using a different branch
	return cmd.Run()
}
<<<<<<< Updated upstream

func createPullRequest() error {
	prURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/pulls", repoOwner, repoName)

	prData := map[string]interface{}{
		"title": "Automated Pull Request",
		"head":  "your-feature-branch", // Replace with the branch you pushed to
		"base":  prBaseBranch,
		"body":  "This is an automated pull request.",
	}

	jsonData, err := json.Marshal(prData)
	if err != nil {
		return fmt.Errorf("error marshaling PR data: %w", err)
	}

	req, err := http.NewRequest("POST", prURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to create pull request: %s", resp.Status)
	}

	fmt.Println("Pull request created successfully!")
	return nil
}
=======
```
>>>>>>> Stashed changes
