package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

const (
	ReceiptStatusUndecided = iota
	ReceiptStatusValid
	ReceiptStatusInvalid
)

type Receipt struct {
	ID             int    `json:"id"`
	PrescriptionId int    `json:"prescription_id"`
	FilePath       string `json:"file_path" gorm:"column:file_path"`
	Status         int    `json:"status" gorm:"column:status"`
}

func (r *Receipt) Create(db *gorm.DB) error {
	return db.Create(r).Error
}

func (r *Receipt) Save(db *gorm.DB) error {
	return db.Save(r).Error
}

func (r *Receipt) UpdateStatus(status int, db *gorm.DB) error {
	r.Status = status
	return r.Save(db)
}

func ListReceipts(prescriptionId string, db *gorm.DB) []Receipt {
	var rs []Receipt
	db.Where("prescription_id = ?", prescriptionId).Find(&rs)
	return rs
}

func FetchReceipt(prescriptionId int, db *gorm.DB) Receipt {
	r := Receipt{}
	if err := db.Where("prescription_id = ?", prescriptionId).First(&r).Error; err != nil {
		fmt.Println("[ERROR] ", err.Error())
		fmt.Println("[ERROR] Will return empty receipt")
	}
	return r
}

func RecipeStringToStatus(status string) int {
	switch status {
	case "valid":
		return ReceiptStatusValid
	case "invalid":
		return ReceiptStatusInvalid
	default:
		return ReceiptStatusInvalid
	}
}
