package usecases_test

import (
	"task-microservice/src/domain/mocks"
	usecases "task-microservice/src/domain/use-cases/delete-task"
	"testing"
)

func TestDeleteTaskUseCase(t *testing.T) {
	t.Run("Task Not Found Error", func(t *testing.T) {
		deleteTaskUseCase := usecases.NewDeleteTaskUseCase(
			mocks.NewFetchTaskByIdRepositoryMock(),
			mocks.NewDeleteTaskRepository())

		expectedError := "task not found"

		err := deleteTaskUseCase.Execute(2)

		if err == nil || err.Error() != expectedError {
			t.Errorf("Expected error message '%s', but got: %v", expectedError, err)
		}
	})
}
