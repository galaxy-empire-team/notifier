package mission

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"

	"github.com/galaxy-empire-team/notifier/internal/models"
)

func (s *NotificationStorage) GetNotifications(ctx context.Context, userID uuid.UUID, offset uint16, limit uint16) ([]models.Notification, error) {
	const getNotificationsQuery = `
		SELECT 
			id,
			notification_id, 
			data, 
			is_read,
			created_at
		FROM session_beta.user_notifications
		WHERE user_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3;
	`

	rows, err := s.DB.Query(ctx, getNotificationsQuery, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("DB.Query(): %w", err)
	}
	defer rows.Close()

	var notifications []models.Notification
	var body []byte
	for rows.Next() {
		var notification models.Notification
		if err := rows.Scan(
			&notification.ID,
			&notification.NotificationID,
			&body,
			&notification.IsRead,
			&notification.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("rows.Scan(): %w", err)
		}

		if err := json.Unmarshal(body, &notification.Data); err != nil {
			return nil, fmt.Errorf("json.Unmarshal(): %w", err)
		}

		notifications = append(notifications, notification)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Err(): %w", err)
	}

	return notifications, nil
}
