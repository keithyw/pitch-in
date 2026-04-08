package agent

import (
	"context"
	"fmt"
	"log/slog"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/artifact"
	"google.golang.org/adk/runner"
	"google.golang.org/adk/tool"
	"google.golang.org/genai"
)

type Runnable interface {
	Run(ctx context.Context, args any) (map[string]any, error)
}

// avoid parameter bloat which stinks because it's just a wrapper
// over a wrapper/config
type RunnerServiceConfig struct {
	Agent agent.Agent
	SessionManager SessionManager
	ArtifactService artifact.Service
	Log *slog.Logger
}

type RunnerService struct {
	ctx            context.Context
	Agent agent.Agent
	AgentSessionManager SessionManager
	Runner *runner.Runner
	tools map[string]tool.Tool
	log *slog.Logger
}

func NewRunnerService(ctx context.Context, conf RunnerServiceConfig) (*RunnerService, error) {

	// literally going through hoops through this crap
	config := runner.Config{
		AppName: conf.Agent.Name(),
		Agent: conf.Agent,
		SessionService: conf.SessionManager.SessionService,
	}

	// gollum/smeagol choking deagol scene. y u no make ArtifactService public or have
	// setter??????
	if conf.ArtifactService != nil {
		config.ArtifactService = conf.ArtifactService
	}

	runner, err := runner.New(config)

	if err != nil {
		conf.Log.Error("failed to create runner service", "error", err)
		return nil, err
	}

	return &RunnerService{
		ctx: ctx,
		Agent: conf.Agent,
		AgentSessionManager: conf.SessionManager,
		Runner: runner,
		tools: map[string]tool.Tool{},
		log: conf.Log,
	}, nil
}

func (r *RunnerService) RegisterTool(t tool.Tool) {
	r.tools[t.Name()] = t
}

func (r *RunnerService) handleToolCall(call *genai.FunctionCall) (map[string]any, error) {
	t, ok := r.tools[call.Name]
	if !ok {
		r.log.Error("Tool not registered", "tool", call.Name)
		return nil, fmt.Errorf("tool '%s' requested but not registered", call.Name)
	}

	runnable, ok := t.(Runnable)
	if !ok {
		r.log.Error("Tool is not executable", "tool", call.Name)
		return nil, fmt.Errorf("tool %s is not executable", call.Name)
	}

	return runnable.Run(r.ctx, call.Args)
}

func (r *RunnerService) Run(message *genai.Content) (string, error) {
	msg := message
	sessionID := r.AgentSessionManager.GetSessionID()
	if sessionID == "" {
		r.log.Error("Session ID not set")
		return "", fmt.Errorf("Sesssion ID not set and/or AgentSessionManager has not started via start_session()")
	}

	cfg := agent.RunConfig{}
	for {
		var toolResponses []*genai.Part

		r.log.Info("Starting runner")

		for event, err:= range r.Runner.Run(r.ctx, r.AgentSessionManager.User, sessionID, msg, cfg) {
			if err != nil {
				r.log.Error("Runner error", "error", err)
				return "", err
			}
			
			if event.Content == nil || len(event.Content.Parts) == 0 {
				r.log.Info("No content received but still running")
				continue
			}

			if event.IsFinalResponse() {
				r.log.Info("Final response received", "response", event.Content.Parts[0].Text)
				return event.Content.Parts[0].Text, nil
			}

			r.log.Info("Examining additional parts from event")
			for _, part := range event.Content.Parts {
				if part.FunctionCall != nil {
					res, err := r.handleToolCall(part.FunctionCall)
					if err != nil {
						r.log.Error("Error handling tool call", "error", err)
						return "", err
					}

					if res == nil {
						continue
					}

					toolResponses = append(toolResponses, &genai.Part{
						FunctionResponse: &genai.FunctionResponse{
							Name: part.FunctionCall.Name,
							Response: res,
							ID: part.FunctionCall.ID,
						},					
					})
				}
			}
		}

		if len(toolResponses) > 0 {
			msg = &genai.Content{
				Role: "tool",
				Parts: toolResponses,
			}
			continue
		}

		return "", fmt.Errorf("streamed closed without a final response")
	}
}