package main

import "fmt"

func (t *Tasks) AddTask(T Task) {
	t.Task = append(t.Task, T)
	t.SaveToFile("tasks.json")
}

func (t *Tasks) DeleteTask(ID string) {
	for i, task := range t.Task {
		if task.ID == ID {
			t.Task = append(t.Task[:i], t.Task[i+1:]...)
			t.SaveToFile("tasks.json")
			return
		}
	}
}

func (t *Tasks) UpdateTask(T Task) {
	for i, task := range t.Task {
		if task.ID == T.ID {
			if T.Description != "" {
				t.Task[i].Description = T.Description
			}
			if !T.StartTime.IsZero() {
				t.Task[i].StartTime = T.StartTime
			}
			if !T.EndTime.IsZero() {
				t.Task[i].EndTime = T.EndTime
			}
			if string(T.Status) != "" {
				t.Task[i].Status = T.Status
			}
		}
	}
	t.SaveToFile("tasks.json")
}

func (t *Tasks) updateTask(ID string, Status Status) {
	for i, task := range t.Task {
		if task.ID == ID {
			t.Task[i].Status = Status
		}
	}
	t.SaveToFile("tasks.json")
}

func (t *Tasks) listAllTasks() {
	for _, task := range t.Task {
		fmt.Println("TASKS")
		fmt.Printf("Task ID: %s\n", task.ID)
		fmt.Printf("Description: %s\n", task.Description)
		fmt.Printf("Start Time: %s\n", task.StartTime.Format("Monday, 02-Jan-2006 15:04:05"))
		fmt.Printf("End Time: %s\n", task.EndTime.Format("Monday, 02-Jan-2006 15:04:05"))
		fmt.Printf("Status: %s\n", task.Status)
	}
}

func (t *Tasks) ListTasksByStatus(status Status) {
	for _, task := range t.Task {
		if task.Status == status {
			fmt.Println("TASKS")
			fmt.Printf("Task ID: %s\n", task.ID)
			fmt.Printf("Description: %s\n", task.Description)
			fmt.Printf("Start Time: %s\n", task.StartTime.Format("Monday, 02-Jan-2006 15:04:05"))
			fmt.Printf("End Time: %s\n", task.EndTime.Format("Monday, 02-Jan-2006 15:04:05"))
			fmt.Printf("Status: %s\n", task.Status)
		}
	}
}
