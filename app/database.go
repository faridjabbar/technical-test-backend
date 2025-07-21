package app

import (
	"log"
	"os"
	"technical-test-backend/model/domain"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase(user, host, password, port, db string) *gorm.DB {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + db + "?parseTime=true"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	err = database.AutoMigrate(
		// Migrasi Database
		&domain.User{},
		&domain.History{},
		&domain.TraditionalFood{},
	)
	if err != nil {
		panic("failed to auto migrate schema")
	}
	database.Migrator().CreateConstraint(&domain.History{}, "users")

	return database
}
