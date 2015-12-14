package models

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
)

const (
	TREATMENT_STATUS_ACTIVE = iota
	TREATMENT_STATUS_INACTIVE
	TREATMENT_STATUS_FINISHED
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
	StartDate     int64          `json:"start_date"`
	FinishDate    int64          `json:"finish_date"`
	Email         string         `json:"email" sql:"-"`
	Prescriptions []Prescription `json:"prescriptions"`
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

	if err := db.Where(t).Find(&ts, t.Base.BuildQuery()).Error; err != nil {
		return nil, err
	}

	for i, o := range ts {
		prescriptions := []Prescription{}

		if err := db.Model(o).Related(&prescriptions, "Prescriptions").Error; err != nil {
			fmt.Println("[ERROR] ", err.Error())
			return nil, err
		}

		o.Prescriptions = prescriptions
		ts[i] = o
	}

	return ts, nil
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
