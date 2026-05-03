package notificationhandlers

import (
	"context"

	"github.com/google/uuid"

	"github.com/galaxy-empire-team/notifier/internal/models"
)

type NotificationService interface {
	GetNotificationsCount(ctx context.Context, userID uuid.UUID, afterID uint64) (uint64, error)
	GetNotifications(ctx context.Context, userID uuid.UUID, beforeID uint64, afterID uint64, limit uint16) ([]models.Notification, error)
	SetReadFlag(ctx context.Context, userID uuid.UUID, notificationIDs []uint64) error
}
