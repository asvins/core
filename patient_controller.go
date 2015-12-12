package main

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/asvins/core/models"
	"github.com/asvins/router/errors"
)

func FillPatientIdWIthUrlValue(p *models.Patient, params url.Values) error {
	id, err := strconv.Atoi(params.Get("id"))
	if err != nil {
		return err
	}
	p.ID = id

	return nil
}

func handlePatientRetrieve(w http.ResponseWriter, r *http.Request) errors.Http {
	p := models.Patient{}
	if err := BuildStructFromQueryString(&p, r.URL.Query()); err != nil {
		return errors.BadRequest(err.Error())
	}

	p.Base.Query = r.URL.Query()

	ps, err := p.Retrieve(db)
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	if len(ps) == 0 {
		return errors.NotFound("record not found")
	}
	rend.JSON(w, http.StatusOK, ps)

	return nil
}

func handlePatientUpdate(w http.ResponseWriter, r *http.Request) errors.Http {
	p := models.Patient{}

	if err := BuildStructFromReqBody(&p, r.Body); err != nil {
		return errors.BadRequest(err.Error())
	}

	if err := FillPatientIdWIthUrlValue(&p, r.URL.Query()); err != nil {
		return errors.BadRequest(err.Error())
	}

	if err := p.Update(db); err != nil {
		return errors.InternalServerError(err.Error())
	}

	rend.JSON(w, http.StatusOK, p)
	return nil
}
