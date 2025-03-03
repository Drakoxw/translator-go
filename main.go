package main

import (
	"log"

	"translate/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", handlers.HelloWord)

	r.GET("/translate", handlers.TranslateHandlerV2)

	r.POST("/translate", handlers.TranslateHandler)

	log.Println("Servidor iniciado en el puerto 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
