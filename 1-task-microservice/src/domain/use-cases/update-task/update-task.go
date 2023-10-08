package usecases

import (
	"errors"
	"task-microservice/src/domain/contracts"
	"task-microservice/src/domain/entity"
)

type updateTaskUseCase struct {
	fetchTaskByIdRepository contracts.FetchTaskByIdRepository
	updateTaskRepository    contracts.UpdateTaskRepository
}

func NewUpdateTaskUseCase(
	fetchTaskByIdRepository contracts.FetchTaskByIdRepository,
	updateTaskRepository contracts.UpdateTaskRepository) updateTaskUseCase {
	return updateTaskUseCase{fetchTaskByIdRepository, updateTaskRepository}
}

func (t updateTaskUseCase) Execute(id uint, title, description string, userId uint) error {
	taskDb, err := t.fetchTaskByIdRepository.Execute(id)

	if err != nil {
		return err
	}

	//Check if the task not found
	if taskDb.Id == 0 {
		return errors.New("task not found")
	}

	task := entity.Task{
		Title:       title,
		Description: description,
		UserId:      userId,
	}

	//Entity validation
	if err = task.Validation(); err != nil {
		return err
	}

	//Update task
	if err = t.updateTaskRepository.Execute(id, task); err != nil {
		return err
	}

	return nil
}
