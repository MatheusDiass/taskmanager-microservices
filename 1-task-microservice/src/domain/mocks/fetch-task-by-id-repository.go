package mocks

import (
	"task-microservice/src/domain/dtos"
)

type fetchTaskByIdRepositoryMock struct{}

func NewFetchTaskByIdRepositoryMock() fetchTaskByIdRepositoryMock {
	return fetchTaskByIdRepositoryMock{}
}

func (t fetchTaskByIdRepositoryMock) Execute(id uint) (dtos.Task, error) {
	if id == 1 {
		return dtos.Task{
			Id:          id,
			Title:       "Title",
			Description: "Description",
			UserId:      1,
		}, nil
	}

	return dtos.Task{}, nil
}
