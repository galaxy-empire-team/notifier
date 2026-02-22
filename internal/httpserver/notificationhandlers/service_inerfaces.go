package notificationhandlers

import (
	"context"

	"github.com/google/uuid"

	"github.com/galaxy-empire-team/notifier/internal/models"
)

type NotificationService interface {
	GetNotifications(ctx context.Context, userID uuid.UUID, offset uint16, limit uint16) ([]models.Notification, error)
}
