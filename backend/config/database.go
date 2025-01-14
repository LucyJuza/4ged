package config

import (
	"fmt"
	"os"

	"backend/entities"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB         *gorm.DB
	dbhost     = ""
	dbport     = ""
	dbuser     = ""
	dbpassword = ""
	dbname     = ""
)

// Get variables from env vars
func init() {
	dbhost = os.Getenv("DBHOST")
	dbport = os.Getenv("DBPORT")
	dbuser = os.Getenv("DBUSER")
	dbpassword = os.Getenv("DBPASSWORD")
	dbname = os.Getenv("DBNAME")
}

// Connect tries to connect to a database with the given environment variables
func Connect() error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
		dbuser,
		dbpassword,
		dbhost,
		dbport,
	)

	log.Info("Attempting to connect with ", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		TranslateError:         true,
		CreateBatchSize:        100,
	})
	if err != nil {
		log.Error("Couldn't open connection to db")
		panic(err)
	}

	_ = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbname + ";")

	dsn = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbuser,
		dbpassword,
		dbhost,
		dbport,
		dbname,
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		TranslateError:         true,
		CreateBatchSize:        100,
	})
	if err != nil {
		log.Error("Couldn't open connection to db")
		panic(err)
	}

	if err := db.AutoMigrate(&entities.Game{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&entities.Genre{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&entities.Play{}); err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&entities.Player{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&entities.User{}); err != nil {
		panic(err)
	}

	DB = db

	return nil
}
