package identifier

import (
	"github.com/keithyw/pitch-in/pkg/agent"
	"github.com/keithyw/pitch-in/pkg/command"
	"google.golang.org/adk/artifact"
	"google.golang.org/genai"
)

type IdentifierAgentCommand struct {
	cfg agent.AgentCommandConfig
	agent.BaseAgentCommand
}

func NewIdentifierAgentCommand(cfg agent.AgentCommandConfig) agent.AgentCommandInterface {
	agt := IdentifierAgentCommand{
		cfg: cfg,
		BaseAgentCommand: agent.BaseAgentCommand{
			Ctx: cfg.Ctx,
			Params: cfg.Params,
			Tools: cfg.Tools,
			ArtifactService: cfg.ArtifactService,
			DisableTransfer: cfg.DisableTransfer,
			InternalData: make(map[string]any),
			PromptData: make(map[string]any),
			PromptTemplate: "image_identification_prompt.tmpl",
			PromptTemplateManager: cfg.TemplateManager,
			Log: cfg.Log,
		},
	}
	agt.BaseAgentCommand.Impl = &agt
	return &agt
}

func (c *IdentifierAgentCommand) GenerateSessionKey() string {
	return "identifier-agent-key"
}

func (c *IdentifierAgentCommand) GetInputContent() *genai.Content {
	prompt := c.Params.GetValue("prompt", "").(string)
	return genai.NewContentFromText(prompt, genai.RoleUser)
}

func (c *IdentifierAgentCommand) GetOutputSchema() *genai.Schema {
	return &genai.Schema{
		Type: genai.TypeObject,
		Properties: map[string]*genai.Schema{
			"name": { Type: genai.TypeString },
			"description": { Type: genai.TypeString },
			"tags": {
				Type: genai.TypeArray,
				Items: &genai.Schema{
					Type: genai.TypeString,
				},
			},
		},
		Required: []string{"name", "description", "tags"},
	}
}

func (c *IdentifierAgentCommand) Handle(output string) (*command.CommandResults, error) {
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

func (c *IdentifierAgentCommand) PreProcess() error {
	imageData := c.Params.GetValue("image_data", nil).([]byte)
	c.InternalData["image_data"] = imageData
	c.PromptData["prompt"] = c.Params.GetValue("prompt", "")
	c.PromptData["mime_type"] = c.Params.GetValue("mime_type", "")
	return nil
}

func (c *IdentifierAgentCommand) PostProcess() error {
	return nil
}

func (c *IdentifierAgentCommand) ArtifactSetup() error {
	imageData := c.Params.GetValue("image_data", nil).([]byte)
	mimeType := c.Params.GetValue("mime_type", "").(string)

	_, err := c.ArtifactService.Save(
		c.Ctx,
		&artifact.SaveRequest{
			AppName: c.Params.GetAgentName(),
			UserID: c.Params.GetUserID(),
			SessionID: c.SessionManager.GetSessionID(),
			FileName: c.Params.GetValue("filename", "").(string),
			Part: genai.NewPartFromBytes(imageData, mimeType),
		},
	)
	if err != nil {
		c.Log.Error("Failed to save artifact", "error", err)
		return err
	}
	return nil
}