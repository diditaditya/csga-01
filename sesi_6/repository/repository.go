package repository

import (
	"fmt"
	"log"
	"os"
	prd "sesi_6/product"
	wh "sesi_6/warehouse"

	"sesi_6/repository/product"
	"sesi_6/repository/warehouse"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db            *gorm.DB
	ProductRepo   prd.ProductRepo
	WarehouseRepo wh.WarehouseRepo
}

func New() (*Repository, error) {
	repo := Repository{}
	repo.init()
	return &repo, nil
}

func (repo *Repository) init() {
	db, err := connectToDB()
	if err != nil {
		os.Exit(1)
	}
	repo.db = db
	repo.WarehouseRepo = warehouse.New(db)
	repo.ProductRepo = product.New(db)
}

func connectToDB() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
		return nil, err
	}

	return db, nil
}
