package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	Database_path := "host=localhost user=postgres password=root dbname=scanner port=5432 sslmode=disable"
	db,err := gorm.Open(postgres.Open(Database_path), &gorm.Config{})
	if err != nil {
		panic("DB not connected "+err.Error())
	}
	db.AutoMigrate(&ScanRecord{})
	DB = db
	fmt.Println("DB initialised succesfully.")
}

type ScanRecord struct {
	gorm.Model
	IP string
	START string
	END string
	OPEN_PORTS []int `gorm:"serializer:json"`
}

func AddRecord(record ScanRecord) {
	ConnectDB()
	DB.Create(&record)
}

func GetRecord() *gorm.DB {
	ConnectDB()
	var records []ScanRecord
	result := DB.Find(&records)
	return result
}