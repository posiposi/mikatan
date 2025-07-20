package request

type SignUpRequest struct {
	Name     string
	Email    string
	Password string
}

type LogInRequest struct {
	Email    string
	Password string
}