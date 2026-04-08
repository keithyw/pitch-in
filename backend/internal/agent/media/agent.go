package media

import (
	"github.com/keithyw/pitch-in/pkg/agent"
	"github.com/keithyw/pitch-in/pkg/command"
	"google.golang.org/genai"
)

type CastResearchAgent struct {
	cfg agent.AgentCommandConfig
	agent.BaseAgentCommand
}

func NewCastResearchAgent(cfg agent.AgentCommandConfig) agent.AgentCommandInterface {	
	agt := CastResearchAgent{
		cfg: cfg,
		BaseAgentCommand: agent.BaseAgentCommand{
			Ctx: cfg.Ctx,
			Params: cfg.Params,
			Tools: cfg.Tools,
			ArtifactService: cfg.ArtifactService,
			DisableTransfer: cfg.DisableTransfer,
			PromptData: make(map[string]any),
			PromptTemplate: "test_prompt.tmpl",
			PromptTemplateManager: cfg.TemplateManager,
			Log: cfg.Log,
		},
	}
	agt.BaseAgentCommand.Impl = &agt
	return &agt
}

func (c *CastResearchAgent) GenerateSessionKey() string {
	return "case-research-agent-key"
}

func (c *CastResearchAgent) GetInputContent() *genai.Content {
	return genai.NewContentFromText("Generate cast information for a given movie", genai.RoleUser)
}

func (c *CastResearchAgent) GetOutputSchema() *genai.Schema {
	return &genai.Schema{
		Type: genai.TypeObject,
		Properties: map[string]*genai.Schema{
			"cast": {
				Type: genai.TypeArray,
				Items: &genai.Schema{
					Type: genai.TypeObject,
					Properties: map[string]*genai.Schema{
						"actor_name": { Type: genai.TypeString },
						"character_name": { Type: genai.TypeString },
					},
					Required: []string{"actor_name", "character_name"},
				},
			},
			"summary": { Type: genai.TypeString },
		},
		Required: []string{"cast"},
	}
}

func (c *CastResearchAgent) Handle(output string) (*command.CommandResults, error) {
	c.Log.Info("Handling output", "output", output)
	data, err := c.ParseOutput(output)
	if err != nil {
		c.Log.Error("Failed to parse output", "error", err)
		return &command.CommandResults{
			IsSuccess: false,
			Errors: err.Error(),
		}, nil
	}

	return &command.CommandResults{
		Data: data,
		IsSuccess: true,
	}, nil	
}

// func (c *BaseAgentCommand) PreProcess() error {
func (c *CastResearchAgent) PreProcess() error {
	c.Log.Info("setting prompt data")	
	c.PromptData["medium_type"] = c.Params.GetValue("medium_type", "")
	c.PromptData["title"] = c.Params.GetValue("title", "")
	c.PromptData["year"] = c.Params.GetValue("year", "")	
	c.Log.Info("prompt title", "title", c.PromptData["title"])
	return nil
}

func (c *CastResearchAgent) PostProcess() error {
	return nil
}