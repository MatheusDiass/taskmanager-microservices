package contracts

import "github.com/matheusdiass/taskmanager-microservices/internal/domain/entity"

type UpdateTaskRepository interface {
	Execute(id uint, task entity.Task) error
}
