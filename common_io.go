package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	authModels "github.com/asvins/auth/models"
	"github.com/asvins/common_io"
	operationsModel "github.com/asvins/operations/models"
	"github.com/asvins/utils/config"
	"github.com/asvins/warehouse/models"
)

func setupCommonIo() {
	cfg := common_io.Config{}

	err := config.Load("common_io_config.gcfg", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	/*
	*	Producer
	 */
	producer, err = common_io.NewProducer(cfg)
	if err != nil {
		log.Fatal(err)
	}

	/*
	*	Consumer
	 */
	consumer = common_io.NewConsumer(cfg)

	/*
	*	Topics
	 */
	consumer.HandleTopic("user_created", handleUserCreated)
	consumer.HandleTopic("patient_updated", handlePatientUpdated)
	consumer.HandleTopic("pack_created", handlePackCreated)
	consumer.HandleTopic("subscription_updated", handleSubscriptionUpdated)

	if err = consumer.StartListening(); err != nil {
		log.Fatal(err)
	}
}

func userUpdated(msg []byte) {
	var usr authModels.User
	json.Unmarshal(msg, &usr)
	// update user: COMO PEGAR OS EMAILS ANTIGOS??
	switch usr.Scope {
	case "patient":
	case "medic":
	case "pharmacist":
	default:
		break
	}
}

/*
*	Handlers
 */
func handleUserCreated(msg []byte) {
	var usr authModels.User
	err := json.Unmarshal(msg, &usr)

	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
	}

	switch usr.Scope {
	case "patient":
		p := Patient{}
		p.ID = usr.ID
		p.Name = usr.FirstName + " " + usr.LastName
		p.Email = usr.Email

		err = p.Save()

		return

	case "medic":
		m := Medic{}
		m.ID = usr.ID
		m.Name = usr.FirstName + " " + usr.LastName
		m.Email = usr.Email

		err = m.Create()

		return

	case "pharmacist":
		p := Pharmacist{}
		p.ID = usr.ID
		p.Name = usr.FirstName + " " + usr.LastName
		p.Email = usr.Email

		err = p.Save()
		return
	}
	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
	}

}

/*
*	Senders
 */
func sendProductCreated(m *Medication) {
	topic, _ := common_io.BuildTopicFromCommonEvent(common_io.EVENT_CREATED, "product")
	p := models.Product{}

	/*
	*	Bind
	 */
	p.ID = m.ID
	p.Name = m.Name

	/*
	*	json Marshal
	 */
	b, err := json.Marshal(&p)
	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
		return
	}

	producer.Publish(topic, b)
}

func sendTreatmentCreated(t *Treatment) {
	topic, _ := common_io.BuildTopicFromCommonEvent(common_io.EVENT_CREATED, "treatment")
	p := operationsModel.Pack{}

	/*
	* Bind
	 */
	p.Owner = strconv.Itoa(t.PatientId)
	p.Supervisor = strconv.Itoa(t.PharmacistId)
	p.PackType = "medication"

	rand.Seed(time.Now().UTC().UnixNano())
	tc := md5.Sum([]byte(strconv.Itoa(rand.Intn(10000))))
	p.TrackingCode = string(tc[:])
	p.Status = operationsModel.PackStatusScheduled
	p.PackHash = t.BuildPackHash()

	/*
	*	marshal
	 */
	b, err := json.Marshal(&p)
	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
		return
	}

	producer.Publish(topic, b)
}
