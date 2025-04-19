package auth

import "golang.org/x/crypto/bcrypt"

const COST = 10

func HashPassword(Password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(Password), COST)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CompareWithString(Password string, HashedPassword []byte) bool {
	result := bcrypt.CompareHashAndPassword(HashedPassword, []byte(Password))
	if result != nil && result == bcrypt.ErrMismatchedHashAndPassword {
		return false
	}
	return true
}
