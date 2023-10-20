package factories

import (
	factories "github.com/matheusdiass/taskmanager-microservices/internal/main/factories/use-cases"
	"github.com/matheusdiass/taskmanager-microservices/internal/presentation/controllers"
)

type CreateTaskControllerFactory struct{}

func (fac CreateTaskControllerFactory) Create() controllers.CreateTaskController {
	createTaskUseCaseFactory := factories.CreateTaskUseCaseFactory{}
	return controllers.NewCreateTaskController(createTaskUseCaseFactory.Create())
}
