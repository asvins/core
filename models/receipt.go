package models

import "github.com/jinzhu/gorm"

const (
	ReceiptStatusUndecided = iota
	ReceiptStatusValid
	ReceiptStatusInvalid
)

type Receipt struct {
	gorm.Model
	TreatmentID int    `json:"treatment_id" gorm:"column:treatment_id"`
	FilePath    string `json:"file_path" gorm:"column:file_path"`
	Status      int    `json:"status" gorm:"column:status"`
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

func ListReceipts(treatmentId string, db *gorm.DB) []Receipt {
	var rs []Receipt
	db.Where("treatment_id = ?", treatmentId).Find(&rs)
	return rs
}

func FetchReceipt(treatmentId string, db *gorm.DB) Receipt {
	var r Receipt
	db.Where("treatment_id = ?", treatmentId).First(&r)
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
