package contracts

import "task-microservice/src/main/types"

type HttpAdapter interface {
	CreateRoute(route types.Route)
}
