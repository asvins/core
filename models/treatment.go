package models

import (
	"strconv"

	"github.com/jinzhu/gorm"
)

const (
	TREATMENT_STATUS_ACTIVE = iota
	TREATMENT_STATUS_INACTIVE
)

type Treatment struct {
	Base
	ID            int            `json:"id"`
	MedicId       int            `json:"medic_id"`
	PatientId     int            `json:"patient_id"`
	PharmacistId  int            `json:"pharmacist_id"`
	Title         string         `json:"title"`
	Status        int            `json:"status"`
	Dose          string         `json:"dose"`
	StartDate     int            `json:"start_date"`
	FinishDate    int            `json:"finish_date"`
	Prescriptions []Prescription `json:"prescriptions"`
	Receipts      []Receipt      `json:"receipts"`
}

func (t *Treatment) Save(db *gorm.DB) error {
	return db.Create(t).Error
}

func (t *Treatment) Update(db *gorm.DB) error {
	return db.Save(t).Error
}

func (t *Treatment) Delete(db *gorm.DB) error {
	return db.Delete(t).Error
}

func (t *Treatment) Retrieve(db *gorm.DB) ([]Treatment, error) {
	var ts []Treatment

	err := db.Where(t).Find(&ts, t.Base.BuildQuery()).Error
	return ts, err
}

/*
*	Gambetas
 */
func (t *Treatment) BuildPackHash() string {
	hash := ""
	for _, prescription := range t.Prescriptions {
		hash += strconv.Itoa(prescription.MedicationId) + ","
	}

	if len(hash) != 0 {
		return hash[:len(hash)-1]
	}
	return hash
}
