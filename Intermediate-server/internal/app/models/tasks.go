package models

import "time"

type Task struct {
	ID        int
	Status    string
	CreaednAt time.Time
	UpdatedAt time.Time
}
