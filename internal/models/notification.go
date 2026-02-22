package models

import (
	"time"

	"github.com/galaxy-empire-team/bridge-api/pkg/consts"
)

type Notification struct {
	ID        consts.NotificationID
	IsRead    bool
	Data      map[string]any
	CreatedAt time.Time
}
