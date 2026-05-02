package notification

import (
	"context"

	"github.com/google/uuid"

	"github.com/galaxy-empire-team/notifier/internal/models"
)

const (
	notificationsLimit = 30
)

type notificationStorage interface {
	GetNotificationsBeforeID(ctx context.Context, userID uuid.UUID, beforeID uint64, limit uint16) ([]models.Notification, error)
	GetNotificationsAfterID(ctx context.Context, userID uuid.UUID, afterID uint64, limit uint16) ([]models.Notification, error)
	SetReadFlag(ctx context.Context, userID uuid.UUID, notificationIDs []uint64) error
}

type Service struct {
	notificationStorage notificationStorage
}

func New(notificationStorage notificationStorage) *Service {
	return &Service{
		notificationStorage: notificationStorage,
	}
}
