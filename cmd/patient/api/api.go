package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func routerSetup() *mux.Router {
	r := mux.NewRouter()

	// patient handlers
	r.HandleFunc("/patient", registerPatient).Methods("POST")
	r.HandleFunc("/patient", getPatient).Methods("GET")
	r.HandleFunc("/patient", updatePatient).Methods("PUT")
	r.HandleFunc("/patient/list", listPatients).Methods("GET")
	r.HandleFunc("patient", deletePatient).Methods("DELETE")

	return r
}

// Start brings up REST API
func Start() {
	routes := routerSetup()

	config := getConfiguration()
	port := fmt.Sprintf(":%s", config.Port)
	if err := http.ListenAndServe(port, routes); err != nil {
		log.Fatal(err)
	}
}
