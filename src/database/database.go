package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB é uma variável global que mantém a conexão com o banco de dados
var DB *gorm.DB

// InitDatabase inicializa a conexão com o banco de dados PostgreSQL
func InitDatabase() {
	dsn := "user=postgres password=israelsenha dbname=atvs3 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrar esquemas de banco de dados
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}

	DB = db
}

// User representa um modelo simples de usuário
type User struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

// GetDB retorna a instância do banco de dados
func GetDB() *gorm.DB {
	return DB
}
