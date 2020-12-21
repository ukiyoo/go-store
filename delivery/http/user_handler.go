package http

import (
	"github.com/gin-gonic/gin"
	"go-store/domain"
	"go-store/usecase"
	"net/http"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecase
}

func NewUserHandler(r *gin.Engine, us usecase.UserUsecase) *gin.Engine {
	handler := &UserHandler{
		UserUsecase: us,
	}

	r.GET("/ping", Ping)
	r.POST("/login", handler.Login)

	return r
}

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func (u *UserHandler) Login(c *gin.Context) {
	var user domain.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = u.UserUsecase.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
