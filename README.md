# Git Commit Automation

This project automates the process of creating multiple Git commits over a specified date range. It generates a single file with commit messages, commits the changes, and pushes them to a remote repository.

## Disclaimer

This project is intended for educational and demonstration purposes only. The automated creation of multiple Git commits may not reflect best practices in version control and should be used responsibly. Users are encouraged to understand the implications of mass commits in a collaborative environment and to adapt the script to fit their specific needs. The author, AshokShau, is not responsible for any unintended consequences resulting from the use of this tool.


## Overview

- **Author**: AshokShau
- **Purpose**: To automate Git commits for testing or demonstration purposes.
- **Language**: Go

## Features

- Generates multiple commits for each day within a specified date range.
- Uses customizable commit messages and author information.
- Handles file creation and appending automatically.
- Pushes changes to the specified remote repository.

## Prerequisites

- Go installed on your machine.
- Git installed and configured with a remote repository.
- Ensure you have the necessary permissions to push changes to the repository.

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/AshokShau/commit-generator.git
cd commit-generator
```

### Modify the Configuration

##### Edit the constants in `main.py` or `main.go` to set your desired configuration

- `fileName`: Name of the file where commit messages will be stored.
- `startDate`: The start date for generating commits (format: YYYY-MM-DD).
- `endDate`: The end date for generating commits (format: YYYY-MM-DD).
- `commitMessage`: The base message for the commits.
- `commitsPerDay`: Number of commits to create per day.
- `authorName`: Your name as it will appear in commits.
- `authorEmail`: Your email as it will appear in commits.
- `branchName`: The name of the branch to push changes to.

### Run the Program

To run the program, execute:

```bash
go run main.go
```
or 
```bash
python3 main.py
```

The program will create a file with commit messages, generate commits, and push the changes to the remote repository.

### Verify Commits

After the program completes, you can verify the commits in your Git repository:

```bash
git log
```

## Contributing

Contributions are welcome! If you have suggestions for improvements or new features, feel free to create an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Contact

For any inquiries, you can reach me at [AshokShau](https://github.com/AshokShau) or on Telegram at [AshokShau](https://t.me/AshokShau).
