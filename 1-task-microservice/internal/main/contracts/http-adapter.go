package contracts

import "github.com/matheusdiass/taskmanager-microservices/internal/main/types"

type HttpAdapter interface {
	CreateRoute(route types.Route)
}
