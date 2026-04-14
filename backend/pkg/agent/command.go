package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/keithyw/pitch-in/pkg/command"
	"github.com/keithyw/pitch-in/pkg/templating"
	"google.golang.org/adk/artifact"
	"google.golang.org/adk/tool"
	"google.golang.org/genai"
)

// used for passing into the concrete
// agent's command constructor
type AgentCommandConfig struct {
	Ctx context.Context
	Params AgentCommandParametersInterface	
	Tools []tool.Tool	
	ArtifactService artifact.Service
	DisableTransfer *bool
	TemplateManager templating.TemplateManager
	Log *slog.Logger
}

type AgentCommandInterface interface {
	GenerateSessionKey() string
	GetInputContent() *genai.Content
	GetOutputSchema() *genai.Schema
	Handle(output string) (*command.CommandResults, error)
	PreProcess() error
	PostProcess() error
	Execute() command.CommandResults
	ArtifactSetup() error
}

type BaseAgentCommand struct {
	Ctx context.Context
	Agent *LLMAgent
	AgentModel *AgentModel
	Params AgentCommandParametersInterface
	SessionManager *SessionManager
	Runner *RunnerService
	PromptTemplate string
	PromptData map[string]any
	Tools []tool.Tool
	InternalData map[string]any
	Results *command.CommandResults
	ArtifactService artifact.Service
	PromptTemplateManager templating.TemplateManager
	DisableTransfer *bool
	Log *slog.Logger
	Impl AgentCommandInterface
}

func (c *BaseAgentCommand) GenerateSessionKey() string {
	return ""
}

func (c *BaseAgentCommand) generateAgent() error {
	config := NewAgentConfig()
	model, err := NewAgentModel(c.Ctx, config)
	if err != nil {
		c.Log.Error("failed to create agent model", "error", err)
		return err
	}
	c.AgentModel = model
	c.Log.Info("Agent Model created")
	prompt, err := c.generatePrompt()
	if err != nil {
		c.Log.Error("failed to generate prompt", "error", err)
		return fmt.Errorf("failed to generate prompt: %w", err)
	}

	c.Log.Info("prompt", "prompt", prompt)

	agent, err := NewLLMAgent(
		*model,
		c.Params.GetAgentName(),
		c.Params.GetDescription(),
		prompt,
		c.Impl.GetOutputSchema(),
		c.Tools,
		c.DisableTransfer,
	)
	if err != nil {
		c.Log.Error("failed to create new LLM agent", "error", err)
		return err
	}
	c.Agent = agent
	c.Log.Info("Agent created")
	return nil
}

func (c *BaseAgentCommand) generatePrompt() (string, error) {
	c.Log.Info("prompt template", "template", c.PromptTemplate)
	c.Log.Info("prompt data", "data", c.PromptData)
	return c.PromptTemplateManager.Render(c.PromptTemplate, c.PromptData)
}

func (c *BaseAgentCommand) generateSession() {
	sessionManager := NewSessionManager(
		c.Ctx, 
		c.Params.GetAgentName(), 
		c.Params.GetUserID(),
		c.Impl.GenerateSessionKey(),
	)
	sessionManager.StartSession()
	c.SessionManager = sessionManager
}

func (c *BaseAgentCommand) generateRunner() error {
	if c.Agent.Agent == nil {
		return fmt.Errorf("agent is nil")
	}

	config := RunnerServiceConfig{
		Agent: c.Agent.Agent,
		SessionManager: *c.SessionManager,
		Log: c.Log,
	}
	if c.ArtifactService != nil {		
		c.ArtifactSetup()
		config.ArtifactService = c.ArtifactService
	}

	runner, err := NewRunnerService(
		c.Ctx,
		config,
	)
	if err != nil {
		c.Log.Error("Failed instantiating new runner service", "error", err)
		return err
	}
	
	for _, t := range c.Tools {
		c.Log.Info("Registering tool", "tool", t.Name())
		runner.RegisterTool(t)
	}

	c.Runner = runner
	
	return nil
}

func (c *BaseAgentCommand) ParseOutput(output string) (map[string]any, error) {
	if output == "" {
		c.Log.Error("No response from agent")
		return nil, fmt.Errorf("no response from agent")
	}

	var data map[string]any
	err := json.Unmarshal([]byte(output), &data); if err != nil {
		c.Log.Error("Failed to parse output", "error", err)
		return nil, err
	}

	c.Log.Info("Output parsed", "data", data)

	return data, nil
}

func (c *BaseAgentCommand) GetInputContent() *genai.Content {
	return nil
}

func (c *BaseAgentCommand) GetOutputSchema() *genai.Schema {
	return nil
}

func (c *BaseAgentCommand) Handle(output string) (*command.CommandResults, error) {
	return nil, nil
}

func (c *BaseAgentCommand) PreProcess() error {
	return nil
}

func (c *BaseAgentCommand) PostProcess() error {
	return nil
}

func (c *BaseAgentCommand) Execute() command.CommandResults {
	res := command.CommandResults{}

	if c.Params.Validate() == false {
		c.Log.Error("Invalid parameters error", "error", "Invalid parameters")
		res.IsSuccess = false
		res.Errors = "Invalid parameters"
		return res
	}

	c.Log.Info("parameters validated")

	if err := c.Impl.PreProcess(); err != nil {
		c.Log.Error("Preprocess error", "error", err)
		res.IsSuccess = false
		res.Errors = err.Error()
		return res
	}

	c.Log.Info("Preprocess complete")

	if err := c.generateAgent(); err != nil {
		c.Log.Error("Agent generation error", "error", err)
		res.IsSuccess = false
		res.Errors = err.Error()
		return res
	}

	c.Log.Info("Agent generated")

	c.generateSession()

	c.Log.Info("Session generated")

	if err := c.generateRunner(); err != nil {
		c.Log.Error("Runner generation error", "error", err)
		res.IsSuccess = false
		res.Errors = err.Error()
		return res
	}

	c.Log.Info("Runner generated")

	output, err := c.Runner.Run(c.Impl.GetInputContent())
	if err != nil {
		c.Log.Error("Runner error", "error", err)
		res.IsSuccess = false
		res.Errors = err.Error()
		return res
	}
	commandResults, err := c.Impl.Handle(output)
	if err != nil {
		res.IsSuccess = false
		res.Errors = err.Error()
		return res
	}
	c.Results = commandResults

	if err = c.Impl.PostProcess(); err != nil {
		res.IsSuccess = false
		res.Errors = err.Error()
		return res
	}

	if commandResults == nil {
		c.Log.Error("No results found")
		return command.CommandResults{ IsSuccess: false, Errors: "No results found"}
	}

	return *c.Results
}

func (c *BaseAgentCommand) ArtifactSetup() error {
	return nil
}