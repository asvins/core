package models

import (
	"reflect"
	"strconv"
	"time"

	om "github.com/asvins/operations/models"
	sm "github.com/asvins/subscription/models"
	"github.com/jinzhu/gorm"
)

type FeedEvent struct {
	gorm.Model
	PatientID   int    `json:"patient_id" gorm:"column:patient_id"`
	Title       string `json:"title"`
	Tags        string `json:"tags"`
	Description string `json:"desc"`
	Hypermidia  string `json:"link"`
}

func FindFeedEvents(from time.Time, patientID int, db *gorm.DB) ([]FeedEvent, error) {
	var fes []FeedEvent
	err := db.Where("patient_id = ? AND updated_at > ?", patientID, from).Find(&fes).Error
	return fes, err
}

func (e *FeedEvent) Create(db *gorm.DB) error {
	return db.Create(e).Error
}

func (e *FeedEvent) Save(db *gorm.DB) error {
	return db.Save(e).Error
}

func (e *FeedEvent) Delete(db *gorm.DB) error {
	return db.Delete(e).Error
}

// LOGIC
func NewEvent(i interface{}) *FeedEvent {
	e := &FeedEvent{}
	switch reflect.TypeOf(i).Name() {
	case "Subscription":
		s, _ := i.(sm.Subscription)
		e.Title = "Assinatura Atualizada"
		e.Description = "Seus dados de pagamento foram atualizados. Isso pode significar que um pagamento foi realizado, ou que um endereço de entrega foi modificado."
		e.Tags = "subscription"
		e.PatientID, _ = strconv.Atoi(s.Owner)
		break

	case "Patient":
		p, _ := i.(Patient)
		if p.CPF == "" {
			e.Title = "Bem vindo ao Asvins!"
			e.Description = "Cadastro realizado com sucesso!<br>Não esqueça de terminar seu cadastro!"
			e.Tags = "profile"
			e.PatientID = p.ID
		} else {
			e.Title = "Dados Atualizados"
			e.Description = "Os dados de sua conta foram atualizados com sucesso!"
			e.Tags = "profile"
			e.PatientID = p.ID
		}
		break

	case "Box":
		box, _ := i.(om.Box)
		switch box.Status {
		case om.BOX_DELIVERED:
			e.Title = "Pedido Entregue"
			e.Description = "Seu pedido de " + time.Unix(box.StartDate, 0).String() + " até " + time.Unix(box.EndDate, 0).String() + " já está saiu para a entrega!"
			break

		case om.BOX_SHIPED:
			e.Title = "Pedido Enviado"
			e.Description = "Seu pedido de " + time.Unix(box.StartDate, 0).String() + " até " + time.Unix(box.EndDate, 0).String() + " já foi enviado pela transportadora!"
			break

		case om.BOX_SCHEDULED:
			e.Title = "Pedido Agendado"
			e.Description = "Seu pedido de " + time.Unix(box.StartDate, 0).String() + " até " + time.Unix(box.EndDate, 0).String() + " já foi agendado."
			break

		case om.BOX_PENDING:
			e.Title = "Pedido Agendado"
			e.Description = "Seu pedido de " + time.Unix(box.StartDate, 0).String() + " até " + time.Unix(box.EndDate, 0).String() + " Foi recebido. Termine seu cadastro para dar continuidade ao processo de envio!."
		}
		e.Title = "Atualizações do Envio"
		e.Tags = "shipment"
		e.PatientID = box.PatientId
		break

	case "Subscriber":
		break
	default:
		return nil
	}

	return e
}
