package routes

import (
	"movie-sto/config"
	"movie-sto/handlers"
	"movie-sto/redis"
	"movie-sto/repository"
	"movie-sto/service"

	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.Engine) {
	handler := handlers.NewAuthHandler(service.NewAuthServices(repository.NewAuthRepository(config.DB, redis.RDB)))
	route := router.Group("/api/auth")
	{
		route.POST("/login", handler.LoginAuthor())
	}
}
