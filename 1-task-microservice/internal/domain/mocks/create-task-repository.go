package mocks

import "github.com/matheusdiass/taskmanager-microservices/internal/domain/entity"

type createTaskRepository struct{}

func NewCreateTaskRepository() createTaskRepository {
	return createTaskRepository{}
}

func (r createTaskRepository) Execute(task entity.Task) error {
	return nil
}
