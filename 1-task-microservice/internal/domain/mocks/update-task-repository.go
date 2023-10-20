package mocks

import "github.com/matheusdiass/taskmanager-microservices/internal/domain/entity"

type updateTaskRepository struct{}

func NewUpdateTaskRepositoryMock() updateTaskRepository {
	return updateTaskRepository{}
}

func (r updateTaskRepository) Execute(id uint, task entity.Task) error {
	return nil
}
