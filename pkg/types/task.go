package types

type Task struct {
	TaskType string   `json:"task_type"`
	TaskTime string   `json:"task_time"`
	IDs      []uint64 `json:"ids"`
}
