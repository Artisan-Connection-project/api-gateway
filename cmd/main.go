package main

import (
	"api_gateway/api"
	"api_gateway/api/handlers"
	"api_gateway/genproto/product_service"
	"api_gateway/logger"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	logger.InitLogger()
	log := logger.GetLogger()

	router := gin.Default()

	auth := handlers.NewAuthHandler()

	productConn, err := grpc.NewClient(":50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	proClient := product_service.NewProductServiceClient(productConn)

	product := handlers.NewProductHandler(proClient)
	mainHandler := handlers.NewMainHandler(auth, product, log)

	server := api.NewServer(mainHandler)

	server.InitRoutes(router)

	router.Run(":8080")
}
