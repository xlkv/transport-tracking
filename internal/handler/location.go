package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"tracking.xlkv.com/internal/domain"
	"tracking.xlkv.com/internal/response"
	"tracking.xlkv.com/internal/service"
)

type LocationHandler struct {
	service service.LocationService
}

func NewLocationHandler(service service.LocationService) *LocationHandler {
	return &LocationHandler{
		service: service,
	}
}

func (s *LocationHandler) Create(w http.ResponseWriter, r *http.Request) {

	var req struct {
		Lat    float64 `json:"lat"`
		Lng    float64 `json:"lng"`
		TripID int64   `json:"trip_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		response.Error(w, http.StatusBadRequest, "bad request")
		return
	}

	location, err := s.service.Create(r.Context(), req.Lat, req.Lng, req.TripID)

	if err != nil {
		if errors.Is(err, domain.ErrInvalidParam) {
			response.Error(w, http.StatusBadRequest, "invalid trip id")
			return
		}
		response.Error(w, http.StatusInternalServerError, "server error")
		return
	}

	response.Success(w, http.StatusCreated, map[string]interface{}{
		"location": location,
	})
}
