package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/asvins/router"
	"github.com/asvins/router/errors"
	"github.com/asvins/router/logger"
	"github.com/unrolled/render"
)

func DiscoveryHandler(w http.ResponseWriter, req *http.Request) errors.Http {
	prefix := strings.Join([]string{ServerConfig.Server.Addr, ServerConfig.Server.Port}, ":")

	//add discovery links here
	discoveryMap := map[string]string{"discovery": prefix + "/api/discovery"}

	rend.JSON(w, http.StatusOK, discoveryMap)
	return nil
}

func fetchRecipe(w http.ResponseWriter, r *http.Request) errors.Http {
	rend := render.New()
	treatmentId := r.URL.Query().Get("treatment_id")
	imgPath := r.URL.Path[len("/api/receipt"):]
	fmt.Println(imgPath + treatmentId)
	rend.JSON(w, 200, "{}")
	return nil
}

func DefRoutes() *router.Router {
	r := router.NewRouter()

	r.Handle("/api/discovery", router.GET, DiscoveryHandler, []router.Interceptor{})

	//MEDIC
	r.Handle("/api/medic/:medic_id/profile", router.GET, DiscoveryHandler, []router.Interceptor{})
	r.Handle("/api/medic/registration", router.POST, DiscoveryHandler, []router.Interceptor{}) // validar escopo de farmaceutico //DONE

	//TREATMENTS
	r.Handle("/api/treatments", router.GET, DiscoveryHandler, []router.Interceptor{})
	r.Handle("/api/treatments", router.PUT, DiscoveryHandler, []router.Interceptor{})
	r.Handle("/api/treatments", router.DELETE, DiscoveryHandler, []router.Interceptor{})
	r.Handle("/api/treatments", router.POST, DiscoveryHandler, []router.Interceptor{})
	r.Handle("/api/treatments/:treatment_id/ship", router.POST, DiscoveryHandler, []router.Interceptor{})

	//RECEIPT
	r.Handle("/api/receipt/:treatment_id", router.GET, fetchRecipe, []router.Interceptor{})               // TODO: VINIX
	r.Handle("/api/receipt/:treatment_id/validate", router.PUT, DiscoveryHandler, []router.Interceptor{}) // TODO: VINIX
	r.Handle("/api/receipt/:treatment_id", router.POST, DiscoveryHandler, []router.Interceptor{})         // TODO: VINIX

	//PATIENT
	r.Handle("/api/patient/:patient_id/feed", router.GET, DiscoveryHandler, []router.Interceptor{}) // TODO: VINIX

	//PHARMACIST
	r.Handle("/api/medications", router.GET, retreiveMedication, []router.Interceptor{})
	r.Handle("/api/medications/:id", router.PUT, updateMedication, []router.Interceptor{})
	r.Handle("/api/medications/:id", router.DELETE, deleteMedication, []router.Interceptor{})
	r.Handle("/api/medications", router.POST, insertMedication, []router.Interceptor{})

	r.Handle("/api/patient/registration", router.POST, DiscoveryHandler, []router.Interceptor{}) // validar escopo de medico
	// interceptors
	r.AddBaseInterceptor("/", logger.NewLogger())
	return r
}
