package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/saravana-1010/go-bookstore-users-api/domain/users"
	"github.com/saravana-1010/go-bookstore-users-api/services"
	"github.com/saravana-1010/go-bookstore-users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	// var user users.User
	user := users.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Code, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Code, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Params.ByName("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id")
		c.JSON(err.Code, err)
		return
	}
	result, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Code, getErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func SearchUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"status": "Not impl"})
}