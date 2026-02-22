package mission

import (
	"github.com/galaxy-empire-team/notifier/internal/db"
)

// Embed txManager requires different naming -> can't use 'storage' storage name :()
type NotificationStorage struct {
	DB db.DB
}

func New(db db.DB) *NotificationStorage {
	return &NotificationStorage{
		DB: db,
	}
}
