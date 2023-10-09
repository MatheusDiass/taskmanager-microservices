package usecases

import (
	"errors"
	"task-microservice/src/domain/contracts"
)

type deleteTaskUseCase struct {
	fetchTaskByIdRepository contracts.FetchTaskByIdRepository
	deleteTaskRepository    contracts.DeleteTaskRepository
}

func NewDeleteTaskUseCase(
	fetchTaskByIdRepository contracts.FetchTaskByIdRepository,
	deleteTaskRepository contracts.DeleteTaskRepository) deleteTaskUseCase {
	return deleteTaskUseCase{fetchTaskByIdRepository, deleteTaskRepository}
}

func (t deleteTaskUseCase) Execute(id uint) error {
	taskDb, err := t.fetchTaskByIdRepository.Execute(id)

	if err != nil {
		return err
	}

	//Check if the task not found
	if taskDb.Id == 0 {
		return errors.New("task not found")
	}

	//Delete task
	if err = t.deleteTaskRepository.Execute(id); err != nil {
		return err
	}

	return nil
}
