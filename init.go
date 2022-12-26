package main

import (
	"fmt"
	"log"
	"music_manager/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDb(c Config) *gorm.DB {
	log.Printf("%+v", c.Db)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=GMT",
		c.Db.Host, c.Db.User, c.Db.Pass, c.Db.Name, c.Db.Port)
	db, err := gorm.Open(postgres.Open(dsn), nil)
	if err != nil {
		logrus.
			WithError(err).
			Fatal("failed to init db")
	}

	db.AutoMigrate(models.GetAllModels()...)

	return db
}

func initLogger() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyTime:  "log_time",
		},
	})
}

func initServer(c Config) *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return e
}
