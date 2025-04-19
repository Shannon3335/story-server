package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/shannon3335/story-server/internal/types"
)

func TestSignup(t *testing.T) {
	service := &mockUserStore{}
	handler := NewAuthHandler(service)

	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		payload := types.User{
			FirstName: "Shannon",
			LastName:  "Fernandes",
			Email:     "shannon",
			Password:  "sdfg",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		e := echo.New()
		e.POST("signup", handler.Signup)
		e.ServeHTTP(rr, req)

		if rr.Code != http.StatusUnprocessableEntity {
			t.Fail()
		}
	})
}

type mockUserStore struct {
}

func (m *mockUserStore) SignupUser(types.User) error {
	return nil
}
func (m *mockUserStore) Login(Username string, Password string) (bool, error) {
	return true, nil
}
func (m *mockUserStore) GetUser(Email string) (*types.User, error) {
	return nil, nil
}
