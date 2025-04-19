package dto

type SignupPayload struct {
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,alphanum"`
}

type StartStoryPayload struct {
	MainCharacter string `json:"mainCharacter" validate:"required"`
	Villain       string `json:"villain"`
	Setting       string `json:"setting"`
}
