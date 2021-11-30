package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	config "github.com/jussmor/blog/config"
	entities "github.com/jussmor/blog/internal/entities"
)


type PostgresDB interface {
	DB() *gorm.DB
}

type postgresDB struct {
	db *gorm.DB
}

// ConnectDB connecto to db
func ConnectDB() PostgresDB  {
	var err error 
	var db *gorm.DB

	// Config logger database
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	//Get value env
	dbHost := config.GetEnv("DB_HOST")
	dbUser := config.GetEnv("DB_USER")
	dbPassword := config.GetEnv("DB_PASSWORD")
	dbName := config.GetEnv("DB_NAME")
	dbPort := config.GetEnv("DB_PORT")
	dbSsl := config.GetEnv("DB_SSL")
	dbTimeZone := config.GetEnv("DB_TIME_ZONE")

	// Connect db
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",dbHost,dbUser,dbPassword,dbName,dbPort,dbSsl,dbTimeZone)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Println(fmt.Sprintf("Error to loading Database %s", err))
		return nil
	}

	fmt.Println("Connection Opened to Database")
	db.AutoMigrate(&entities.User{}, &entities.Category{}, &entities.Post{}, &entities.Tag{}, &entities.PostTag{})
	fmt.Println("Database Migrated")

	return &postgresDB{
		db: db,
	}
}

func (c postgresDB) DB() *gorm.DB {
	return c.db
}