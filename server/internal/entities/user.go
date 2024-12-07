package entities

type User struct {
	ID        int    `json:"id" db:"id"`
	Password  string `json:"password" db:"password"`
	Email     string `json:"email" db:"email"`
	Role      string `json:"role" db:"role"`
	Name      string `json:"name" db:"name"`
	Surname   string `json:"surname" db:"surname"`
	ThirdName string `json:"third_name" db:"third_name"`
}

type CreateUserRequest struct {
	Password  string `json:"password"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	ThirdName string `json:"third_name"`
}

type CreateUserResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ID           int    `json:"id"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ID           int    `json:"id"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
