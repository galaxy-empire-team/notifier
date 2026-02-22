package notificationhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/galaxy-empire-team/notifier/internal/httpserver/middleware"
)

func GetUserNotifications(notificationService NotificationService) func(c *gin.Context) {
	return func(c *gin.Context) {
		userID, err := middleware.RetrieveUserID(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, ErrorResponse{
				Err: err.Error(),
			})

			return
		}

		var req SearchParamsRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Err: "invalid request body",
			})
			return
		}

		notifications, err := notificationService.GetNotifications(c.Request.Context(), userID, req.Offset, req.Limit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Err: err.Error(),
			})
		}

		c.JSON(http.StatusOK, toTransportNotifications(notifications))
	}
}
