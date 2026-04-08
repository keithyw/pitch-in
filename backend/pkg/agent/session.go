package agent

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"google.golang.org/adk/session"
)

type SessionManager struct {
	ctx context.Context
	Name string
	User string
	key string
	sessionID string
	SessionService session.Service
}

func NewSessionManager(ctx context.Context, name string, user string, key string) *SessionManager {
	svc := session.InMemoryService()
	return &SessionManager{
		Name: name,
		User: user,
		key: key,
		SessionService: svc,
	}
}

func (s *SessionManager) GenerateSessionID() string {
	u := uuid.New().String()[:8]
	
	safeName := strings.ReplaceAll(s.Name, ":", "_")
	safeKey := strings.ReplaceAll(s.key, ":", "_")
	
	s.sessionID = fmt.Sprintf("%s:%s:%s", u, safeName, safeKey)
	return s.sessionID
}

func (s *SessionManager) GetSessionID() string {
	return s.sessionID
}

func (s *SessionManager) StartSession() (*session.CreateResponse, error) {
	if s.sessionID == "" {
		s.GenerateSessionID()
	}

	res, err := s.SessionService.Create(s.ctx, &session.CreateRequest{
		AppName:   s.Name,
		UserID:    s.User,
		SessionID: s.sessionID,
	})
	
	if err != nil {
		return nil, fmt.Errorf("failed to create ADK session: %w", err)
	}

	return res, nil
}