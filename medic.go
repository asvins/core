package main

import "github.com/jinzhu/gorm"

// TODO
const (
	MEDIC_SPECIALTY_ = iota
)

type Medic struct {
	gorm.Model
	Name       string `json:"name" gorm:"column:name"`
	CRM        string `json:"crm" gorm:"column:crm"`
	CPF        string `json:"cpf" gorm:"column:cpf"`
	Specialty  string `json:"specialty" gorm:"column:specialty"`
	Email      string `json:"email" gorm:"column:email"`
	Treatments []Treatment
}

func (m *Medic) Create() error {
	return db.Create(m).Error
}

func (m *Medic) Update() error {
	return db.Save(m).Error
}

func FindMedicByID(id string, m Medic) error {
	return db.Where("id = ?").First(&m).Error
}

func ListMedics(ms []Medic) error {
	return db.Where("").Find(&ms).Error
}

func FindMedicByEmail(email string, m Patient) error {
	return db.Where("email = ?", email).First(&m).Error
}
