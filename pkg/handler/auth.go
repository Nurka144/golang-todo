package handler

import (
	"net/http"

	"github.com/Nurka144/todo-app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signup(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponseHandler(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateToUser(input)

	if err != nil {
		newErrorResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type AuthSignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) signin(c *gin.Context) {
	var input AuthSignInRequest

	if err := c.BindJSON(&input); err != nil {
		newErrorResponseHandler(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)

	if err != nil {
		newErrorResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
