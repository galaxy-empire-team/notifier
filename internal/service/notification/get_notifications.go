package notification

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/galaxy-empire-team/notifier/internal/models"
)

func (s *Service) GetNotifications(ctx context.Context, userID uuid.UUID, offset uint16, limit uint16) ([]models.Notification, error) {
	notifications, err := s.notificationStorage.GetNotifications(ctx, userID, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("notificationStorage.GetNotifications(): %w", err)
	}

	return notifications, nil
}
