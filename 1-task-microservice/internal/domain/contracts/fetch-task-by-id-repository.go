package contracts

import (
	"github.com/matheusdiass/taskmanager-microservices/internal/domain/dtos"
)

type FetchTaskByIdRepository interface {
	Execute(id uint) (dtos.Task, error)
}
