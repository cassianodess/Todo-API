package interfaces

import "gorm.io/gorm"

type DatabaseInterface interface {
	Connect() (*gorm.DB, error)
}