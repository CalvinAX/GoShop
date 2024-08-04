package initializers

import "github.com/CalvinAX/GoShop/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Product{})
}
