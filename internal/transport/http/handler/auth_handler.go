package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"tasker_go/internal/auth"
	"tasker_go/internal/repository"
	"tasker_go/internal/service"
	"tasker_go/internal/transport/http/dto"
	"tasker_go/internal/transport/http/responder"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, "invalid json")
		return
	}

	if err := dto.Validate.Struct(req); err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, dto.FormatValidationError(err))
		return
	}

	if err := h.authService.Register(r.Context(), &req); err != nil {
		if errors.Is(err, repository.ErrEmailAlreadyExists) {
			responder.RespondWithError(w, http.StatusConflict, err.Error())
		} else {
			responder.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	responder.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "done"})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, "invalid json")
		return
	}

	if err := dto.Validate.Struct(req); err != nil {
		responder.RespondWithError(w, http.StatusBadRequest, dto.FormatValidationError(err))
		return
	}

	token, err := h.authService.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		if errors.Is(err, auth.ErrInvalidCredentials) {
			responder.RespondWithError(w, http.StatusUnauthorized, err.Error())
		} else {
			responder.RespondWithError(w, http.StatusInternalServerError, "internal error")
		}
		return
	}

	responder.RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})
}
