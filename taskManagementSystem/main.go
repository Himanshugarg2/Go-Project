// Requirements
// The application should run from the command line, accept user actions and inputs as arguments,
// and store the tasks in a json file. The user should be able to:
//
// Add, Update, and Delete tasks
// Mark a task as in progress or done
// List all tasks
// List all tasks that are done
// List all tasks that are not done
// List all tasks that are in progress

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

func readLine(prompt string) string {
	fmt.Println(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func parseTimeInput(prompt string, defaultTime time.Time) time.Time {
	input := readLine(prompt)
	if input == "" {
		return defaultTime
	}
	parsed, err := time.Parse("2006-01-02 15:04:05", input)
	if err != nil {
		fmt.Println("Invalid format, using default.")
		return defaultTime
	}
	return parsed
}

func takeInput() Task {
	id := readLine("Enter Task ID:")
	description := readLine("Enter Task Description:")

	startTime := parseTimeInput(
		"Enter Start Time (YYYY-MM-DD HH:MM:SS): leave blank to use current time",
		time.Now(),
	)

	endTime := parseTimeInput(
		"Enter End Time (YYYY-MM-DD HH:MM:SS): leave blank to use current time + 2 days",
		time.Now().Add(48*time.Hour),
	)

	fmt.Println("Enter Status:leave blank for default/current status")
	fmt.Println("1 = Made, 2 = In Progress, 3 = Done")
	var choice int
	fmt.Scanln(&choice)

	var status Status
	switch choice {
	case 1:
		status = StatusMade
	case 2:
		status = StatusInProgress
	case 3:
		status = StatusDone
	default:
		status = StatusMade
	}

	return Task{
		ID:          id,
		Description: description,
		StartTime:   startTime,
		EndTime:     endTime,
		Status:      status,
	}
}

func PrintChoices() {
	fmt.Println("1. Add Task")
	fmt.Println("2. Update Task")
	fmt.Println("3. Update Task Status")
	fmt.Println("4. Delete Task")
	fmt.Println("5. List All Tasks")
	fmt.Println("6. List Done Tasks")
	fmt.Println("7. List Not Done Tasks")
	fmt.Println("8. List In Progress Tasks")
	fmt.Println("9. Exit")
}

func main() {
	tasks := Tasks{}
	for {
		tasks.LoadFromFile("tasks.json")

		var choice int
		PrintChoices()
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			newTask := takeInput()
			tasks.AddTask(newTask)

		case 2:
			updatedTask := takeInput()
			tasks.UpdateTask(updatedTask)

		case 3:
			id := readLine("Enter Task ID to update status:")
			fmt.Println("Enter 1 for Made, 2 for In Progress, 3 for Done")
			var statusChoice int
			fmt.Scanln(&statusChoice)

			switch statusChoice {
			case 1:
				tasks.updateTask(id, StatusMade)
			case 2:
				tasks.updateTask(id, StatusInProgress)
			case 3:
				tasks.updateTask(id, StatusDone)
			}

		case 4:
			id := readLine("Enter Task ID to delete:")
			tasks.DeleteTask(id)

		case 5:
			tasks.listAllTasks()

		case 6:
			tasks.ListTasksByStatus(StatusDone)

		case 7:
			tasks.ListTasksByStatus(StatusMade)

		case 8:
			tasks.ListTasksByStatus(StatusInProgress)

		case 9:
			return
		}
	}
}
