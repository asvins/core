package main

import "github.com/jinzhu/gorm"

type Pharmacist struct {
	gorm.Model
	Name   string `json:"name" gorm:"column:name"`
	CRF    string `json:"crf" gorm:"column:crf"`
	Email  string `json:"email" gorm:"column:email"`
	Avatar string `json:"avatar" gorm:"column:avatar"`
}

func (m *Pharmacist) Create() error {
	return db.Create(m).Error
}

func (m *Pharmacist) Update() error {
	return db.Save(m).Error
}

func FindPharmacistByID(id string, m Pharmacist) error {
	return db.Where("id = ?").First(&m).Error
}

func ListPharmacists(ms []Pharmacist) error {
	return db.Where("").Find(&ms).Error
}

func FindPharmacistByEmail(email string, m Patient) error {
	return db.Where("email = ?", email).First(&m).Error
}
