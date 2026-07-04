package main

import (
	"fmt"
	"time"

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
	ID uint    `json:"id" gorm:"primaryKey"`
	IP string   
	START string
	END string
	OPEN_PORTS []int `json:"open_ports" gorm:"serializer:json"`
	CREATED_AT time.Time
}

func AddRecord(record ScanRecord) {
	DB.Create(&record)
}

func GetRecord() ([]ScanRecord , error) {
	var records []ScanRecord
	result := DB.Find(&records)
	return records, result.Error
}