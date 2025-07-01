package handler

import "github.com/gin-gonic/gin"

// напишем пустые обработчики и присвоим их к маршрутам.
// опишем пустые хендлеры для эндпоинтов регистрации и авторизации

// Хендлер во фреймворке Gin - функция,
// которая должна в качестве параметра в себе иметь
// указатель на объект gin.Context

func (h *Handler) signUp(c *gin.Context) {

}

func (h *Handler) signIn(c *gin.Context) {

}
