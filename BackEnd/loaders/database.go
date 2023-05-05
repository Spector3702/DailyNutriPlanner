package loaders

import (
	"BackEnd/configs"
	"BackEnd/entity"

	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Create a new connection to database
func setDBConnection() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configs.GetEnvConfigs("DB_USER"),
		configs.GetEnvConfigs("DB_PASSWORD"),
		configs.GetEnvConfigs("DB_HOST"),
		configs.GetEnvConfigs("DB_PORT"),
		configs.GetEnvConfigs("DB_NAME"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Connect DB failed: ", err)
		panic(err)
	}
	db.AutoMigrate(&entity.RecommendedNutrition{})

	return db.Session(&gorm.Session{PrepareStmt: true})
}

var DB *gorm.DB

func Init() {
	DB = setDBConnection()
}

// Close a connection between app and database
func CloseDBConnection() {
	dbSQL, err := DB.DB()
	if err != nil {
		panic("Failed to close connection to database")
	}

	dbSQL.Close()
}
