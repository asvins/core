package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/asvins/core/models"
	"github.com/asvins/router/errors"
)

func validateRecipe(w http.ResponseWriter, r *http.Request) errors.Http {
	rcpt := models.FetchReceipt(r.URL.Query().Get("treatment_id"), db)
	if r.ParseForm() != nil {
		return errors.BadRequest("Invalid input")
	}
	if rcpt.UpdateStatus(models.RecipeStringToStatus(r.Form.Get("status")), db) != nil {
		return errors.NotFound("Not found")
	}
	rend.JSON(w, 200, "{}")
	return nil
}

func fetchRecipe(w http.ResponseWriter, r *http.Request) errors.Http {
	treatmentId := r.URL.Query().Get("treatment_id")
	rId := r.URL.Query().Get("receipt_id")
	if rId == "" { //FETCH ALL: discovery
		rs := models.ListReceipts(treatmentId, db)
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
		fmt.Println(err)
		return errors.BadRequest("Invalid file")
	}

	defer file.Close()

	id, _ := strconv.Atoi(treatmentId)
	rcpt := &models.Receipt{TreatmentID: id, FilePath: "", Status: models.ReceiptStatusUndecided}
	rcpt.Create(db)
	rcpt.FilePath = "upload/" + treatmentId + "/" + strconv.Itoa(int(rcpt.ID))
	rcpt.Save(db) // ID incremental :'(
	os.MkdirAll("upload/"+treatmentId, 0777)
	out, err := os.Create("upload/" + treatmentId + "/" + strconv.Itoa(int(rcpt.ID)))
	fmt.Println(err)
	_, err = io.Copy(out, file)
	fmt.Println(err)
	if err != nil {
		return errors.InternalServerError("Error uploading image")
	}
	rend.JSON(w, 200, rcpt)
	return nil
}
