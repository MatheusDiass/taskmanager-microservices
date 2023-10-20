package usecases

import (
	"errors"

	"github.com/matheusdiass/taskmanager-microservices/internal/domain/contracts"
	"github.com/matheusdiass/taskmanager-microservices/internal/domain/entity"
)

type CreateTaskUseCase struct {
	fetchUserByIdRepository contracts.FetchUserByIdRepository
	createTaskRepository    contracts.CreateTaskRepository
}

func NewCreateTaskUseCase(
	fetchUserByIdRepository contracts.FetchUserByIdRepository,
	createTaskRepository contracts.CreateTaskRepository) CreateTaskUseCase {
	return CreateTaskUseCase{fetchUserByIdRepository, createTaskRepository}
}

func (t CreateTaskUseCase) Execute(title, description string, userId uint) error {
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
