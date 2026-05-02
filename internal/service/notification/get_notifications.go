package notification

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/galaxy-empire-team/notifier/internal/models"
)

func (s *Service) GetNotifications(ctx context.Context, userID uuid.UUID, beforeID uint64, afterID uint64, limit uint16) ([]models.Notification, error) {
	if limit == 0 || limit > notificationsLimit {
		return nil, ErrInvalidLimit
	}

	if beforeID > 0 && afterID > 0 {
		return nil, ErrInvalidPagination
	}

	if beforeID > 0 {
		notifications, err := s.notificationStorage.GetNotificationsBeforeID(ctx, userID, beforeID, limit)
		if err != nil {
			return nil, fmt.Errorf("notificationStorage.GetNotificationsBeforeID(): %w", err)
		}

		return notifications, nil
	}

	notifications, err := s.notificationStorage.GetNotificationsAfterID(ctx, userID, afterID, limit)
	if err != nil {
		return nil, fmt.Errorf("notificationStorage.GetNotificationsAfterID(): %w", err)
	}

	return notifications, nil
}
