package agent

import "github.com/keithyw/pitch-in/pkg/command"

type AgentCommandParametersInterface interface {
	command.CommandParameterInterface
	GetAgentName() string
	GetDescription() string
	GetUserID() string
}

type AgentCommandParameters struct {
	agentName string
	description string
	userID string
	command.BaseParameters
}

func NewAgentCommandParameters(agentName string, description string, userID string, data map[string]any) AgentCommandParametersInterface{
	return &AgentCommandParameters{
		agentName: agentName,
		description: description,
		userID: userID,
		BaseParameters: command.BaseParameters{
			Data: data,
		},
	}
}

func (a *AgentCommandParameters) GetAgentName() string {
	return a.agentName
}

func (a *AgentCommandParameters) GetDescription() string {
	return a.description
}

func (a *AgentCommandParameters) GetUserID() string {
	return a.userID
}