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
	ID              int    `json:"id"`
	ActiveAgent     string `json:"active_agent"`
	Label           int    `json:"label"`
	Dosage          string `json:"dosage"`
	Bula            string `json:"bula"` // sei lá como é bula em inglês...
	Type            int    `json:"type"`
	IntekeMean      int    `json:"intake_mean"`
	Name            string `json:"name"`
	BrRegister      string `json:"br_register"`
	TerapeuticClass string `json:"terapeutic_class"`
	Manufacturer    string `json:"manufacturer"`
}

// Struct that will be used as intermediate on controlllers
type IntermediateMedication struct {
	ID              int    `json:"id"`
	ActiveAgent     string `json:"active_agent"`
	Label           int    `json:"label"`
	Dosage          string `json:"dosage"`
	Bula            string `json:"bula"` // sei lá como é bula em inglês...
	Type            int    `json:"type"`
	IntekeMean      int    `json:"intake_mean"`
	Name            string `json:"name"`
	BrRegister      string `json:"br_register"`
	TerapeuticClass string `json:"terapeutic_class"`
	Manufacturer    string `json:"manufacturer"`
}

func (m *Medication) Save() error {
	return db.Create(m).Error
}

func (m *Medication) Update() error {
	return db.Update(m).Error
}

func (m *Medication) Delete() error {
	return db.Delete(m).Error
}

func (m *Medication) Retreive() ([]Medication, error) {
	var medications []Medication
	err := db.Where(m).Find(&medications).Error
	return medications, err
}
