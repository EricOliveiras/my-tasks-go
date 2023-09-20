package requests

type BasicAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	BasicAuth
}
