
# GitHub User Activity CLI


A simple Go command-line tool to fetch and display recent public activity for any GitHub user using the GitHub API.

## Features
- Fetches recent activity for a specified GitHub username
- Displays event types and details (pushes, repo creation, membership, etc.)
- Handles user not found and no activity cases

## Usage

### 1. Navigate to the project directory
```
cd Projects/Go-Project/gitHubUserActivity
```

### 2. Run the program
```
go run main.go
```

### 3. Enter the GitHub username when prompted
```
Enter User's Github username
> kamranahmedse
```

### 4. Example Output
```
--------------------------------------------------
Event 1:
Event Type: PushEvent
Event Details:
 - Commit: 123abc | Message: Update README
Repository Name: kamranahmedse/developer-roadmap
Created At: 2025-08-26T10:00:00Z
--------------------------------------------------
Event 2:
Event Type: CreateEvent
Event Details:
 - Created repository: new-repo
Repository Name: kamranahmedse/new-repo
Created At: 2025-08-25T09:00:00Z
...
```

## How it works
- Prompts for a GitHub username
- Fetches recent events from `https://api.github.com/users/<username>/events`
- Parses and displays event details in a readable format

## Requirements
- Go 1.18 or higher
- Internet connection

## Notes
- Only public events are shown (private activity is not available via the API)
- The tool uses the standard Go library (no third-party dependencies)

## License
MIT
