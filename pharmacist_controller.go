package main

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/asvins/core/models"
	"github.com/asvins/router/errors"
)

func FillPharmacistIdWIthUrlValue(p *models.Pharmacist, params url.Values) error {
	id, err := strconv.Atoi(params.Get("id"))
	if err != nil {
		return err
	}
	p.ID = id

	return nil
}

func handlePharmacistRetrieve(w http.ResponseWriter, r *http.Request) errors.Http {
	p := models.Pharmacist{}
	if err := BuildStructFromQueryString(&p, r.URL.Query()); err != nil {
		return errors.BadRequest(err.Error())
	}

	p.Base.Query = r.URL.Query()

	pharms, err := p.Retrieve(db)
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	if len(pharms) == 0 {
		return errors.NotFound("record not found")
	}
	rend.JSON(w, http.StatusOK, pharms)

	return nil
}

func handlePharmacistUpdate(w http.ResponseWriter, r *http.Request) errors.Http {
	p := models.Pharmacist{}

	if err := BuildStructFromReqBody(&p, r.Body); err != nil {
		return errors.BadRequest(err.Error())
	}

	if err := FillPharmacistIdWIthUrlValue(&p, r.URL.Query()); err != nil {
		return errors.BadRequest(err.Error())
	}

	if err := p.Update(db); err != nil {
		return errors.InternalServerError(err.Error())
	}

	rend.JSON(w, http.StatusOK, p)
	return nil
}
