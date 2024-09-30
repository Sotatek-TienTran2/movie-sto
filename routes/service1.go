package routes

import (
	"movie-sto/config"
	"movie-sto/handlers"
	"movie-sto/middlewares"
	"movie-sto/proto"
	"movie-sto/redis"
	"movie-sto/repository"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func Service(router *gin.Engine, conn *grpc.ClientConn) {
	client := proto.NewAddAuthorServiceClient(conn)
	route := router.Group("/api/service/author")
	jwtMiddleware := middlewares.NewJWTMiddleware(repository.NewAuthorRepository(config.DB), redis.RDB)
	//authMiddleware := middlewares.NewAuthMiddlewares(redis.RDB)
	handlers := handlers.NewService1Handler(client)
	{
		route.GET("/:id", handlers.FindMovieByIdAuthor())
		//route.POST("/login", authMiddleware.CheckRegistration(), handlers.LoginGRPC())
		route.GET("/logout", jwtMiddleware.Verify(), handlers.LogoutAuthor())
	}
}
