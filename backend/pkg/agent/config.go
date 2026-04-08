package agent

import "os"

type AgentConfig struct {
	GoogleAPIKey string
	AgentModel string
}

func NewAgentConfig() *AgentConfig {
	return &AgentConfig{
		GoogleAPIKey: os.Getenv("GOOGLE_API_KEY"),
		AgentModel: os.Getenv("AGENT_MODEL"),
	}
}