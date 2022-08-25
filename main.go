package main

import (
	"fmt"
	"log"
	"todo/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to the TODO APP !")

	r := gin.Default()

	r.Use(cors.Default()) //it will allow all origin

	// r.GET("/test", Test)
	routes.TodoRouters(r)
	log.Fatal(r.Run())
}

func Test(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Welcome to gin"})
}
