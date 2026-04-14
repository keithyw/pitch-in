package media

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/keithyw/pitch-in/internal/agent/media"
	"github.com/keithyw/pitch-in/pkg/agent"
	"github.com/keithyw/pitch-in/pkg/templating"
	"google.golang.org/adk/tool"
)

type MediaService interface {
	QueryShow(ctx context.Context, data map[string]any) (map[string]any, error)
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

func (s *mediaServiceImpl) QueryShow(ctx context.Context, data map[string]any) (map[string]any, error) {
	// factory := agenttool.NewAgentToolManager(ctx, s.manager)
	// searchTool := mytool.NewSearchTool(factory)
	// googleSearch, err := searchTool.CreateSearchTool()
	// if err != nil {
	// 	s.log.Error("Failed to create search tool", "error", err)
	// 	return nil, err
	// }

	// lazy
	description := "Expert researcher that uses Google Search to find and validate cast lists for movies and TV shows."
	cfg := agent.AgentCommandConfig{
		Ctx: ctx,
		Params: media.NewMediaAgentParameters("media-query-agent", description, "show-user", data, s.log),
		TemplateManager: s.manager,
		Log: s.log,
		Tools: []tool.Tool{
			// geminitool.GoogleSearch{},
			// googleSearch,
		},
	}
	agt := media.NewCastResearchAgent(cfg)
	res := agt.Execute()
	if res.IsSuccess == false {
		s.log.Error("Agent failed to run", "errors", res.Errors)
		return nil, fmt.Errorf("agent failed to run: %s", res.Errors)
	}
	return res.Data, nil
}