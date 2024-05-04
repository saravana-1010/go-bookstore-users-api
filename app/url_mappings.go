package app

import (
	"github.com/saravana-1010/go-bookstore-users-api/controllers/ping"
	"github.com/saravana-1010/go-bookstore-users-api/controllers/users"
)

func mapUrls()  {
	router.GET("/ping", ping.Ping)
	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	// router.GET("/users/search", users.SearchUser)
}