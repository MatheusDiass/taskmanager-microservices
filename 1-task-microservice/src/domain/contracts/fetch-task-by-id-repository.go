package contracts

import (
	"task-microservice/src/domain/dtos"
)

type FetchTaskByIdRepository interface {
	Execute(id uint) (dtos.Task, error)
}
