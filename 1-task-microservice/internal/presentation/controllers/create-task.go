package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/matheusdiass/taskmanager-microservices/internal/domain/entity"
	usecases "github.com/matheusdiass/taskmanager-microservices/internal/domain/use-cases/create-task"
	"github.com/matheusdiass/taskmanager-microservices/internal/presentation/helpers"
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
		helpers.ErrorResponse(w, http.StatusUnprocessableEntity, err)
	}

	var taskRequest entity.Task
	if err = json.Unmarshal(body, &taskRequest); err != nil {
		helpers.ErrorResponse(w, http.StatusBadRequest, err)
	}

	if err = controller.createTaskUseCase.Execute(
		taskRequest.Title,
		taskRequest.Description,
		taskRequest.UserId); err != nil {
		helpers.ErrorResponse(w, http.StatusBadRequest, err)
	}
}
