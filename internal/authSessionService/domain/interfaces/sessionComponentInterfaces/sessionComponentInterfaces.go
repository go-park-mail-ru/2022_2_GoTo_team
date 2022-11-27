package sessionComponentInterfaces

import (
	"2022_2_GoTo_team/internal/authSessionService/domain/models"
	"context"
)

type SessionUsecaseInterface interface {
	SessionExists(ctx context.Context, session *models.Session) (bool, error)
	CreateSessionForUser(ctx context.Context, email string, password string) (*models.Session, error)
	RemoveSession(ctx context.Context, session *models.Session) error
	GetUserInfoBySession(ctx context.Context, session *models.Session) (*models.User, error)
	GetUserEmailBySession(ctx context.Context, session *models.Session) (string, error)
}

type SessionRepositoryInterface interface {
	CreateSessionForUser(ctx context.Context, email string) (*models.Session, error)
	GetEmailBySession(ctx context.Context, session *models.Session) (string, error)
	RemoveSession(ctx context.Context, session *models.Session) error
	SessionExists(ctx context.Context, session *models.Session) (bool, error)
	UpdateEmailBySession(ctx context.Context, session *models.Session, newEmail string)
}
