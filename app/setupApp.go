package app

import (
	"log"
	"movie-sto/config"
	"movie-sto/redis"
	"movie-sto/routes"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func RunApp() {
	conn, err := grpc.Dial("0.0.0.0:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	config.InitDatabase()
	redis.NewResdisClient()
	router := gin.Default()
	routes.Service(router, conn)
	routes.MovieRoute(router)
	routes.AuthorRoute(router)
	routes.CategoryRoute(router)
	routes.AuthRoute(router)

	log.Println("Server is running on PORT ", ":8080")
	router.Run(":8080")
}
