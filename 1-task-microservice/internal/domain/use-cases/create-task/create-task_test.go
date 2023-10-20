package usecases_test

import (
	"testing"

	"github.com/matheusdiass/taskmanager-microservices/internal/domain/mocks"
	usecases "github.com/matheusdiass/taskmanager-microservices/internal/domain/use-cases/create-task"
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
