package notification

import (
	"context"

	"github.com/google/uuid"

	"github.com/galaxy-empire-team/notifier/internal/models"
)

type notificationStorage interface {
	GetNotifications(ctx context.Context, userID uuid.UUID, offset uint16, limit uint16) ([]models.Notification, error)
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
