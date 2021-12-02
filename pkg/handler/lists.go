package handler

import (
	"fmt"
	"net/http"

	"github.com/Nurka144/todo-app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) {
	fmt.Println(c.GetStringMap("user_info")["user_id"])
	var input todo.TodoList

	if err := c.BindJSON(&input); err != nil {
		newErrorResponseHandler(c, http.StatusBadRequest, err.Error())
	}

	id, err := h.services.TodoList.CreateTodo(input)

	if err != nil {
		newErrorResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"id": id,
	})
}

func (h *Handler) getAllList(c *gin.Context) {
	todoLists, _ := h.services.TodoList.GetAllTodo()

	c.JSON(200, gin.H{
		"data": todoLists,
	})
}

func (h *Handler) getList(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
