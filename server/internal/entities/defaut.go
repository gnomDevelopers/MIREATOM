package entities

type Message struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
