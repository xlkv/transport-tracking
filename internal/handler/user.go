package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"tracking.xlkv.com/internal/domain"
	"tracking.xlkv.com/internal/response"
	"tracking.xlkv.com/internal/service"
)

type UserHandler struct {
	service     *service.UserService
	authService *service.AuthService
}

func NewUserHandler(service service.UserService, authService service.AuthService) *UserHandler {
	return &UserHandler{
		service:     &service,
		authService: &authService,
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {

	var req struct {
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		response.Error(w, http.StatusBadRequest, "bad request")
		return
	}

	driver, err := h.service.Register(r.Context(), req.Username, req.Name, req.Password)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrAlreadyExists):
			response.Error(w, http.StatusConflict, "username already taken")
		case errors.Is(err, domain.ErrValidation):
			response.Error(w, http.StatusBadRequest, err.Error())
		default:
			response.Error(w, http.StatusInternalServerError, "server error")
		}
		return
	}

	accessToken, refreshToken, err := h.authService.GenereateTokens(r.Context(), driver.ID)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, "server error")
		return
	}

	response.Success(w, http.StatusCreated, map[string]interface{}{
		"driver":        driver,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		response.Error(w, http.StatusBadRequest, "bad request")
		return
	}

	driver, err := h.service.Login(r.Context(), req.Username, req.Password)

	if err != nil {

		if errors.Is(err, domain.ErrNotFound) || errors.Is(err, domain.ErrUnauthorized) {
			response.Error(w, http.StatusUnauthorized, "username or password is ivalid")
			return
		}

		response.Error(w, http.StatusInternalServerError, "server error")
		return
	}

	accessToken, refreshToken, err := h.authService.GenereateTokens(r.Context(), driver.ID)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, "server error")
		return
	}

	response.Success(
		w,
		http.StatusOK,
		map[string]interface{}{
			"driver":        driver,
			"refresh_token": refreshToken,
			"access_token":  accessToken,
		},
	)
}

func (h *UserHandler) Refresh(w http.ResponseWriter, r *http.Request) {

	var req struct {
		RefreshToken string `json:"refresh_token"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	if req.RefreshToken == "" {
		response.Error(w, http.StatusBadRequest, "token is empty")
		return
	}

	accessToken, err := h.authService.Refresh(r.Context(), req.RefreshToken)

	if err != nil {
		if errors.Is(err, domain.ErrUnauthorized) {
			response.Error(w, http.StatusUnauthorized, "invalid token")
			return
		}
		response.Error(w, http.StatusInternalServerError, "server error")
		return
	}

	response.Success(w, http.StatusOK, map[string]interface{}{
		"access_token": accessToken,
	})
}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {

	var req struct {
		RefreshToken string `json:"refresh_token"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	if req.RefreshToken == "" {
		response.Error(w, http.StatusBadRequest, "token is empty")
		return
	}

	err := h.authService.Logout(r.Context(), req.RefreshToken)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, "server error")
		return
	}

	response.Success(w, http.StatusOK, "done")
}
