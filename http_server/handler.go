package http_server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func (s *Server) handleGetOrder(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	id := vars["order_uid"]

	order := s.cache.Get(id)

	err := json.NewEncoder(w).Encode(order)
	if err != nil {
		errorRes(w, http.StatusBadGateway, err.Error())
		return
	}
}

func errorRes(w http.ResponseWriter, statusCode int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if msg != "" {
		_ = json.NewEncoder(w).Encode(msg)
	}
}
