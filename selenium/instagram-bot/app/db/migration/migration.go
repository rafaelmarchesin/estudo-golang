package migration

import (
	"app/db"
	"app/models"
)

// AutoMigration cria a tabela no Banco de Dados se ela não existir
func AutoMigration() {
	db := db.Connect()

	db.AutoMigrate(
		&models.User{},
	)
}
