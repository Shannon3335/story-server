package types

type StartStoryPrompt struct {
	MainCharacter string
	Villain       string
	Setting       string
}

type ContinueStoryPrompt struct {
	StartStoryPrompt
	StoryOptions []string
}

type SavedStoryDetails struct {
}
