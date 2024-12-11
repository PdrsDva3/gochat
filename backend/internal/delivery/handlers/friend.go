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

type FriendHandler struct {
	service service.FriendServ
}

func InitFriendHandler(service service.FriendServ) FriendHandler {
	return FriendHandler{
		service: service,
	}
}

// @Summary Get friends
// @Tags friend
// @Accept  json
// @Produce  json
// @Param id query int true "UserID"
// @Success 200 {object} int "Successfully get user"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /friend/{id} [get]
func (h FriendHandler) Get(c *gin.Context) {
	id := c.Query("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	friend, err := h.service.Get(ctx, aid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"friend": friend})
}

// @Summary GetFriendsInfo friends
// @Tags friend
// @Accept  json
// @Produce  json
// @Param id query int true "UserID"
// @Success 200 {object} int "Successfully get friend info"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /friend/list/{id} [get]
func (h FriendHandler) GetFriendsInfo(c *gin.Context) {
	id := c.Query("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx := c.Request.Context()

	friends, err := h.service.GetFriendsInfo(ctx, aid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"friends": friends})
}

// @Summary Add friend
// @Tags friend
// @Accept  json
// @Produce  json
// @Param data body models.Friend true "user create"
// @Success 200 {object} int "Successfully created user"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /friend [post]
func (h FriendHandler) AddFriend(c *gin.Context) {
	var newFriend models.Friend

	if err := c.ShouldBindJSON(&newFriend); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	err := h.service.AddFriend(ctx, newFriend.IDFriend, newFriend.IDUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"friend": newFriend})
}

// @Summary Delete user
// @Tags friend
// @Accept  json
// @Produce  json
// @Param id_1 query int true "UserID"
// @Param id_2 query int true "UserID"
// @Success 200 {object} int "Successfully deleted"
// @Failure 400 {object} map[string]string "Invalid id"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /friend/{id_1}/{id_2} [delete]
func (h FriendHandler) Delete(c *gin.Context) {
	friendID := c.Query("id_1")
	userID := c.Query("id_2")
	id1, err := strconv.Atoi(friendID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id2, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	err = h.service.Delete(ctx, id1, id2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"delete": "success"})
}
