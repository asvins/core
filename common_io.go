package main

import (
	"encoding/json"
	"fmt"
	"log"

	authModels "github.com/asvins/auth/models"
	"github.com/asvins/common_io"
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

	// users (patient, medic, pharmaceutic)
	consumer.HandleTopic("user_created", handleUserCreated)

	// medication TODO
	consumer.HandleTopic("medication_created", nil)
	consumer.HandleTopic("medication_updated", nil)
	consumer.HandleTopic("medication_deleted", nil)

	if err = consumer.StartListening(); err != nil {
		log.Fatal(err)
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
		p.Name = usr.FirstName + " " + usr.LastName
		p.Email = usr.Email

		err = p.Save()

		return

	case "medic":
		m := Medic{}
		m.Name = usr.FirstName + " " + usr.LastName
		m.Email = usr.Email

		err = m.Create()

		return

	case "pharmaceutic":
		p := Pharmaceutic{}
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
