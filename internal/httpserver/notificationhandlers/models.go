package notificationhandlers

import (
	"time"

	"github.com/galaxy-empire-team/bridge-api/pkg/consts"
)

type NotificationResponse struct {
	NotificationID consts.NotificationID `json:"notificationID"`
	IsRead         bool                  `json:"isRead"`
	Data           map[string]any        `json:"body"`
	CreatedAt      time.Time             `json:"createdAt"`
}

type SearchParamsRequest struct {
	Offset uint16 `json:"offset"`
	Limit  uint16 `json:"limit"`
}

type ErrorResponse struct {
	Err string `json:"err"`
}
