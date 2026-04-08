package media

import (
	"log/slog"

	"github.com/keithyw/pitch-in/pkg/agent"
)

type mediaAgentParameters struct {
	log *slog.Logger
	agent.AgentCommandParameters
}

func NewMediaAgentParameters(agentName string, description string, userID string, data map[string]any, log *slog.Logger) agent.AgentCommandParametersInterface{
	p := agent.NewAgentCommandParameters(agentName, description, userID, data)
	return &mediaAgentParameters{
		AgentCommandParameters: *p.(*agent.AgentCommandParameters),
		log: log,
	}
}

func (a *mediaAgentParameters) Validate() bool {
	if a.GetValue("title", "") == nil || a.GetValue("title", "") == "" {
		a.log.Error("title is required")
		return false
	}
	a.log.Info("title", a.GetValue("title", ""))
	return true
}


