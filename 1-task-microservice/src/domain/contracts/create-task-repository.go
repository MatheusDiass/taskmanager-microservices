package contracts

import "task-microservice/src/domain/entity"

type CreateTaskRepository interface {
	Execute(task entity.Task) error
}
