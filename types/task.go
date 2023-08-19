package types

import "time"

type Task struct {
	ID          int64
	Description string
	Completed   bool
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
