package routes

import (
	"net/http"
	factories "task-microservice/src/main/factories/controllers"
	"task-microservice/src/main/types"
)

var (
	createTaskController = factories.CreateTaskControllerFactory{}.Create()
)

var TaskRoutes = []types.Route{
	{
		Uri:      "/tasks",
		Method:   http.MethodPost,
		Function: createTaskController.Handle,
	},
	{
		Uri:    "/tasks/{id}",
		Method: http.MethodGet,
		Function: func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Get task by id route!"))
		},
	},
	{
		Uri:    "/tasks",
		Method: http.MethodGet,
		Function: func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Get all tasks route!"))
		},
	},
	{
		Uri:    "/tasks/{id}",
		Method: http.MethodPut,
		Function: func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Update task route!"))
		},
	},
	{
		Uri:    "/tasks/{id}",
		Method: http.MethodDelete,
		Function: func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Delete task route!"))
		},
	},
}
