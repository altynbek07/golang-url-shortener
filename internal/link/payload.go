package link

type LinkCreateResponse struct {
	Token string `json:"token"`
}

type LinkCreateRequest struct {
	Url string `json:"url" validate:"required,url"`
}
