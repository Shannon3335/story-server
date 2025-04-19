package auth

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shannon3335/story-server/internal/dto"
	"github.com/shannon3335/story-server/internal/types"
	"github.com/shannon3335/story-server/utils"
)

type AuthHandler struct {
	AuthService AuthService
}

func (a *AuthHandler) Signup(c echo.Context) error {
	var payload dto.SignupPayload
	if err := c.Bind(&payload); err != nil {
		// return c.JSON(http.StatusUnprocessableEntity, types.ServerMessage{Message: "Invalid input structure/values"})
		return c.JSON(http.StatusUnprocessableEntity, types.ServerMessage{Message: err.Error()})
	}
	if err := utils.ValidateStruct(payload); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, types.ServerMessage{Message: err.Error()})
	}
	// check if user exists in the database
	result, err := a.AuthService.GetUser(payload.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	} else if result != nil {
		return c.JSON(http.StatusConflict, types.ServerMessage{Message: "Account with email exists"})
	}
	// create user if not
	u := types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  payload.Password,
	}
	err = a.AuthService.SignupUser(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.ServerMessage{Message: "Error signing up user"})
	}
	return c.JSON(http.StatusOK, types.ServerMessage{Message: "User Signed up"})
}

func (a *AuthHandler) Login(c echo.Context) error {
	var userDetails types.LoginDetails
	if err := c.Bind(userDetails); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, types.ServerMessage{Message: "Invalid input structure/values"})
	}
	if err := utils.ValidateStruct(userDetails); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, types.ServerMessage{Message: err.Error()})
	}
	loggedIn, err := a.AuthService.Login(userDetails.Username, userDetails.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.ServerMessage{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, types.ServerMessage{Message: fmt.Sprintf("Login %t", loggedIn)})
}

func NewAuthHandler(authService AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}
