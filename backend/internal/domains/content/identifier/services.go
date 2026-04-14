package identifier

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"google.golang.org/adk/artifact"

	agt "github.com/keithyw/pitch-in/internal/agent/identifier"
	"github.com/keithyw/pitch-in/pkg/agent"
	"github.com/keithyw/pitch-in/pkg/templating"
)

type IdentifierService interface {
	IdentifyImage(ctx context.Context, prompt, filanem string, reader io.Reader) (map[string]any, error)
}

type identifierServiceImpl struct {
	manager templating.TemplateManager
	log *slog.Logger
}

func NewIdentifierService(manager templating.TemplateManager, log *slog.Logger) IdentifierService {
	return &identifierServiceImpl{
		manager: manager,
		log: log,
	}
}

func (s *identifierServiceImpl) IdentifyImage(ctx context.Context, prompt, filename string, reader io.Reader) (map[string]any, error) {
	imageData, err := io.ReadAll(reader)
	if err != nil {
		s.log.Error("Failed reading image", "error", err)
		return nil, fmt.Errorf("failed reading image: %w", err)
	}

	buf := 512
	if len(imageData) < buf {
		buf = len(imageData)
	}
	detectedType := http.DetectContentType(imageData[:buf])
	agentName := "image-identification-agent"
	userName := "show-user"	
	description := "Expert researcher that identifies a section of an image and its properties"
	artifactService := artifact.InMemoryService()
	disableTransfer := true

	paramData := map[string]any{
		"prompt": prompt,
		"mime_type": detectedType,
		"filename": filename,
		"image_data": imageData,
	}
	
	cfg := agent.AgentCommandConfig{
		Ctx: ctx,
		Params: agt.NewIdentifierAgentParameters(
			agentName,
			description,
			userName,
			paramData,
			s.log,
		),
		ArtifactService: artifactService,
		DisableTransfer: &disableTransfer,
		TemplateManager: s.manager,
		Log: s.log,
	}
	agent := agt.NewIdentifierAgentCommand(cfg)
	res := agent.Execute()
	if res.IsSuccess == false {
		s.log.Error("Agent failed to run", "errors", res.Errors)
		return nil, fmt.Errorf("agent failed to run: %s", res.Errors)
	}
	return res.Data, nil
}