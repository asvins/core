package main

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/asvins/core/models"
	"github.com/asvins/router/errors"
)

func FillTreatmentIdWIthUrlValue(t *models.Treatment, params url.Values) error {
	id, err := strconv.Atoi(params.Get("id"))
	if err != nil {
		return err
	}
	t.ID = id

	return nil
}

func retreiveTreatments(w http.ResponseWriter, r *http.Request) errors.Http {
	t := models.Treatment{}
	if err := BuildStructFromQueryString(&t, r.URL.Query()); err != nil {
		return errors.BadRequest(err.Error())
	}

	t.Base.Query = r.URL.Query()

	treatments, err := t.Retrieve(db)
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	if len(treatments) == 0 {
		return errors.NotFound("record not found")
	}
	rend.JSON(w, http.StatusOK, treatments)

	return nil
}

func deleteTreatment(w http.ResponseWriter, r *http.Request) errors.Http {
	t := models.Treatment{}
	if err := FillTreatmentIdWIthUrlValue(&t, r.URL.Query()); err != nil {
		return errors.BadRequest(err.Error())
	}

	if err := t.Delete(db); err != nil {
		return errors.InternalServerError(err.Error())
	}

	rend.JSON(w, http.StatusOK, t.ID)
	return nil
}

func updateTreatment(w http.ResponseWriter, r *http.Request) errors.Http {
	t := models.Treatment{}

	if err := BuildStructFromReqBody(&t, r.Body); err != nil {
		return errors.BadRequest(err.Error())
	}

	if err := FillTreatmentIdWIthUrlValue(&t, r.URL.Query()); err != nil {
		return errors.BadRequest(err.Error())
	}

	if err := t.Update(db); err != nil {
		return errors.InternalServerError(err.Error())
	}

	rend.JSON(w, http.StatusOK, t)
	return nil
}

func insertTreatment(w http.ResponseWriter, r *http.Request) errors.Http {
	t := models.Treatment{}
	if err := BuildStructFromReqBody(&t, r.Body); err != nil {
		return errors.BadRequest(err.Error())
	}

	if err := t.Save(db); err != nil {
		return errors.InternalServerError(err.Error())
	}

	sendTreatmentCreated(&t)

	rend.JSON(w, http.StatusOK, t)
	return nil
}

func validataTreatment(w http.ResponseWriter, r *http.Request) errors.Http {
	id, err := strconv.Atoi(r.URL.Query().Get("treatment_id"))
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	if err := db.Exec("UPDATE treatments SET status = 0 WHERE id = ?", id).Error; err != nil {
		return errors.InternalServerError(err.Error())
	}

	t := models.Treatment{ID: id}
	sendTreatmentApproved(&t)
	return nil
}
