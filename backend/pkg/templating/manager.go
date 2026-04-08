package templating

import (
	"fmt"

	"github.com/flosch/pongo2/v7"
)

type TemplateManager struct {
	templateSet *pongo2.TemplateSet
}

func NewTemplateManager(templateDir string) *TemplateManager {
	loader := pongo2.MustNewLocalFileSystemLoader(templateDir)
	ts := pongo2.NewSet("agent-prompts", loader)

	return &TemplateManager{
		templateSet: ts,
	}
}

func (m *TemplateManager) Render(template string, data map[string]any) (string, error) {
	tpl, err := m.templateSet.FromFile(template)
	if err != nil {
		return "", fmt.Errorf("failed to load template: %s: %w", template, err)
	}
	
	return tpl.Execute(pongo2.Context(data))
}