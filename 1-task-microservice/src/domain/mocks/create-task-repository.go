package mocks

import "task-microservice/src/domain/entity"

type createTaskRepository struct {
}

func NewCreateTaskRepository() createTaskRepository {
	return createTaskRepository{}
}

func (r createTaskRepository) Execute(task entity.Task) error {
	return nil
}
