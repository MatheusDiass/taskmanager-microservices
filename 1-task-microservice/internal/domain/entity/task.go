package entity

import (
	"errors"
)

type Task struct {
	Title       string
	Description string
	UserId      uint
}

func (t Task) Validation() error {
	if len(t.Title) > 30 {
		return errors.New("the title cannot be longer than 30 characters")
	}

	if t.Title == "" {
		return errors.New("the title cannot be empty")
	}

	if len(t.Description) > 300 {
		return errors.New("the description cannot be longer than 300 characters")
	}

	if t.Description == "" {
		return errors.New("the description cannot be empty")
	}

	if t.UserId == 0 {
		return errors.New("the user id has an invalid value")
	}

	return nil
}
