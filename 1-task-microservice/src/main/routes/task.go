package routes

import (
	"net/http"
	"task-microservice/src/main/types"
)

var TaskRoutes = []types.Route{
	{
		Uri:    "/tasks",
		Method: http.MethodPost,
		Function: func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Save task route!"))
		},
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
