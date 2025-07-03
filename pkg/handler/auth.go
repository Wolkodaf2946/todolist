package handler

import (
	"net/http"

	"github.com/Wolkodaf2946/todolist"
	"github.com/gin-gonic/gin"
)

// напишем пустые обработчики и присвоим их к маршрутам.
// опишем пустые хендлеры для эндпоинтов регистрации и авторизации

// Хендлер во фреймворке Gin - функция,
// которая должна в качестве параметра в себе иметь
// указатель на объект gin.Context

func (h *Handler) signUp(c *gin.Context) {
	var input todolist.User // записываем данные из JSON от пользователей

	if err := c.BindJSON(&input); err != nil { // ссылка на объект, у которого хотим распарсить тело JSON
		newErrorResponse(c, http.StatusBadRequest, err.Error()) // пользователь предоставил некорректные данные в запросе
		return
	}
	// теперь мы должны передать данные на слой ниже, в service, где реализована бизнес-логика регистрации

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {

}
