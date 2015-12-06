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
	file, _, err := r.FormFile("receipt")
	if err != nil {
		return errors.BadRequest("Invalid file")
	}

	defer file.Close()

	id, _ := strconv.Atoi(treatmentId)
	rcpt := &Receipt{TreatmentID: id, FilePath: "", Status: ReceiptStatusUndecided}
	rcpt.Create()
	rcpt.FilePath = "upload/" + treatmentId + "/" + strconv.Itoa(int(rcpt.ID))
	rcpt.Save() // ID incremental :'(
	os.MkdirAll("upload/"+treatmentId, 0644)
	out, _ := os.Create("upload/" + treatmentId + "/" + strconv.Itoa(int(rcpt.ID)))
	_, err = io.Copy(out, file)
	if err != nil {
		return errors.InternalServerError("Error uploading image")
	}
	rend.JSON(w, 200, rcpt)
	return nil
}
