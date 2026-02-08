package initializers

import "valentoins/models"

func SyncDb() {
	DB.AutoMigrate(&models.Valentines{})
}
