package delivery

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/SiwakornSitti/practice-ai/be-agent/internal/user/domain"
	"github.com/golang-jwt/jwt/v5"
)

type userHandler struct {
	usecase   domain.UserUsecase
	jwtSecret string
}

// NewUserHandler initializes the user HTTP routes
func NewUserHandler(mux *http.ServeMux, usecase domain.UserUsecase, secret string) {
	handler := &userHandler{
		usecase:   usecase,
		jwtSecret: secret,
	}

	// Public routes
	mux.HandleFunc("POST /register", handler.Register)
	mux.HandleFunc("POST /login", handler.Login)

	// Protected routes
	mux.HandleFunc("GET /users/me", handler.authMiddleware(handler.GetProfile))
	mux.HandleFunc("PUT /users/me", handler.authMiddleware(handler.UpdateProfile))
	mux.HandleFunc("DELETE /users/me", handler.authMiddleware(handler.DeleteAccount))
}

// WriteJSON is a helper for JSON responses
func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// WriteError is a helper for error responses
func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

// Context key for user ID
type contextKey string

const userIDKey contextKey = "userID"

func (h *userHandler) authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			WriteError(w, http.StatusUnauthorized, domain.ErrUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			WriteError(w, http.StatusUnauthorized, domain.ErrUnauthorized)
			return
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(h.jwtSecret), nil
		})

		if err != nil || !token.Valid {
			WriteError(w, http.StatusUnauthorized, domain.ErrUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			WriteError(w, http.StatusUnauthorized, domain.ErrUnauthorized)
			return
		}

		userID, ok := claims["sub"].(string)
		if !ok {
			WriteError(w, http.StatusUnauthorized, domain.ErrUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (h *userHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req domain.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteError(w, http.StatusBadRequest, errors.New("invalid request body"))
		return
	}

	res, err := h.usecase.Register(r.Context(), &req)
	if err != nil {
		if err == domain.ErrEmailAlreadyInUse {
			WriteError(w, http.StatusConflict, err)
			return
		}
		if err == domain.ErrInvalidCredentials {
			WriteError(w, http.StatusBadRequest, err)
			return
		}
		WriteError(w, http.StatusInternalServerError, err)
		return
	}

	WriteJSON(w, http.StatusCreated, res)
}

func (h *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req domain.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteError(w, http.StatusBadRequest, errors.New("invalid request body"))
		return
	}

	res, err := h.usecase.Login(r.Context(), &req)
	if err != nil {
		if err == domain.ErrInvalidCredentials {
			WriteError(w, http.StatusUnauthorized, err)
			return
		}
		WriteError(w, http.StatusInternalServerError, err)
		return
	}

	WriteJSON(w, http.StatusOK, res)
}

func (h *userHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(userIDKey).(string)

	user, err := h.usecase.GetProfile(r.Context(), userID)
	if err != nil {
		if err == domain.ErrUserNotFound {
			WriteError(w, http.StatusNotFound, err)
			return
		}
		WriteError(w, http.StatusInternalServerError, err)
		return
	}

	WriteJSON(w, http.StatusOK, user)
}

func (h *userHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(userIDKey).(string)

	var req struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteError(w, http.StatusBadRequest, errors.New("invalid request body"))
		return
	}

	user, err := h.usecase.UpdateProfile(r.Context(), userID, req.Name)
	if err != nil {
		if err == domain.ErrUserNotFound {
			WriteError(w, http.StatusNotFound, err)
			return
		}
		WriteError(w, http.StatusInternalServerError, err)
		return
	}

	WriteJSON(w, http.StatusOK, user)
}

func (h *userHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(userIDKey).(string)

	if err := h.usecase.DeleteAccount(r.Context(), userID); err != nil {
		if err == domain.ErrUserNotFound {
			WriteError(w, http.StatusNotFound, err)
			return
		}
		WriteError(w, http.StatusInternalServerError, err)
		return
	}

	WriteJSON(w, http.StatusNoContent, nil)
}
