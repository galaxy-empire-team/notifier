package httpserver

import (
	"github.com/galaxy-empire-team/notifier/internal/httpserver/notificationhandlers"
)

func (hs *HttpServer) RegisterRoutes(
	notificationService notificationhandlers.NotificationService,
) {
	// ----- Notification Routes -----
	hs.apiRouter.GET("/notifications", notificationhandlers.GetNotifications(notificationService))
	hs.apiRouter.GET("/notifications/count", notificationhandlers.GetNotificationsCount(notificationService))
	hs.apiRouter.POST("/notifications/read", notificationhandlers.SetReadFlag(notificationService))
	hs.apiRouter.POST("/notifications/delete", notificationhandlers.SetDeletedFlag(notificationService))
}
