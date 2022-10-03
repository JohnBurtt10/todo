package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"

	"github.com/JohnBurtt10/go/app/models"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	// finish imports
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "mytestdb"
)

var (
	// DBConn for gorm
	DBConn *gorm.DB

	// SessionStore for session storage
	SessionStore *session.Store
	// cfg          *config.DatabaseConfig
)

type Database struct {
	*gorm.DB
}

func Setup() {
	if DBConn != nil {
		return
	}

	if err := connect(); err != nil {
		log.Panicf("error could not connect database (%s)", err.Error())
	}

	log.Println("Connected to DB")

	if err := DBConn.AutoMigrate(&models.Task{},
		&models.User{},
		&models.Task{}); err != nil {
		log.Panicf("Failed to automigrate %s", err.Error())
	}
}

func connect() error {
	var err error
	connectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", host, port, user, dbname, password)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)
	DBConn, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	return nil
}
