package models

import "github.com/jinzhu/gorm"

const (
	PATIENT_GENDER_MALE = iota
	PATIENT_GENDER_FEMALE
	PATIENT_GENDER_OTHER
)

/*
*	Patient struct
 */
type Patient struct {
	Base
	ID             int         `json:"id"`
	Name           string      `json:"name"`
	CPF            string      `json:"cpf" gorm:"column:cpf"`
	MedicalHistory string      `json:"medical_history"`
	Gender         int         `json:"gender"`
	Weight         string      `json:"weight"`
	Email          string      `json:"email"`
	Treatments     []Treatment `json:"treatments"`
}

func (p *Patient) Save(db *gorm.DB) error {
	return db.Create(p).Error
}

func (p *Patient) Update(db *gorm.DB) error {
	return db.Save(p).Error
}

func (p *Patient) Delete(db *gorm.DB) error {
	return db.Delete(p).Error
}

func (p *Patient) Retreive(db *gorm.DB) ([]Patient, error) {
	var ps []Patient

	err := db.Where(p).Find(&ps, p.Base.BuildQuery()).Error
	return ps, err
}
