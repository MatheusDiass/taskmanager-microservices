package factories

import (
	usecases "github.com/matheusdiass/taskmanager-microservices/internal/domain/use-cases/create-task"
	"github.com/matheusdiass/taskmanager-microservices/internal/infra/adapters"
	"github.com/matheusdiass/taskmanager-microservices/internal/infra/repositories"
)

type CreateTaskUseCaseFactory struct{}

func (fac CreateTaskUseCaseFactory) Create() usecases.CreateTaskUseCase {
	databaseAdapter := adapters.NewGormAdapter()
	createTaskRepository := repositories.NewCreateTaskRepository(databaseAdapter)
	fetchUserByIdRepository := repositories.NewGetUserByIdRepository(databaseAdapter)
	return usecases.NewCreateTaskUseCase(fetchUserByIdRepository, createTaskRepository)
}
