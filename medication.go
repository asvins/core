package main

import "strconv"

// Medication possible labels
const (
	MEDICATION_LABEL_NONE = iota
	MEDICATION_LABEL_YELLOW
	MEDICATION_LABEL_RED
	MEDICATION_LABEL_BLACK
)

// Medication type eg: liquid, pill etc..
const (
	MEDICATION_TYPE_LIQUID = iota
	MEDICATION_TYPE_PILL
	MEDICATION_TYPE_OINTMENT
)

// Struct that will be used as intermediate on controlllers
type IntermediateMedication struct {
	ID              int    `json:"id"`
	ActiveAgent     string `json:"active_agent"`
	Label           string `json:"label"`
	Dosage          string `json:"dosage"`
	Bula            string `json:"bula"` // sei lá como é bula em inglês...
	Type            string `json:"type"`
	Name            string `json:"name"`
	BrRegister      string `json:"br_register"`
	TerapeuticClass string `json:"terapeutic_class"`
	Manufacturer    string `json:"manufacturer"`
	PrescriptionId  int    `json:"prescription_id"`
}

func (im *IntermediateMedication) LabelEnum() int {
	switch im.Label {
	case "none":
		return MEDICATION_LABEL_NONE

	case "yellow":
		return MEDICATION_LABEL_YELLOW

	case "red":
		return MEDICATION_LABEL_RED

	case "black":
		return MEDICATION_LABEL_BLACK

	default:
		return -1
	}
}

func (im *IntermediateMedication) TypeEnum() int {
	switch im.Type {
	case "liquid":
		return MEDICATION_TYPE_LIQUID

	case "pill":
		return MEDICATION_TYPE_PILL

	case "ointment":
		return MEDICATION_TYPE_OINTMENT

	default:
		return -1
	}
}

func (im *IntermediateMedication) Medication() *Medication {
	return &Medication{
		ActiveAgent:     im.ActiveAgent,
		Label:           im.LabelEnum(),
		Dosage:          im.Dosage,
		Bula:            im.Bula,
		Type:            im.TypeEnum(),
		Name:            im.Name,
		BrRegister:      im.BrRegister,
		TerapeuticClass: im.TerapeuticClass,
		Manufacturer:    im.Manufacturer,
	}
}

// Medication struct
type Medication struct {
	Base
	ID              int    `json:"id"`
	ActiveAgent     string `json:"active_agent"`
	Label           int    `json:"label"`
	Dosage          string `json:"dosage"`
	Bula            string `json:"bula"` // sei lá como é bula em inglês...
	Type            int    `json:"type"`
	Name            string `json:"name"`
	BrRegister      string `json:"br_register"`
	TerapeuticClass string `json:"terapeutic_class"`
	Manufacturer    string `json:"manufacturer"`
}

func (m *Medication) String() string {
	return "ID: " + strconv.Itoa(m.ID) + " Name: " + m.Name + " ActiveAgent: " + m.ActiveAgent
}

func (m *Medication) Save() error {
	return db.Create(m).Error
}

func (m *Medication) Update() error {
	return db.Save(m).Error
}

func (m *Medication) Delete() error {
	return db.Delete(m).Error
}

func (m *Medication) Retreive() ([]Medication, error) {
	var medications []Medication

	err := db.Where(m).Find(&medications, m.Base.BuildQuery()).Error
	return medications, err
}
