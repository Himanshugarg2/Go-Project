
# Go Project Collection

This repository contains two command-line projects written in Go:

1. **Task Management System**  
2. **GitHub User Activity CLI**  https://roadmap.sh/projects/github-user-activity

---

## 1. Task Management System

A simple command-line Task Management System. Allows you to add, update, delete, and list tasks, as well as save/load them in JSON format.

**Features:**
- Add, update, and delete tasks
- Mark tasks as done, not done, or in progress
- List all tasks or filter by status
- Save and load tasks to/from a JSON file

**Folder:** `taskManagementSystem/`

**How to Run:**
```sh
cd Projects/Go-Project/taskManagementSystem
go run .
```
Follow the on-screen menu to manage your tasks.

**Example Menu:**
```
Task Management System Menu:
1. Add Task
2. List All Tasks
3. List Tasks by Status
4. Update Task Status
5. Delete Task
6. Save Tasks to File
7. Load Tasks from File
0. Exit
```

---

## 2. GitHub User Activity CLI

A Go CLI tool to fetch and display recent public activity for any GitHub user using the GitHub API.

**Features:**
- Fetches recent activity for a specified GitHub username
- Displays event types and details (pushes, repo creation, membership, etc.)
- Handles user not found and no activity cases

**Folder:** `gitHubUserActivity/`

**How to Run:**
```sh
cd Projects/Go-Project/gitHubUserActivity
go run main.go
```
Enter the GitHub username when prompted.

**Example Output:**
```
--------------------------------------------------
Event 1:
Event Type: PushEvent
Event Details:
 - Commit: 123abc | Message: Update README
Repository Name: user/repo
Created At: 2025-08-26T10:00:00Z
--------------------------------------------------
```

---

## Requirements
- Go 1.18 or higher

## License
MIT
