package auth

import (
	"encoding/json"
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/pkg/res"
	"net/http"
	"regexp"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload LoginRequest
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			res.Json(w, err.Error(), 402)
			return
		}

		if payload.Email == "" {
			res.Json(w, "Email required", 402)
			return
		}

		reg, _ := regexp.Compile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)

		if !reg.MatchString(payload.Email) {
			res.Json(w, "Not email", 402)
			return
		}

		if payload.Password == "" {
			res.Json(w, "Password required", 402)
			return
		}

		data := LoginResponse{
			Token: "1234",
		}
		res.Json(w, data, 200)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Register")
	}
}
