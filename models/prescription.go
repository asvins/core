package models

import "github.com/jinzhu/gorm"

const (
	PRESCRIPTION_FREQ_4H = iota
	PRESCRIPTION_FREQ_6H
	PRESCRIPTION_FREQ_8H
	PRESCRIPTION_FREQ_12H
	PRESCRIPTION_FREQ_24H
)

type Prescription struct {
	Base
	ID           int     `json:"id"`
	TreatmentId  int     `json:"treatment_id"`
	MedicationId int     `json:"medication_id"`
	StartingAt   int64   `json:"starting_at"`
	FinishingAt  int64   `json:"finishing_at"`
	Frequency    int     `json:"frequency"`
	Receipt      Receipt `json:"receipt" sql:"-"`
	ReceiptId    int     `json:"receipt_id"`
}

func (p *Prescription) Save(db *gorm.DB) error {
	return db.Create(p).Error
}

func (p *Prescription) Update(db *gorm.DB) error {
	return db.Save(p).Error
}

func (p *Prescription) Delete(db *gorm.DB) error {
	return db.Delete(p).Error
}

func (p *Prescription) Retreive(db *gorm.DB) ([]Prescription, error) {
	var ps []Prescription

	err := db.Where(p).Find(&ps, p.Base.BuildQuery()).Error
	if err != nil {
		return nil, err
	}

	for i, presc := range ps {
		presc.Receipt = FetchReceipt(presc.ReceiptId, db)
		ps[i] = presc
	}

	return ps, err
}
