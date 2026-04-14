package agent

import (
	"context"

	"google.golang.org/adk/model"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/genai"
)

// Just the model itself
type AgentModel struct {
	Config *AgentConfig
	Ctx context.Context
	Model model.LLM
}

func NewAgentModel(ctx context.Context, config *AgentConfig) (*AgentModel, error) {
	model, err := gemini.NewModel(ctx, config.AgentModel, &genai.ClientConfig{
		APIKey: config.GoogleAPIKey,
	})

	if err != nil {
		return nil, err
	}
	return &AgentModel{
		Config: config,
		Ctx: ctx,
		Model: model,
	}, nil
}