package data

import "time"

type TaskStatus string

const (
	PENDING   TaskStatus = "pending"
	COMPLETED            = "completed"
)

type Todo struct {
	ID        int64      `json:"id"`
	Task      string     `json:"task"`
	Status    TaskStatus `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
