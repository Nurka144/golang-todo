package handler

import (
	"fmt"
	"net/http"

	"github.com/Nurka144/todo-app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signup(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		fmt.Println("ascasc")
		newErrorResponseHandler(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("123")
	id, err := h.services.Authorization.CreateToUser(input)

	if err != nil {
		newErrorResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) signin(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": 123123,
	})
}
