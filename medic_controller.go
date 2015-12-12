package main

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/asvins/core/models"
	"github.com/asvins/router/errors"
)

func FillMedicIdWIthUrlValue(m *models.Medic, params url.Values) error {
	id, err := strconv.Atoi(params.Get("id"))
	if err != nil {
		return err
	}
	m.ID = id

	return nil
}

func handleMedicRetrieve(w http.ResponseWriter, r *http.Request) errors.Http {
	m := models.Medic{}
	if err := BuildStructFromQueryString(&m, r.URL.Query()); err != nil {
		return errors.BadRequest(err.Error())
	}

	m.Base.Query = r.URL.Query()

	medics, err := m.Retrieve(db)
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	if len(medics) == 0 {
		return errors.NotFound("record not found")
	}
	rend.JSON(w, http.StatusOK, medics)

	return nil
}

func handleMedicUpdate(w http.ResponseWriter, r *http.Request) errors.Http {
	m := models.Medic{}

	if err := BuildStructFromReqBody(&m, r.Body); err != nil {
		return errors.BadRequest(err.Error())
	}

	if err := FillMedicIdWIthUrlValue(&m, r.URL.Query()); err != nil {
		return errors.BadRequest(err.Error())
	}

	if err := m.Update(db); err != nil {
		return errors.InternalServerError(err.Error())
	}

	rend.JSON(w, http.StatusOK, m)
	return nil
}
