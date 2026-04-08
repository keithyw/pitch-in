package agent

import (
	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/tool"
	"google.golang.org/genai"
)

type LLMAgent struct {
	model AgentModel
	name string
	Agent agent.Agent
	Prompt string
}

func NewLLMAgent(model AgentModel, name string, description string, prompt string, schema *genai.Schema, tools []tool.Tool, disableTransfer *bool) (*LLMAgent, error) {
	cfg := llmagent.Config{
		Name: name,		
		Description: description,
		Model: model.Model,
		Instruction: prompt,
		OutputSchema: schema,
		Tools: tools,
	}
	
	if disableTransfer != nil && *disableTransfer {
		cfg.DisallowTransferToParent = true
		cfg.DisallowTransferToPeers = true
	}

	myagent, err := llmagent.New(cfg)
	if err != nil {
		return nil, err
	}
	return &LLMAgent{
		model: model,
		name: name,
		Agent: myagent,		
	}, nil
}