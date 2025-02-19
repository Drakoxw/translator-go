package main

import (
	"log"

	"translate/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/translate", func(c *gin.Context) {
		handlers.TranslateHandler(c.Writer, c.Request)
	})

	r.PUT("/translate", func(c *gin.Context) {
		handlers.TranslateHandler(c.Writer, c.Request)
	})

	log.Println("Servidor iniciado en el puerto 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
