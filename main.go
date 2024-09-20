package main

import (
	"net/http"

	"github.com/bmendescrypex/CRUD_g/src/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	router := gin.Default()

	// Rota raiz
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Bem-vindo à API",
		})
	})

	// Rota de teste
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080")
}
