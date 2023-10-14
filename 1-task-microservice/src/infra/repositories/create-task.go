package repositories

import (
	"task-microservice/src/domain/entity"
	"task-microservice/src/infra/adapters"
)

type CreateTaskRepository struct {
	databaseAdapter adapters.GormAdapter
}

func NewCreateTaskRepository(
	databaseAdapter adapters.GormAdapter,
) CreateTaskRepository {
	return CreateTaskRepository{databaseAdapter}
}

func (repository CreateTaskRepository) Execute(task entity.Task) error {
	result := repository.databaseAdapter.GetDb().Table("sch_task.tbl_task").Create(&task)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
