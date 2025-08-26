# Task Management System (Go)

A simple command-line Task Management System written in Go. This project allows you to add, update, delete, and list tasks, as well as save/load them in JSON format.

https://roadmap.sh/projects/task-tracker

## Features
- Add, update, and delete tasks
- Mark tasks as done, not done, or in progress
- List all tasks or filter by status
- Save and load tasks to/from a JSON file

## Folder Structure
```
taskManagementSystem/
├── CRUDfunctions.go
├── FileManagementFunctions.go
├── go.mod
├── main.go
├── models.go
├── tasks.json
```

## How to Run

1. **Navigate to the project directory:**
   ```sh
   cd Projects/Go-Project/taskManagementSystem
   ```

2. **Run the program:**
   ```sh
   go run .
   ```

3. **Follow the on-screen menu** to add, update, delete, or list tasks.

## Saving and Loading Tasks
- Tasks are saved to `tasks.json` in the same directory.
- You can load tasks from this file when you restart the program.

## Requirements
- Go 1.18 or higher

## Example Usage
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
Enter your choice: 1
Enter Task ID: 1
Enter Task Name: Example
Enter Description: Example task
Select Status: 1. done 2. not done 3. in progress
Task added.
```

## License
MIT
