package main

import "github.com/jinzhu/gorm"

const (
	GenderMale = iota
	GenderFemale
	GenderOther
)

type Patient struct {
	gorm.Model
	Name           string `json:"name" gorm:"column:name"`
	CPF            string `json:"cpf" gorm:"column:cpf"`
	Email          string `json:"email" gorm:"column:email"`
	MedicalHistory string `json:"medical_history" gorm:"column:medical_history"`
	Weight         string `json:"weight" gorm:"column:weight"`
	Gender         int    `json:"gender" gorm:"column:gender"`
	Avatar         string `json:"avatar" gorm:"column:avatar"`
}

func (m *Patient) Create() error {
	return db.Create(m).Error
}

func (m *Patient) Update() error {
	return db.Save(m).Error
}

func FindPatientByID(id string, m Patient) error {
	return db.Where("id = ?").First(&m).Error
}

func ListPatients(ms []Patient) error {
	return db.Where("").Find(&ms).Error
}
