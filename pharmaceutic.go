package main

/*
*	Pharmaceutic struct
 */
type Pharmaceutic struct {
	Base
	Name      string `json:"name"`
	CPF       string `json:"cpf" gorm:"column:cpf"`
	Specialty string `json:"specialty"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

func (p *Pharmaceutic) Save() error {
	return db.Create(p).Error
}

func (p *Pharmaceutic) Update() error {
	return db.Save(p).Error
}

func (p *Pharmaceutic) Delete() error {
	return db.Delete(p).Error
}

func (p *Pharmaceutic) Retreive() ([]Pharmaceutic, error) {
	var ps []Pharmaceutic

	err := db.Where(p).Find(&ps, p.Base.BuildQuery()).Error
	return ps, err
}
