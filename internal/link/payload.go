package link

type CreateResponse struct {
	Token string `json:"token"`
}

type CreateRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
