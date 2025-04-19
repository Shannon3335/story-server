package aigen

import (
	"fmt"

	"github.com/shannon3335/story-server/internal/types"
)

type AiService interface {
	StartStory(*types.StartStoryPrompt) (string, error)
}

type aiService struct {
	apiKey string
}

func NewAiService(apiKey string) AiService {
	return &aiService{
		apiKey: apiKey,
	}
}

func (a *aiService) StartStory(promptDetails *types.StartStoryPrompt) (string, error) {
	return fmt.Sprintf("Your main character is %s, villain is %s. They will battle in the setting %s", promptDetails.MainCharacter, promptDetails.Villain, promptDetails.Setting), nil
}
