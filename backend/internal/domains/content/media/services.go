package media

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/keithyw/pitch-in/internal/agent/media"
	"github.com/keithyw/pitch-in/pkg/agent"
	"github.com/keithyw/pitch-in/pkg/templating"
)

type MediaService interface {
	QueryShow(ctx context.Context) (map[string]any, error)
}

type mediaServiceImpl struct {
	manager templating.TemplateManager
	log *slog.Logger
}

func NewMediaService(manager templating.TemplateManager, log *slog.Logger) MediaService {
	return &mediaServiceImpl{
		manager: manager,
		log: log,
	}
}

func (s *mediaServiceImpl) QueryShow(ctx context.Context) (map[string]any, error) {
	params := agent.NewAgentCommandParameters("media query agent", "asks about a shows cast", "show-user", nil)	
	cfg := agent.AgentCommandConfig{
		Ctx: ctx,
		Params: params,
		TemplateManager: s.manager,
		Log: s.log,
	}
	agt := media.NewCastResearchAgent(cfg)
	res := agt.Execute()
	if res.IsSuccess == false {
		s.log.Error("Agent failed to run", "errors", res.Errors)
		return nil, fmt.Errorf("agent failed to run: %s", res.Errors)
	}
	return res.Data, nil
}