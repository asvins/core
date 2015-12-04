package main

// Medication possible labels
const (
	MEDICATION_LABEL_NONE = iota
	MEDICATION_LABEL_YELLOW
	MEDICATION_LABEL_RED
	MEDICATION_LABEL_BLACK
)

// Medication intakeMeans
const (
	MEDICATION_INTAKE_oral = iota
)

// Medication type eg: liquid, pill etc..
const (
	MEDICATION_TYPE_LIQUID = iota
	MEDICATION_TYPE_PILL
)

// Medication struct
type Medication struct {
	ActiveAgent     string
	Label           int
	Dosage          string
	Bula            string // sei lá como é bula em inglês...
	Type            int
	IntekeMeans     int
	Name            string
	BrRegister      string
	TerapeuticClass string
	Manufector      string
}

func (m *Medication) Save() error {
	return db.Create(m).Error
}

func (m *Medication) Update() error {
	return nil
}

func (m *Medication) Delete() error {
	return nil
}

func (m *Medication) Retreive() error {
	return nil
}
