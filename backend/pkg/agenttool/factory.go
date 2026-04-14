package agenttool

import (
	"context"

	"github.com/keithyw/pitch-in/pkg/agent"
	"github.com/keithyw/pitch-in/pkg/templating"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/agenttool"
)

type AgentToolManager interface {
	CreateAgentTool(name, description, template string, params map[string]any, tools []tool.Tool) (tool.Tool, error)
}

type agentToolManagerImpl struct {
	ctx context.Context
	TemplateManager templating.TemplateManager
}

func NewAgentToolManager(ctx context.Context, templateManager templating.TemplateManager) AgentToolManager {
	return &agentToolManagerImpl{
		ctx: ctx,
		TemplateManager: templateManager,
	}
}

func (m *agentToolManagerImpl) CreateAgentTool(name, description, template string, params map[string]any, tools []tool.Tool) (tool.Tool, error) {
	prompt, err := m.TemplateManager.Render(template, params)
	if err != nil {
		return nil, err
	}
	// fk it; note: cache this thing somewhere and pull it out
	model, err := agent.NewAgentModel(m.ctx, agent.NewAgentConfig())
	if err != nil {
		return nil, err
	}

	agt, err := agent.NewLLMAgent(*model, name, description, prompt, nil, tools, nil)
	if err != nil {
		return nil, err
	}
	
	return agenttool.New(agt.Agent, nil), nil
}