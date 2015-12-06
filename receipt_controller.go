package main

import (
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/asvins/router/errors"
)

const (
	ReceiptStatusUndecided = iota
	ReceiptStatusValid
	ReceiptStatusInvalid
)

func validateRecipe(w http.ResponseWriter, r *http.Request) errors.Http {
	rcpt := FetchReceipt(r.URL.Query().Get("treatment_id"))
	if r.ParseForm() != nil {
		return errors.BadRequest("Invalid input")
	}
	if rcpt.UpdateStatus(recipeStringToStatus(r.Form.Get("status"))) != nil {
		return errors.NotFound("Not found")
	}
	rend.JSON(w, 200, "{}")
	return nil
}

func fetchRecipe(w http.ResponseWriter, r *http.Request) errors.Http {
	treatmentId := r.URL.Query().Get("treatment_id")
	rId := r.URL.Query().Get("receipt_id")
	if rId == "" { //FETCH ALL: discovery
		rs := ListReceipts(treatmentId)
		rend.JSON(w, 200, rs)
		return nil
	}
	http.ServeFile(w, r, "upload/"+treatmentId+"/"+rId)
	return nil
}

func uploadRecipe(w http.ResponseWriter, r *http.Request) errors.Http {
	treatmentId := r.URL.Query().Get("treatment_id")
	file, header, err := r.FormFile("receipt")
	if err != nil {
		return errors.BadRequest("Invalid file")
	}

	defer file.Close()

	r := &Receipt{TreatmentID: strconv.Atoi(treatmentId), FilePath: "", Status: ReceiptStatusUndecided}
	r.Create()
	r.FilePath = "upload/" + treatmentId + "/" + r.ID
	r.Save() // ID incremental :'(
	os.MkdirAll("upload/"+treatmentId, 0644)
	out, _ := os.Create("upload/" + treatmentId + "/" + r.ID)
	_, err = io.Copy(out, file)
	if err != nil {
		return errors.InternalServerError("Error uploading image")
	}
	rend.JSON(w, 200, r)
}
