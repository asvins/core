package main

import (
	"net/http"
	"strings"

	"github.com/asvins/router"
	"github.com/asvins/router/errors"
	"github.com/asvins/router/logger"
)

func DiscoveryHandler(w http.ResponseWriter, req *http.Request) errors.Http {
	prefix := strings.Join([]string{ServerConfig.Server.Addr, ServerConfig.Server.Port}, ":")

	//add discovery links here
	discoveryMap := map[string]string{"discovery": prefix + "/api/discovery"}

	rend.JSON(w, http.StatusOK, discoveryMap)
	return nil
}

func DefRoutes() *router.Router {
	r := router.NewRouter()

	r.Handle("/api/discovery", router.GET, DiscoveryHandler, []router.Interceptor{})

	//MEDIC
	r.Handle("/api/medic/:medic_id/profile", router.GET, DiscoveryHandler, []router.Interceptor{})
	r.Handle("/api/medic/:id", router.PUT, handleMedicUpdate, []router.Interceptor{})
	r.Handle("/api/medic", router.GET, handleMedicRetrieve, []router.Interceptor{})

	//TREATMENTS
	r.Handle("/api/treatments", router.GET, retreiveTreatments, []router.Interceptor{})
	r.Handle("/api/treatments/:id", router.PUT, updateTreatment, []router.Interceptor{})
	r.Handle("/api/treatments/:id", router.DELETE, deleteTreatment, []router.Interceptor{})
	r.Handle("/api/treatments", router.POST, insertTreatment, []router.Interceptor{})
	r.Handle("/api/treatments/:treatment_id/validate", router.POST, validataTreatment, []router.Interceptor{})

	//RECEIPT
	r.Handle("/api/receipt/:prescription_id", router.GET, fetchRecipe, []router.Interceptor{})
	r.Handle("/api/receipt/:prescription_id/validate", router.PUT, validateRecipe, []router.Interceptor{})
	r.Handle("/api/receipt/:prescription_id", router.POST, uploadRecipe, []router.Interceptor{})

	//PATIENT
	r.Handle("/api/patient/:patient_id/feed", router.GET, handleGetFeed, []router.Interceptor{})
	r.Handle("/api/patient/:id", router.PUT, handlePatientUpdate, []router.Interceptor{})
	r.Handle("/api/patient", router.GET, handlePatientRetrieve, []router.Interceptor{})

	//PHARMACIST
	r.Handle("/api/pharmacist/:id", router.GET, handlePharmacistUpdate, []router.Interceptor{})
	r.Handle("/api/pharmacist", router.GET, handlePharmacistRetrieve, []router.Interceptor{})

	//MEDICATIONS
	r.Handle("/api/medications", router.GET, retreiveMedication, []router.Interceptor{})
	r.Handle("/api/medications/:id", router.PUT, updateMedication, []router.Interceptor{})
	r.Handle("/api/medications/:id", router.DELETE, deleteMedication, []router.Interceptor{})
	r.Handle("/api/medications", router.POST, insertMedication, []router.Interceptor{})

	// interceptors
	r.AddBaseInterceptor("/", logger.NewLogger())
	return r
}
