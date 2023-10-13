package main

import (
	"fmt"
	"net/http"
	"task-microservice/src/main/adapters"
	"task-microservice/src/main/factories"
	"task-microservice/src/main/types"
)

func main() {
	runMessage := "🚀 task service is running at http://localhost:5000"

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
	http.ListenAndServe(":5000", httpAdapter.GetRouter())
}
