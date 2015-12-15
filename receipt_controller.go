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
	pres_id, err := strconv.Atoi(r.URL.Query().Get("prescription_id"))
	if err != nil {
		return errors.BadRequest(err.Error())
	}

	rcpt := models.FetchReceipt(pres_id, db)
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
	prescriptionId := r.URL.Query().Get("prescription_id")
	rId := r.URL.Query().Get("receipt_id")
	if rId == "" { //FETCH ALL: discovery
		rs := models.ListReceipts(prescriptionId, db)
		rend.JSON(w, 200, rs)
		return nil
	}
	http.ServeFile(w, r, "upload/"+prescriptionId+"/"+rId)
	return nil
}

func uploadRecipe(w http.ResponseWriter, r *http.Request) errors.Http {
	prescriptionId := r.URL.Query().Get("prescription_id")
	file, _, err := r.FormFile("receipt")
	if err != nil {
		fmt.Println(err)
		return errors.BadRequest("Invalid file")
	}

	defer file.Close()

	pid, _ := strconv.Atoi(prescriptionId)
	rcpt := &models.Receipt{FilePath: "", Status: models.ReceiptStatusUndecided, PrescriptionId: pid}
	rcpt.Create(db)
	rcpt.FilePath = "upload/" + prescriptionId + "/" + strconv.Itoa(int(rcpt.ID))
	rcpt.Save(db) // ID incremental :'(
	prescr := models.Prescription{ID: rcpt.PrescriptionId}
	fmt.Println("[DEBUG] prescri query obj: ", prescr)
	prescrs, err := prescr.Retreive(db)
	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
	}

	fmt.Println("[DEBUG] prescrs retrieved ", prescrs)
	if prescrs != nil && len(prescrs) >= 1 {
		fmt.Println("[DEBUG] Saving rcpt.id: ", rcpt.ID, " on prescriptionId: ", prescrs[0].ID)
		prescrs[0].ReceiptId = rcpt.ID
		if err := prescrs[0].Update(db); err != nil {
			fmt.Println("[ERROR] ", err.Error())
		}
	}

	os.MkdirAll("upload/"+prescriptionId, 0777)
	out, err := os.Create("upload/" + prescriptionId + "/" + strconv.Itoa(int(rcpt.ID)))
	fmt.Println(err)
	_, err = io.Copy(out, file)
	fmt.Println(err)
	if err != nil {
		return errors.InternalServerError("Error uploading image")
	}
	rend.JSON(w, 200, rcpt)
	return nil
}
