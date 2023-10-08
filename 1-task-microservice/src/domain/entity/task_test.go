package entity

import (
	"testing"
)

func TestTaskValidation(t *testing.T) {
	t.Run("Invalid Title (Too long)", func(t *testing.T) {
		task := Task{
			Title:       "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			Description: "Some description",
			UserId:      1,
		}

		err := task.Validation()
		expectedError := "the title cannot be longer than 30 characters"

		if err == nil || err.Error() != expectedError {
			t.Errorf("Expected error message '%s', but got: %v", expectedError, err)
		}
	})

	t.Run("Empty Title", func(t *testing.T) {
		task := Task{
			Title:       "",
			Description: "Some description",
			UserId:      1,
		}

		err := task.Validation()
		expectedError := "the title cannot be empty"

		if err == nil || err.Error() != expectedError {
			t.Errorf("Expected error message '%s', but got: %v", expectedError, err)
		}
	})

	t.Run("Invalid Description (Too long)", func(t *testing.T) {
		task := Task{
			Title:       "Some title",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla sagittis tristique odio, sed tincidunt augue luctus eget. Cras congue quam lacus, vitae rhoncus eros suscipit at. Phasellus metus eros, condimentum id metus ac, ornare suscipit nisi. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas.",
			UserId:      1,
		}

		err := task.Validation()
		expectedError := "the description cannot be longer than 300 characters"

		if err == nil || err.Error() != expectedError {
			t.Errorf("Expected error message '%s', but got: %v", expectedError, err)
		}
	})

	t.Run("Empty Description", func(t *testing.T) {
		task := Task{
			Title:       "Some title",
			Description: "",
			UserId:      1,
		}

		err := task.Validation()
		expectedError := "the description cannot be empty"

		if err == nil || err.Error() != expectedError {
			t.Errorf("Expected error message '%s', but got: %v", expectedError, err)
		}
	})

	t.Run("Invalid user id", func(t *testing.T) {
		task := Task{
			Title:       "Some title",
			Description: "Some description",
			UserId:      0,
		}

		err := task.Validation()
		expectedError := "the user id has an invalid value"

		if err == nil || err.Error() != expectedError {
			t.Errorf("Expected error message '%s', but got: %v", expectedError, err)
		}
	})
}
