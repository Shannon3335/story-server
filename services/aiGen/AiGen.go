package aigen

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shannon3335/story-server/internal/dto"
	"github.com/shannon3335/story-server/internal/types"
)

type AiHandler struct {
	AiService AiService
}

func NewAiHandler(AiService AiService) *AiHandler {
	return &AiHandler{
		AiService: AiService,
	}
}

func (a *AiHandler) TestHello(c echo.Context) error {
	var payload dto.StartStoryPayload
	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, types.ServerMessage{Message: "Invalid data for starting a story"})
	}
	str, err := a.AiService.StartStory((*types.StartStoryPrompt)(&payload))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, types.ServerMessage{Message: str})
}
