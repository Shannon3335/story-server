package aigen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/shannon3335/story-server/internal/dto"
	"github.com/shannon3335/story-server/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestGenerateStory(t *testing.T) {
	testAIService := &mockAiGenService{}
	handler := NewAiHandler(testAIService)

	t.Run("should fail if connection to AI can't be established", func(t *testing.T) {
		e := echo.New()
		e.POST("/generate", handler.TestHello)

		mockData := &dto.StartStoryPayload{
			MainCharacter: "Arthur Pendragon",
			Villain:       "Moriarty",
			Setting:       "Ancient Roman",
		}
		marshalledData, _ := json.Marshal(mockData)
		req := httptest.NewRequest(http.MethodPost, "/generate", bytes.NewReader(marshalledData))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rr := httptest.NewRecorder()

		e.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		expectedOutput := fmt.Sprintf("Your main character is %s, villain is %s. They will battle in the setting %s", mockData.MainCharacter, mockData.Villain, mockData.Setting)
		var response types.ServerMessage
		if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
			t.Fatalf("Failed to unmarshal JSON response: %v", err)
		}
		assert.Equal(t, expectedOutput, response.Message)

	})
}

type mockAiGenService struct{}

func (a *mockAiGenService) StartStory(promptDetails *types.StartStoryPrompt) (string, error) {
	return fmt.Sprintf("Your main character is %s, villain is %s. They will battle in the setting %s", promptDetails.MainCharacter, promptDetails.Villain, promptDetails.Setting), nil
}
