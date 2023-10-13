package main

import (
	"fmt"
	"net/http"
	"task-microservice/src/main/adapters"
	"task-microservice/src/main/config"
	"task-microservice/src/main/factories"
	"task-microservice/src/main/types"
)

func init() {
	config.SetupEnvironmentVariables()
}

func main() {
	runMessage := fmt.Sprintf("🚀 task service is running at http://localhost:%d", config.Port)

	httpAdapter := adapters.NewChiAdapter()
	httpAdapter.CreateRoute(types.Route{
		Uri:    "/",
		Method: http.MethodGet,
		Function: func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(runMessage))
		},
	})

	//Create routes
	factories.BuildRoutes(httpAdapter)

	fmt.Println(runMessage)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), httpAdapter.GetRouter())
}
