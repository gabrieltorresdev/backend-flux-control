package db

import (
	"github.com/gabrieltorresdev/backend-flux-control/internal/infrastructure/persistence/gorm/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormDB(config *viper.Viper) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(config.GetString("db.connection_string")), &gorm.Config{})
}

func NewGormDBWithAutoMigrate(config *viper.Viper) (*gorm.DB, error) {
	db, err := NewGormDB(config)
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&model.User{}, &model.Category{}, &model.Transaction{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
