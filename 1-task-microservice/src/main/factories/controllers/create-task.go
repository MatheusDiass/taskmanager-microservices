package factories

import (
	factories "task-microservice/src/main/factories/use-cases"
	"task-microservice/src/presentation/controllers"
)

type CreateTaskControllerFactory struct{}

func (fac CreateTaskControllerFactory) Create() controllers.CreateTaskController {
	createTaskUseCaseFactory := factories.CreateTaskUseCaseFactory{}
	return controllers.NewCreateTaskController(createTaskUseCaseFactory.Create())
}
