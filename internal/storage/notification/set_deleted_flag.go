package mission

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *NotificationStorage) SetDeletedFlag(ctx context.Context, userID uuid.UUID, notificationIDs []uint64) error {
	const getNotificationsQuery = `
		UPDATE session_beta.user_notifications
		SET is_deleted = true
		WHERE user_id = $1 AND id = ANY($2);
	`

	_, err := s.DB.Exec(ctx, getNotificationsQuery, userID, notificationIDs)
	if err != nil {
		return fmt.Errorf("DB.Exec(): %w", err)
	}

	return nil
}
