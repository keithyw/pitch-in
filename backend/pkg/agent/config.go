package agent

import "os"

type AgentConfig struct {
	GoogleAPIKey string
	AgentModel string
	APIKey string
	BaseURL string
}

func NewAgentConfig() *AgentConfig {
	return &AgentConfig{
		GoogleAPIKey: os.Getenv("GOOGLE_API_KEY"),
		AgentModel: os.Getenv("AGENT_MODEL"),
		APIKey: os.Getenv("GOOGLE_API_KEY"),
		BaseURL: os.Getenv("GEMINI_BASE_URL"),
	}
}