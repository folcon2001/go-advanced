package auth

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResonse struct {
	Token string `json:"token"`
}
