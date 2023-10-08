package usecases

import (
	"errors"
	"task-microservice/src/domain/contracts"
	"task-microservice/src/domain/entity"
)

type createTaskUseCase struct {
	fetchUserByIdRepository contracts.FetchUserByIdRepository
	createTaskRepository    contracts.CreateTaskRepository
}

func NewCreateTaskUseCase(
	fetchUserByIdRepository contracts.FetchUserByIdRepository,
	createTaskRepository contracts.CreateTaskRepository) createTaskUseCase {
	return createTaskUseCase{fetchUserByIdRepository, createTaskRepository}
}

func (t createTaskUseCase) Execute(title, description string, userId uint) error {
	task := entity.Task{Title: title, Description: description, UserId: userId}

	//Entity validation
	if err := task.Validation(); err != nil {
		return err
	}

	userIdDb, err := t.fetchUserByIdRepository.Execute(task.UserId)

	if err != nil {
		return err
	}

	//Check if the user not found
	if userIdDb == 0 {
		return errors.New("user not found")
	}

	//Create task
	if err = t.createTaskRepository.Execute(task); err != nil {
		return err
	}

	return nil
}
