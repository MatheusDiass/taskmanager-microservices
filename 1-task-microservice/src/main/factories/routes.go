package factories

import (
	"task-microservice/src/main/contracts"
	"task-microservice/src/main/routes"
	"task-microservice/src/main/types"
)

func BuildRoutes(httpAdapter contracts.HttpAdapter) {
	for _, route := range routes.TaskRoutes {
		httpAdapter.CreateRoute(types.Route{
			Uri:      route.Uri,
			Method:   route.Method,
			Function: route.Function,
		})
	}
}
