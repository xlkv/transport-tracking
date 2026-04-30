package handler

import (
	"net/http"

	"tracking.xlkv.com/internal/response"
	"tracking.xlkv.com/internal/serivce"
)

type StatusHandler struct {
	service *serivce.StatusService
}

func NewStatusHandler(service *serivce.StatusService) *StatusHandler {
	return &StatusHandler{
		service: service,
	}
}

func (s *StatusHandler) Status(w http.ResponseWriter, r *http.Request) {

	response.WriteJSON(w, r, http.StatusOK, map[string]string{"status": "ok"})

}
