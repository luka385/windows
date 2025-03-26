package domain

import "time"

type Task struct {
	ID          string
	Title       string
	Description string
	Date        time.Time
	Completed   bool
}
