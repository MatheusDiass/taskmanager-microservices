package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"task-microservice/src/domain/entity"
	usecases "task-microservice/src/domain/use-cases/create-task"
)

type CreateTaskController struct {
	createTaskUseCase usecases.CreateTaskUseCase
}

func NewCreateTaskController(createTaskUseCase usecases.CreateTaskUseCase) CreateTaskController {
	return CreateTaskController{createTaskUseCase}
}

func (controller CreateTaskController) Handle(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var taskRequest entity.Task
	if err = json.Unmarshal(body, &taskRequest); err != nil {
		log.Fatal(err)
	}

	if err = controller.createTaskUseCase.Execute(
		taskRequest.Title,
		taskRequest.Description,
		taskRequest.UserId); err != nil {
		log.Fatal(err)
	}
}
