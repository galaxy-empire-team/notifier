package notificationhandlers

import (
	"time"

	"github.com/galaxy-empire-team/bridge-api/pkg/consts"
)

type NotificationResponse struct {
	ID             uint64                `json:"id"`
	NotificationID consts.NotificationID `json:"notificationID"`
	IsRead         bool                  `json:"isRead"`
	Data           map[string]any        `json:"body"`
	CreatedAt      time.Time             `json:"createdAt"`
}

type SearchParamsRequest struct {
	Offset uint16 `json:"offset"`
	Limit  uint16 `json:"limit"`
}

type NotificationIDsRequest struct {
	NotificationIDs []uint64 `json:"notificationIDs"`
}

type ErrorResponse struct {
	Err string `json:"err"`
}
