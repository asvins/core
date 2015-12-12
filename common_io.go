package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	authModels "github.com/asvins/auth/models"
	"github.com/asvins/common_io"
	"github.com/asvins/core/models"
	"github.com/asvins/utils/config"
	warehouseModels "github.com/asvins/warehouse/models"
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
	consumer.HandleTopic("activate_treatments", handleActivateTreatments)

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
		p := models.Patient{}
		p.ID = usr.ID
		p.Name = usr.FirstName + " " + usr.LastName
		p.Email = usr.Email

		err = p.Save(db)

		return

	case "medic":
		m := models.Medic{}
		m.ID = usr.ID
		m.Name = usr.FirstName + " " + usr.LastName
		m.Email = usr.Email

		err = m.Create(db)

		return

	case "pharmacist":
		p := models.Pharmacist{}
		p.ID = usr.ID
		p.Name = usr.FirstName + " " + usr.LastName
		p.Email = usr.Email

		err = p.Save(db)
		return
	}
	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
	}

}

func handleActivateTreatments(msg []byte) {
	ts := []models.Treatment{}
	err := json.Unmarshal(msg, &ts)
	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
		return
	}

	for _, t := range ts {
		t.Status = models.TREATMENT_STATUS_ACTIVE
		if err := t.Save(db); err != nil {
			fmt.Println("[ERROR] ", err.Error())
			return
		}
	}
}

/*
*	Senders
 */
func sendProductCreated(m *models.Medication) {
	topic, _ := common_io.BuildTopicFromCommonEvent(common_io.EVENT_CREATED, "product")
	p := warehouseModels.Product{}

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

func sendTreatmentCreated(t *models.Treatment) {
	topic, _ := common_io.BuildTopicFromCommonEvent(common_io.EVENT_CREATED, "treatment")

	/*
	*	marshal
	 */

	patient := models.Patient{}
	patient.ID = t.PatientId
	ps, err := patient.Retreive(db)
	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
		return
	}

	if len(ps) != 1 {
		fmt.Println("[FATAL] More then one patient with the same ID...")
		return
	}

	t.Email = ps[0].Email

	b, err := json.Marshal(&t)
	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
		return
	}

	producer.Publish(topic, b)
}

/*
*	Helper
 */

func generateTrackingCode() string {
	rand.Seed(time.Now().UTC().UnixNano())
	h := sha1.New()
	tc := strconv.Itoa(rand.Intn(10000))
	h.Write([]byte(tc))
	return hex.EncodeToString(h.Sum(nil))
}
