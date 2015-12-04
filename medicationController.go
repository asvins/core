package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/asvins/router/errors"
)

func FillMedicationIdWIthUrlValue(m *Medication, params url.Values) error {
	id, err := strconv.Atoi(params.Get("id"))
	if err != nil {
		return err
	}
	m.ID = id

	return nil
}

func retreiveMedication(w http.ResponseWriter, r *http.Request) errors.Http {
	m := Medication{}
	if err := BuildStructFromQueryString(&m, r.URL.Query()); err != nil {
		return errors.BadRequest(err.Error())
	}

	medications, err := m.Retreive()
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	if len(medications) == 0 {
		return errors.NotFound("record not found")
	}
	rend.JSON(w, http.StatusOK, medications)

	return nil
}

func insertMedication(w http.ResponseWriter, r *http.Request) errors.Http {
	im := IntermediateMedication{}
	if err := BuildStructFromReqBody(&im, r.Body); err != nil {
		return errors.BadRequest(err.Error())
	}

	m := im.Medication()

	if err := m.Save(); err != nil {
		return errors.InternalServerError(err.Error())
	}

	rend.JSON(w, http.StatusOK, m)
	return nil
}

func updateMedication(w http.ResponseWriter, r *http.Request) errors.Http {
	im := IntermediateMedication{}

	if err := BuildStructFromReqBody(&im, r.Body); err != nil {
		return errors.BadRequest(err.Error())
	}

	m := im.Medication()

	if err := FillMedicationIdWIthUrlValue(m, r.URL.Query()); err != nil {
		return errors.BadRequest(err.Error())
	}

	fmt.Println(m)

	if err := m.Update(); err != nil {
		return errors.InternalServerError(err.Error())
	}

	rend.JSON(w, http.StatusOK, m)
	return nil
}

func deleteMedication(w http.ResponseWriter, r *http.Request) errors.Http {
	m := Medication{}
	if err := FillMedicationIdWIthUrlValue(&m, r.URL.Query()); err != nil {
		return errors.BadRequest(err.Error())
	}

	if err := m.Delete(); err != nil {
		return errors.InternalServerError(err.Error())
	}

	rend.JSON(w, http.StatusOK, m)
	return nil
}
