package media

import "github.com/keithyw/pitch-in/pkg/agent"

type mediaAgentParameters struct {
	agent.AgentCommandParameters
}

func NewMediaAgentParameters(agentName string, description string, userID string, data map[string]any) agent.AgentCommandParametersInterface{
	return agent.NewAgentCommandParameters(agentName, description, userID, data)
}

func (a *mediaAgentParameters) Validate() bool {
	if a.GetValue("title", "") == nil {
		return false
	}
	return true
}


