package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/mammenj/ginusers/users/config"
	"log"
	"os"
	//Postgres driver
	_ "github.com/lib/pq"

	"github.com/mammenj/ginusers/users/models"
	//Postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Init creates a connection the database
func init() {
	//user := getEnv("PG_USER", "postgres")
	//password := getEnv("PG_PASSWORD", "12345")
	//host := getEnv("PG_HOST", "localhost")
	//port := getEnv("PG_PORT", "5432")
	//database := getEnv("PG_DB", "users")
	config, err := config.GetConfiguration("config.json")
	if err != nil {
		log.Fatal(err)
	}

	user := config.Dbuser
	password := config.Dbpassword
	host := config.Dbhost
	port := config.Dbport
	database := config.Db
	dbengine := config.Dbengine
	dbssl := config.Dbssl


	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		user,
		password,
		host,
		port,
		database,
		dbssl)
	db, err = gorm.Open(dbengine, dbinfo)
	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}
	log.Println("Database connected")

	if !db.HasTable(&models.User{}) {
		err := db.CreateTable(&models.User{})
		if err != nil {
			log.Println("Table already exists")
		}
	}

	db.AutoMigrate(&models.User{})
}

//GetDB returns a db connection
func GetDB() *gorm.DB {
	return db
}

//CloseDB the db connection
func CloseDB() {
	db.Close()
}
