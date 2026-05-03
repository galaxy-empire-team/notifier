package notification

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *Service) GetNotificationsCount(ctx context.Context, userID uuid.UUID, afterID uint64) (uint64, error) {
	notifications, err := s.notificationStorage.GetNotificationsCount(ctx, userID, afterID)
	if err != nil {
		return 0, fmt.Errorf("notificationStorage.GetNotificationsCount(): %w", err)
	}

	return notifications, nil
}
