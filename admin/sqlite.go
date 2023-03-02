package admin

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var db *gorm.DB

func InitDB() {
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,          // Disable color
		},
	)
	dsn := "host=integrationdb.ccxdgh7epa3w.us-east-1.rds.amazonaws.com user=integrationdb password=*Jv3^LK5P8Cc dbname=integrationdb port=5432"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})

	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&Option{})
	if err != nil {
		fmt.Println(err)
		return
	}
}
