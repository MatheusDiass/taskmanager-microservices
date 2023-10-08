package contracts

import "task-microservice/src/domain/entity"

type UpdateTaskRepository interface {
	Execute(id uint, task entity.Task) error
}
