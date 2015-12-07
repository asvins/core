package main

/*
*	Patient struct
 */
type Patient struct {
	Base
	Name           string `json:"name"`
	CPF            string `json:"cpf" gorm:"column:cpf"`
	Label          int    `json:"label"`
	MedicalHistory string `json:"medical_history"`
	Gender         int    `json:"gender"`
	Weight         string `json:"weight"`
	Email          string `json:"email"`
	Avatar         string `json:"avatar"`
}

func (p *Patient) Save() error {
	return db.Create(p).Error
}

func (p *Patient) Update() error {
	return db.Save(p).Error
}

func (p *Patient) Delete() error {
	return db.Delete(p).Error
}

func (p *Patient) Retreive() ([]Patient, error) {
	var ps []Patient

	err := db.Where(p).Find(&ps, p.Base.BuildQuery()).Error
	return ps, err
}
