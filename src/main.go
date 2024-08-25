package main

import (
	"log"
	database "src/database"
	"src/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar o banco de dados
	database.InitDatabase()

	// Configurar o roteador do Gin
	r := gin.Default()

	// Configurar as rotas usando os handlers
	handlers.SetupRoutes(r)

	// Iniciar o servidor
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
