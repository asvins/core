package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

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

	es, err := FindFeedEvents(from, patientID)
	if err != nil {
		return errors.NotFound("No events")
	}
	rend.JSON(w, 200, es)
	return nil
}

func handlePatientUpdated(msg []byte) {
	var p Patient
	err := json.Unmarshal(msg, &p)

	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
	}

	createFeedEvent(p)
}

func handlePackCreated(msg []byte) {
	var p om.Pack
	err := json.Unmarshal(msg, &p)

	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
	}

	createFeedEvent(p)
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
	e := NewEvent(i)
	e.Create()
}
