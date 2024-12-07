package entities

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type CreateUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type CreateUserResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ID           int    `json:"id"`
}

type LoginUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ID           int    `json:"id"`
}
