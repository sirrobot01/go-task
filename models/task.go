package models

type Task struct {
	ID   int    `json:"id"`
	Day  int8   `json:"day"`
	Name string `json:"name"`
}

var (
	Tasks = map[int]*Task{}
	Seq   = 1
)
