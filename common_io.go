package main

import (
	"encoding/json"
	"log"

	"github.com/asvins/auth/models"
	"github.com/asvins/common_io"
	"github.com/asvins/utils/config"
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
	defer producer.TearDown()

	/*
	*	Consumer
	 */
	consumer = common_io.NewConsumer(cfg)
	consumer.HandleTopic("user_created", userCreated)

	if err = consumer.StartListening(); err != nil {
		log.Fatal(err)
	}

	defer consumer.TearDown()

}

// HANDLERS

func userCreated(msg []byte) {
	var usr models.User
	json.Unmarshal(msg, &usr)
	switch usr.Scope {
	case "patient":
		p := Patient{Name: usr.FirstName + usr.LastName, Email: usr.Email}
		p.Create()
	case "medic":
		m := Medic{Name: usr.FirstName + usr.LastName, Email: usr.Email}
		m.Create()
	case "pharmacist":
		p := Pharmacist{Name: usr.FirstName + usr.LastName, Email: usr.Email}
		p.Create()
	default:
		break
	}
}

func userUpdated(msg []byte) {
	var usr models.User
	json.Unmarshal(msg, &usr)
	// update user: COMO PEGAR OS EMAILS ANTIGOS??
	switch usr.Scope {
	case "patient":
	case "medic":
	case "pharmacist":
	default:
		break

}

/*
*	Here can be added the handlers for kafka topics
 */
