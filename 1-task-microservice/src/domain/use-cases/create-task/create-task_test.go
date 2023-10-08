package usecases_test

import (
	"task-microservice/src/domain/mocks"
	usecases "task-microservice/src/domain/use-cases/create-task"
	"testing"
)

func TestCreateTaskUseCase(t *testing.T) {
	t.Run("User Not Found Error", func(t *testing.T) {
		createTaskUseCase := usecases.NewCreateTaskUseCase(
			mocks.NewFetchUserByIdRepositoryMock(),
			mocks.NewCreateTaskRepository())

		expectedError := "user not found"

		err := createTaskUseCase.Execute("Some title", "Some description", 2)

		if err == nil || err.Error() != expectedError {
			t.Errorf("Expected error message '%s', but got: %v", expectedError, err)
		}
	})
}
