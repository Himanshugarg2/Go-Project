package main

import (
	"encoding/json"
	"os"
)

func (t *Tasks) SaveToFile(filename string) error {
	data, err := json.MarshalIndent(t.Task, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func (t *Tasks) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			t.Task = []Task{}
			return t.SaveToFile(filename)
		}
		return err
	}
	return json.Unmarshal(data, &t.Task)
}
