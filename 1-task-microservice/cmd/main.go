package main

import (
	"fmt"
	"net/http"

	config "github.com/matheusdiass/taskmanager-microservices/configs"
	"github.com/matheusdiass/taskmanager-microservices/internal/main/adapters"
	"github.com/matheusdiass/taskmanager-microservices/internal/main/factories"
	"github.com/matheusdiass/taskmanager-microservices/internal/main/types"
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
