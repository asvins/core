package main

/*
*	Pharmacist struct
 */
type Pharmacist struct {
	Base
	Name       string      `json:"name"`
	CRF        string      `json:"crf" gorm:"column:crf"`
	Email      string      `json:"email"`
	Treatments []Treatment `json:"treatments"`
}

func (p *Pharmacist) Save() error {
	return db.Create(p).Error
}

func (p *Pharmacist) Update() error {
	return db.Save(p).Error
}

func (p *Pharmacist) Delete() error {
	return db.Delete(p).Error
}

func (p *Pharmacist) Retreive() ([]Pharmacist, error) {
	var ps []Pharmacist

	err := db.Where(p).Find(&ps, p.Base.BuildQuery()).Error
	return ps, err
}
