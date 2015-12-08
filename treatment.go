package main

import "strconv"

const (
	TREATMENT_STATUS_ACTIVE = iota
	TREATMENT_STATUS_INACTIVE
)

type Treatment struct {
	Base
	ID            int            `json:"email"`
	MedicId       int            `json:"medic_id"`
	PatientId     int            `json:"patient_id"`
	PharmacistId  int            `json:"pharmacist_id"`
	Title         string         `json:"title"`
	Status        int            `json:"status"`
	Dose          string         `json:"dose"`
	Prescriptions []Prescription `json:"prescriptions"`
	Receipts      []Receipt      `json:"receipts"`
}

func (t *Treatment) Save() error {
	return db.Create(t).Error
}

func (t *Treatment) Update() error {
	return db.Save(t).Error
}

func (t *Treatment) Delete() error {
	return db.Delete(t).Error
}

func (t *Treatment) Retreive() ([]Treatment, error) {
	var ts []Treatment

	err := db.Where(t).Find(&ts, t.Base.BuildQuery()).Error
	return ts, err
}

/*
*	Gambetas
 */
func (t *Treatment) BuildPackHash() string {
	hash := ""
	for _, prescription := range t.Prescriptions {
		hash += strconv.Itoa(prescription.MedicationId) + ","
	}

	return hash[:len(hash)-1]
}
