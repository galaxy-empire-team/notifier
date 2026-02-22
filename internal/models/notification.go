package models

import (
	"time"

	"github.com/galaxy-empire-team/bridge-api/pkg/consts"
)

type Notification struct {
	ID             uint64
	NotificationID consts.NotificationID
	IsRead         bool
	Data           map[string]any
	CreatedAt      time.Time
}
