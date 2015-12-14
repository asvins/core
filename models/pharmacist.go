package models

import "github.com/jinzhu/gorm"

/*
*	Pharmacist struct
 */
type Pharmacist struct {
	Base
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	CRF        string      `json:"crf" gorm:"column:crf"`
	Email      string      `json:"email"`
	Specialty  int         `json:"specialty"`
	Treatments []Treatment `json:"treatments"`
}

func (p *Pharmacist) Save(db *gorm.DB) error {
	return db.Create(p).Error
}

func (p *Pharmacist) Update(db *gorm.DB) error {
	return db.Save(p).Error
}

func (p *Pharmacist) Delete(db *gorm.DB) error {
	return db.Delete(p).Error
}

func (p *Pharmacist) Retrieve(db *gorm.DB) ([]Pharmacist, error) {
	var ps []Pharmacist

	err := db.Where(p).Find(&ps, p.Base.BuildQuery()).Error
	return ps, err
}
