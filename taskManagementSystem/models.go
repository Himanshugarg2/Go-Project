package main

import "time"

type Status string

const (
	StatusDone       Status = "done"
	StatusMade       Status = "made"
	StatusInProgress Status = "in progress"
)

type Task struct {
	ID          string    `json:"id"`
	Description string    `json:"Description"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"EndTime"`
	Status      Status    `json:"Status"`
}

type Tasks struct {
	Task []Task
}
