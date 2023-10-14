package factories

import (
	usecases "task-microservice/src/domain/use-cases/create-task"
	"task-microservice/src/infra/adapters"
	"task-microservice/src/infra/repositories"
)

type CreateTaskUseCaseFactory struct{}

func (fac CreateTaskUseCaseFactory) Create() usecases.CreateTaskUseCase {
	databaseAdapter := adapters.NewGormAdapter()
	createTaskRepository := repositories.NewCreateTaskRepository(databaseAdapter)
	fetchUserByIdRepository := repositories.NewGetUserByIdRepository(databaseAdapter)
	return usecases.NewCreateTaskUseCase(fetchUserByIdRepository, createTaskRepository)
}
