package types

type SignUpRequest struct {
	Username string `json:"name"`
	Password string `json:"password"`
}

type ServerMessage struct {
	Message string
}
