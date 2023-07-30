package configurations

import (
	"todo/models"
	"todo/repository"
)

func RunMigrations() error {
	return repository.Context.AutoMigrate(
		&models.Todo{},
	)
}