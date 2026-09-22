package main

import (
	"Portfolio/database"
	"Portfolio/handler"
	"Portfolio/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func check(c *gin.Context) {
	fmt.Println(c)
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func main() {

	r := gin.Default()
	dbpool, err := database.ConnectToDB()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer dbpool.Close()
	repo := &repository.Repository{DB: dbpool}
	genHandler := &handler.GeneratorHandler{Repo: repo}
	r.GET("/ping", handler.HealthCheck)

	r.GET("/api/about", func(c *gin.Context) {
		genHandler.About(c)
	})

	r.GET("/api/projects", func(ctx *gin.Context) {
		genHandler.Projects(ctx)
	})

	r.GET("/api/education", func(ctx *gin.Context) {
		genHandler.Education(ctx)
	})

	r.GET("/api/experience", func(ctx *gin.Context) {
		genHandler.Experience(ctx)
	})

	fmt.Println("Listening at port 8080")
	r.Run(":8080")
}
