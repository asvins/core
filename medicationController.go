package main

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/asvins/core/models"
	"github.com/asvins/router/errors"
)

func FillMedicationIdWIthUrlValue(m *models.Medication, params url.Values) error {
	id, err := strconv.Atoi(params.Get("id"))
	if err != nil {
		return err
	}
	m.ID = id

	return nil
}

func retreiveMedication(w http.ResponseWriter, r *http.Request) errors.Http {
	m := models.Medication{}
	if err := BuildStructFromQueryString(&m, r.URL.Query()); err != nil {
		return errors.BadRequest(err.Error())
	}

	m.Base.Query = r.URL.Query()

	medications, err := m.Retrieve(db)
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
	im := models.IntermediateMedication{}
	if err := BuildStructFromReqBody(&im, r.Body); err != nil {
		return errors.BadRequest(err.Error())
	}

	m := im.Medication()

	if err := m.Save(db); err != nil {
		return errors.InternalServerError(err.Error())
	}

	sendProductCreated(m)

	rend.JSON(w, http.StatusOK, m)
	return nil
}

func updateMedication(w http.ResponseWriter, r *http.Request) errors.Http {
	im := models.IntermediateMedication{}

	if err := BuildStructFromReqBody(&im, r.Body); err != nil {
		return errors.BadRequest(err.Error())
	}

	m := im.Medication()

	if err := FillMedicationIdWIthUrlValue(m, r.URL.Query()); err != nil {
		return errors.BadRequest(err.Error())
	}

	if err := m.Update(db); err != nil {
		return errors.InternalServerError(err.Error())
	}

	rend.JSON(w, http.StatusOK, m)
	return nil
}

func deleteMedication(w http.ResponseWriter, r *http.Request) errors.Http {
	m := models.Medication{}
	if err := FillMedicationIdWIthUrlValue(&m, r.URL.Query()); err != nil {
		return errors.BadRequest(err.Error())
	}

	if err := m.Delete(db); err != nil {
		return errors.InternalServerError(err.Error())
	}

	rend.JSON(w, http.StatusOK, m.ID)
	return nil
}
