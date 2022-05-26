package main

import (
	"github.com/JerbesonViny/nostarve/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type User struct {
	Username		string `json:"username"`
	Password		string `json:"password"`
}

func main() {
	godotenv.Load(); // Load all env variables
	app := gin.Default(); // Create a app


	app.GET("/", func(ctx *gin.Context) {


		ctx.JSON(200, gin.H{
			"err": "nice",
		});
	}); // Creating a route for return one JSON 

	app.POST("/register/", controllers.CreateConsumer);

	app.Run(); // Starting server on port 8080
};