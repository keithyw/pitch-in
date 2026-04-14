package tool

import (
	"github.com/keithyw/pitch-in/pkg/agenttool"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/geminitool"
)

type SearchTool struct {
	factory agenttool.AgentToolManager
}

func NewSearchTool(factory agenttool.AgentToolManager) *SearchTool {
	return &SearchTool{
		factory: factory,
	}
}

func (t *SearchTool) CreateSearchTool() (tool.Tool, error) {
	agtTool, err := t.factory.CreateAgentTool(
		"google_search_agent",
		"Essential tool for finding cast members and media facts. Use this to search Google for movie and show details.",
		"google_search_worker_instructions.tmpl",
		map[string]any{},
		[]tool.Tool{
			geminitool.GoogleSearch{},
		},
	)
	if err != nil {
		return nil, err
	}
	return agtTool, nil
}