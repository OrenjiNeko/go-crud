package migrate

import (
	"go-crud/initializers"
	"go-crud/models"
)

func MigrateDatabases() {
	initializers.DB.AutoMigrate(&models.Post{})
	initializers.DB.AutoMigrate(&models.User{})
}


