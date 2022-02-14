package http_server

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"l0/internal/cache"
	"net/http"
)

type Server struct {
	cache *cache.LocalCache
}

func InitServer(cache *cache.LocalCache) *Server {
	return &Server{
		cache: cache,
	}
}

func (s *Server) Start() error {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/api/v1/order/{order_uid}", s.handleGetOrder).Methods("GET")

	return http.ListenAndServe("127.0.0.1:8080", handlers.CORS()(r))
}
