package mission

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *NotificationStorage) GetNotificationsCount(ctx context.Context, userID uuid.UUID, afterID uint64) (uint64, error) {
	const getNotificationsCountQuery = `
		SELECT 
			COUNT(*)
		FROM session_beta.user_notifications
		WHERE user_id = $1 AND id > $2;
	`

	var count uint64
	if err := s.DB.QueryRow(ctx, getNotificationsCountQuery, userID, afterID).Scan(&count); err != nil {
		return 0, fmt.Errorf("DB.QueryRow(): %w", err)
	}

	return count, nil
}
