package usecases_test

import (
	"task-microservice/src/domain/mocks"
	usecases "task-microservice/src/domain/use-cases/update-task"
	"testing"
)

func TestUpdateTaskUseCase(t *testing.T) {
	t.Run("Task Not Found Error", func(t *testing.T) {
		updateTaskUseCase := usecases.NewUpdateTaskUseCase(
			mocks.NewFetchTaskByIdRepositoryMock(),
			mocks.NewUpdateTaskRepositoryMock())

		expectedError := "task not found"

		err := updateTaskUseCase.Execute(2, "Title", "Description", 1)

		if err == nil || err.Error() != expectedError {
			t.Errorf("Expected error message '%s', but got: %v", expectedError, err)
		}
	})
}
