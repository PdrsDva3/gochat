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

type UserHandler struct {
	service       service.UserServ
	serviceChange service.UserChangeServ
}

func InitUserHandler(service service.UserServ) UserHandler {
	return UserHandler{
		service: service,
	}
}

// @Summary Create user
// @Tags user
// @Accept  json
// @Produce  json
// @Param data body models.UserCreate true "user create"
// @Success 200 {object} int "Successfully created user"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/create [post]
func (h UserHandler) Create(g *gin.Context) {
	var newUser models.UserCreate

	if err := g.ShouldBindJSON(&newUser); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	id, err := h.service.Create(ctx, newUser)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary Get user
// @Tags user
// @Accept  json
// @Produce  json
// @Param id query int true "UserID"
// @Success 200 {object} int "Successfully get user"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/{id} [get]
func (h UserHandler) Get(c *gin.Context) {
	id := c.Query("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	user, err := h.service.Get(ctx, aid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// @Summary Change password
// @Tags user
// @Accept  json
// @Produce  json
// @Param data body models.UserChangePWD true "change password"
// @Success 200 {object} int "Success changing"
// @Failure 400 {object} map[string]string "Invalid id"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/pwd [put]
func (h UserHandler) ChangePWD(g *gin.Context) {
	var user models.UserChangePWD
	if err := g.ShouldBindJSON(&user); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := h.serviceChange.PWD(ctx, user)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"change": "success"})
}

// @Summary Login user
// @Tags user
// @Accept  json
// @Produce  json
// @Param data body models.UserLogin true "user login"
// @Success 200 {object} int "Successfully login user"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/login [post]
func (h UserHandler) Login(g *gin.Context) {
	var user models.UserLogin

	if err := g.ShouldBindJSON(&user); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := g.Request.Context()

	id, err := h.service.Login(ctx, user)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"userID": id})
}

// @Summary Delete user
// @Tags user
// @Accept  json
// @Produce  json
// @Param id query int true "UserID"
// @Success 200 {object} int "Successfully deleted"
// @Failure 400 {object} map[string]string "Invalid id"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/{id} [delete]
func (h UserHandler) Delete(g *gin.Context) {
	userID := g.Query("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err = h.service.Delete(ctx, id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"delete": id})
}
