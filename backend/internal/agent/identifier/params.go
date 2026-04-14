package identifier

import (
	"log/slog"

	"github.com/keithyw/pitch-in/pkg/agent"
	"golang.org/x/exp/slices"
)

type identifierAgentParameters struct {
	log *slog.Logger
	agent.AgentCommandParameters
}

func NewIdentifierAgentParameters(agentName string, description string, userID string, data map[string]any, log *slog.Logger) agent.AgentCommandParametersInterface {
	p := agent.NewAgentCommandParameters(agentName, description, userID, data)
	return &identifierAgentParameters{
		AgentCommandParameters: *p.(*agent.AgentCommandParameters),
		log: log,
	}
}

func (a *identifierAgentParameters) Validate() bool {
	if a.GetValue("prompt", "") == nil || a.GetValue("prompt", "") == "" {
		a.log.Error("prompt is required")
		return false
	}

	if a.GetValue("filename", "") == nil || a.GetValue("filename", "") == "" {
		a.log.Error("filename is required")
		return false
	}

	imageData, ok := a.GetValue("image_data", nil).([]byte)
	if !ok || len(imageData) == 0 {
		a.log.Error("image_data is required")
		return false
	}

	validMimeTypes := []string{"image/jpeg", "image/png", "image/gif", "image/webp"}
	mimeType := a.GetValue("mime_type", "").(string)
	if mimeType == "" {
		a.log.Error("mime_type is required")
		return false
	}

	if !slices.Contains(validMimeTypes, mimeType) {
		a.log.Error("mime_type is invalid")
		return false
	}

	return true
}