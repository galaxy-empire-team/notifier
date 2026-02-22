package httpserver

import (
	"github.com/galaxy-empire-team/notifier/internal/httpserver/notificationhandlers"
)

func (hs *HttpServer) RegisterRoutes(
	notificationService notificationhandlers.NotificationService,
) {
	// ----- Notification Routes -----
	hs.apiRouter.GET("/notification/get", notificationhandlers.GetUserNotifications(notificationService))
	hs.apiRouter.POST("/notification/read", notificationhandlers.SetReadFlag(notificationService))
}
