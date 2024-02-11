package database

import (
	"fmt"

	"github.com/agusheryanto182/go-online-store-mvp/config"
	"github.com/agusheryanto182/go-online-store-mvp/entities"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitialDB(cnf config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		cnf.Database.DbHost, cnf.Database.DbUser, cnf.Database.DbPass, cnf.Database.DbName, cnf.Database.DbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to open database", err.Error())
		return nil
	}
	log.Info("Database connected")
	return db
}

func TableMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		entities.User{},
		entities.Category{},
		entities.Product{},
		entities.Cart{},
		entities.CartItem{},
		entities.Order{},
	)
	if err != nil {
		log.Fatal("Migration table is failed", err.Error())
	} else {
		log.Info("Migration table is success")
	}
}
