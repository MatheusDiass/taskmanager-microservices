package mocks

import "task-microservice/src/domain/entity"

type updateTaskRepository struct{}

func NewUpdateTaskRepositoryMock() updateTaskRepository {
	return updateTaskRepository{}
}

func (r updateTaskRepository) Execute(id uint, task entity.Task) error {
	return nil
}
