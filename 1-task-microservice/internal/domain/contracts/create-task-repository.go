package contracts

import "github.com/matheusdiass/taskmanager-microservices/internal/domain/entity"

type CreateTaskRepository interface {
	Execute(task entity.Task) error
}
