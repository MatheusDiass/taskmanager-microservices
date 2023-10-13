package adapters

import (
	"log"
	"net/http"
	"task-microservice/src/main/types"

	"github.com/go-chi/chi/v5"
)

type ChiAdapter struct {
	router *chi.Mux
}

func NewChiAdapter() ChiAdapter {
	return ChiAdapter{chi.NewRouter()}
}

func (adapter ChiAdapter) CreateRoute(route types.Route) {
	switch route.Method {
	case http.MethodGet:
		adapter.router.Get(route.Uri, route.Function)

	case http.MethodPost:
		adapter.router.Post(route.Uri, route.Function)

	case http.MethodPut:
		adapter.router.Put(route.Uri, route.Function)

	case http.MethodDelete:
		adapter.router.Delete(route.Uri, route.Function)

	default:
		log.Fatal("Invalid http method!")
	}
}

func (adapter ChiAdapter) GetRouter() *chi.Mux {
	return adapter.router
}
