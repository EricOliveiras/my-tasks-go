package requests

type BasicAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	BasicAuth
}

type CreateUserRequest struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BasicAuth
}
