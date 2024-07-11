package models

import "time"

type Task struct {
	ID        int
	Name      string
	Status    string
	CreaednAt time.Time
	UpdatedAt time.Time
}
