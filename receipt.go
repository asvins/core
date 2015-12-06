package main

import "github.com/jinzhu/gorm"

type Receipt struct {
	gorm.Model
	TreatmentID int    `gorm:"column:treatment_id"`
	FilePath    string `gorm:"column:file_path"`
	Status      int    `gorm:"column:status"`
}

func (r *Receipt) Create() error {
	return db.Create(r).Error
}

func (r *Receipt) Save() error {
	return db.Save(r).Error
}

func (r *Receipt) UpdateStatus(status int) error {
	r.Status = status
	return r.Save()
}

func ListReceipts(treatmentId string) []Receipt {
	var rs []Receipt
	db.Where("treatment_id = ?", treatmentId).Find(&rs)
	return rs
}

func FetchReceipt(treatmentId string) Receipt {
	var r Receipt
	db.Where("treatment_id = ?", treatmentId).First(&r)
	return r
}

func recipeStringToStatus(status string) int {
	switch status {
	case "valid":
		return ReceiptStatusValid
	case "invalid":
		return ReceiptStatusInvalid
	default:
		return ReceiptStatusInvalid
	}
}
