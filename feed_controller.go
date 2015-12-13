package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/asvins/core/models"
	om "github.com/asvins/operations/models"
	"github.com/asvins/router/errors"
	sm "github.com/asvins/subscription/models"
)

func handleGetFeed(w http.ResponseWriter, req *http.Request) errors.Http {
	patientId := req.URL.Query().Get("patient_id")
	from := time.Now().AddDate(0, 0, -1)
	patientID, err := strconv.Atoi(patientId)
	if err != nil {
		return errors.BadRequest("Invalid patient id")
	}

	es, err := models.FindFeedEvents(from, patientID, db)
	if err != nil {
		return errors.NotFound("No events")
	}
	rend.JSON(w, 200, es)
	return nil
}

func handlePatientUpdated(msg []byte) {
	var p models.Patient
	err := json.Unmarshal(msg, &p)

	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
	}

	createFeedEvent(p)
}

func handleBoxCreated(msg []byte) {
	var b om.Box
	err := json.Unmarshal(msg, &b)

	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
	}

	createFeedEvent(b)
}

func handleSubscriptionUpdated(msg []byte) {
	var s sm.Subscription
	err := json.Unmarshal(msg, &s)

	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
	}

	createFeedEvent(s)
}

func createFeedEvent(i interface{}) {
	e := models.NewEvent(i)
	e.Create(db)
}
