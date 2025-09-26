package domain

import (
	//"fmt"
	"time"
	//"errors"
)

type Task struct {
	id          int
	description string
	status      string
	createdAt   time.Time
	updatedAt   time.Time
}

func (t *Task) GetID() int {
	return t.id
}

func (t *Task) GetDescription() string {
	return t.description
}

func (t *Task) GetStatus() string {
	return t.status
}

func (t *Task) GetCreatedAt() time.Time {
	return t.createdAt
}

func (t *Task) GetUpdatedAt() time.Time {
	return t.updatedAt
}
func (t *Task) SetDescription(description string) {
	t.description = description
}

func (t *Task) SetID(id int) {
	t.id = id
}
func (t *Task) SetStatus(status string) error {
	if status != "to-do" && status != "in-progress" && status != "done" {
		return ErrWrongStatus
	}
	t.status = status
	return nil
}

func (t *Task) SetCreatedAt(createdAt time.Time) {
	t.createdAt = createdAt
}

func (t *Task) SetUpdatedAt(updatedAt time.Time) {
	t.updatedAt = updatedAt
}
