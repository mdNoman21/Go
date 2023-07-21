package models

type Task struct {
	TaskID      string `json:"taskID"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
type Response struct {
	TaskID    string `json:"id,omitempty"`
	Message   string `json:"message,omitempty"`
	Completed bool   `json:"completed,omitempty"`
}
type TaskStats struct {
	TotalTasks     int `json:"totalTasks"`
	CompletedTasks int `json:"completedTasks"`
}
