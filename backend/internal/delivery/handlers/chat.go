package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"gochat/internal/models"
	"gochat/internal/service"
	"net/http"
	"strconv"
	"time"
)

type ChatHandler struct {
	service service.ChatServ
}

func InitChatHandler(service service.ChatServ) ChatHandler {
	return ChatHandler{
		service: service,
	}
}

// @Summary Create chat
// @Tags chat
// @Accept  json
// @Produce  json
// @Param data body models.ChatCreate true "chat create"
// @Success 200 {object} int "Successfully created chat"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /chat/create [post]
func (h ChatHandler) Create(g *gin.Context) {
	var newChat models.ChatCreate

	if err := g.ShouldBindJSON(&newChat); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	id, err := h.service.Create(ctx, newChat)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary List chat
// @Tags chat
// @Accept  json
// @Produce  json
// @Param id query int true "UserID"
// @Success 200 {object} int "Successfully get list"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /chat/list/{id} [get]
func (h ChatHandler) List(c *gin.Context) {
	id := c.Query("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	list, err := h.service.List(ctx, aid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"list": list})
}

// @Summary Get chat
// @Tags chat
// @Accept  json
// @Produce  json
// @Param id query int true "ChatID"
// @Success 200 {object} int "Successfully get chat"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /chat/{id} [get]
func (h ChatHandler) Get(c *gin.Context) {
	id := c.Query("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	chat, err := h.service.Get(ctx, aid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"chat": chat})
}

// @Summary Delete chat
// @Tags chat
// @Accept  json
// @Produce  json
// @Param id query int true "ID"
// @Success 200 {object} int "Successfully deleted"
// @Failure 400 {object} map[string]string "Invalid id"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /chat/{id} [delete]
func (h ChatHandler) Delete(c *gin.Context) {
	chatID := c.Query("id")
	id, err := strconv.Atoi(chatID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	err = h.service.Delete(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}
