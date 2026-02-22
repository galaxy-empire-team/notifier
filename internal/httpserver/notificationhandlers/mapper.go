package notificationhandlers

import (
	"github.com/galaxy-empire-team/notifier/internal/models"
)

func toTransportNotifications(notifications []models.Notification) []NotificationResponse {
	transportNotifications := make([]NotificationResponse, 0, len(notifications))
	for _, n := range notifications {
		transportNotifications = append(transportNotifications, NotificationResponse{
			NotificationID: n.ID,
			IsRead:         n.IsRead,
			Data:           n.Data,
			CreatedAt:      n.CreatedAt,
		})
	}

	return transportNotifications
}
