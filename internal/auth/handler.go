package auth

import (
	"go/adv-demo/configs"
	"go/adv-demo/pkg/req"
	"go/adv-demo/pkg/res"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

type AuthHandler struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](w, r)

		if err != nil {
			return
		}

		_, err = handler.AuthService.Login(body.Email, body.Password)
		if err != nil {
			res.Json(w, err.Error(), 402)
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
		body, err := req.HandleBody[RegistrationRequest](w, r)

		if err != nil {
			return
		}

		handler.AuthService.Register(body.Email, body.Password, body.Name)

		res.Json(w, nil, 200)
	}
}
