package types

type LoginDetails struct {
	Username string `json:"username" gorm:"uniqueIndex"`
	Password string `json:"password"`
}

func NewLoginDetails(Username string, Password string) *LoginDetails {
	return &LoginDetails{
		Username: Username,
		Password: Password,
	}
}
