package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Authorization() gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		strArr := strings.Split(authHeader, " ")
		if len(strArr) != 2 {
			newErrorResponseHandler(c, http.StatusUnauthorized, "Unauthorized")
			return
		}

		userId, err := h.services.PaseTokenMiddleware(strArr[1])

		if err != nil {
			newErrorResponseHandler(c, http.StatusUnauthorized, "Unauthorized")
			return
		}

		c.Set("user_info", map[string]interface{}{
			"user_id": userId,
		})

	}
}
