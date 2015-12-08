package main

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

func FindFeedEvents(from time.Time, patientID int) ([]FeedEvent, error) {
	var fes []FeedEvent
	err := db.Where("patient_id = ? AND updated_at > ?", patientID, from).Find(&fes).Error
	return fes, err
}

func (e *FeedEvent) Create() error {
	return db.Create(e).Error
}

func (e *FeedEvent) Save() error {
	return db.Save(e).Error
}

func (e *FeedEvent) Delete() error {
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
	case "Patient":
		p, _ := i.(Patient)
		e.Title = "Dados Atualizados"
		e.Description = "Os dados de sua conta foram atualizados com sucesso!"
		e.Tags = "profile"
		e.PatientID = p.ID
	case "Pack":
		p, _ := i.(om.Pack)
		switch p.Status {
		case om.PackStatusDelivered:
			e.Title = "Pedido Entregue"
			e.Description = "Seu pedido de " + p.To.String() + " até " + p.From.String() + " já está saiu para a entrega!"
		case om.PackStatusShipped:
			e.Title = "Pedido Enviado"
			e.Description = "Seu pedido de " + p.To.String() + " até " + p.From.String() + " já foi enviado pela transportadora!"
		case om.PackStatusScheduled:
			e.Title = "Pedido Agendado"
			e.Description = "Seu pedido de " + p.To.String() + " até " + p.From.String() + " já foi agendado."
		case om.PackStatusOnProduction:
			e.Title = "Pedido em Produção"
			e.Description = "Seu pedido de " + p.To.String() + " até " + p.From.String() + " já está sendo produzido."
		}
		e.Title = "Atualizações do Envio"
		e.Tags = "shipment"
		e.PatientID, _ = strconv.Atoi(p.Owner)
		if p.TrackingCode != "" {
			e.Description += "Código de rastreio: " + p.TrackingCode
		}
	default:
		return nil
	}

	return e
}
