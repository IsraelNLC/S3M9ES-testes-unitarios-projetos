package handlers

import (
	"net/http"
	db "src/database"

	// Caminho para o pacote db
	"github.com/gin-gonic/gin"
)

// CreateUserHandler lida com a criação de um novo usuário
func CreateUserHandler(c *gin.Context) {
	var user db.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.GetDB().Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUsersHandler lida com a recuperação de todos os usuários
func GetUsersHandler(c *gin.Context) {
	var users []db.User
	if err := db.GetDB().Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// SetupRoutes configura as rotas do Gin usando os handlers definidos
func SetupRoutes(r *gin.Engine) {
	r.POST("/users", CreateUserHandler)
	r.GET("/users", GetUsersHandler)
}
