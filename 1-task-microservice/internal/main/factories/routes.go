package factories

import (
	"github.com/matheusdiass/taskmanager-microservices/internal/main/contracts"
	"github.com/matheusdiass/taskmanager-microservices/internal/main/routes"
	"github.com/matheusdiass/taskmanager-microservices/internal/main/types"
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
