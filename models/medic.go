package models

import "github.com/jinzhu/gorm"

const (
	MEDIC_SPECIALTY_ENDOCRINO = iota
	MEDIC_SPECIALTY_GASTRO
	MEDIC_SPECIALTY_GINECO
	MEDIC_SPECIALTY_GENERAL
	MEDIC_SPECIALTY_DERMATO
)

type Medic struct {
	Base
	ID         int    `json:"id"`
	Name       string `json:"name" gorm:"column:name"`
	CRM        string `json:"crm" gorm:"column:crm"`
	CPF        string `json:"cpf" gorm:"column:cpf"`
	Specialty  int    `json:"specialty" gorm:"column:specialty"`
	Email      string `json:"email" gorm:"column:email"`
	Treatments []Treatment
}

func (m *Medic) Create(db *gorm.DB) error {
	return db.Create(m).Error
}

func (m *Medic) Update(db *gorm.DB) error {
	return db.Save(m).Error
}

func FindMedicByID(id string, m Medic, db *gorm.DB) error {
	return db.Where("id = ?").First(&m).Error
}

func ListMedics(ms []Medic, db *gorm.DB) error {
	return db.Where("").Find(&ms).Error
}

func FindMedicByEmail(email string, m Patient, db *gorm.DB) error {
	return db.Where("email = ?", email).First(&m).Error
}

func (m *Medic) Retrieve(db *gorm.DB) ([]Medic, error) {
	var ms []Medic

	err := db.Where(m).Find(&ms, m.Base.BuildQuery()).Error
	return ms, err
}
