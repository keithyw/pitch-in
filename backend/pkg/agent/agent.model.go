package agent

import (
	"context"
	"log/slog"

	"google.golang.org/adk/model"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/genai"
)

// Just the model itself
type AgentModel struct {
	Config *AgentConfig
	Ctx context.Context
	Model model.LLM
	log *slog.Logger
}

func NewAgentModel(ctx context.Context, config *AgentConfig, log *slog.Logger) (*AgentModel, error) {
	model, err := gemini.NewModel(ctx, config.AgentModel, &genai.ClientConfig{
		APIKey: config.GoogleAPIKey,
	})
	if err != nil {
		log.Error("failed to create agent model", "error", err)
		return nil, err
	}
	return &AgentModel{
		Config: config,
		Ctx: ctx,
		Model: model,
	}, nil
}