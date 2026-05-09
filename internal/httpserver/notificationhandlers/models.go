package notificationhandlers

import (
	"time"

	"github.com/galaxy-empire-team/bridge-api/pkg/consts"
)

type NotificationResponse struct {
	ID             uint64                `json:"id"`
	Version        uint8                 `json:"version"`
	NotificationID consts.NotificationID `json:"notificationID"`
	IsRead         bool                  `json:"isRead"`
	Data           map[string]any        `json:"body"`
	CreatedAt      time.Time             `json:"createdAt"`
}

type SearchParamsRequest struct {
	Limit    uint16 `json:"limit"`
	BeforeID uint64 `json:"beforeID"`
	AfterID  uint64 `json:"afterID"`
}

type NotificationIDsRequest struct {
	NotificationIDs []uint64 `json:"notificationIDs"`
}

type CountResponse struct {
	Count uint64 `json:"count"`
}

type ErrorResponse struct {
	Err string `json:"err"`
}
