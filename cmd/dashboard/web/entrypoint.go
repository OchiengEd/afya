package web

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Start initiates the microservice
func Start() {
	config := loadConfiguration()

	router := mux.NewRouter()

	// patient handlers
	router.HandleFunc("/patient", registerPatient).Methods("POST")
	router.HandleFunc("/patient", getPatient).Methods("GET")
	router.HandleFunc("/patient", updatePatient).Methods("PUT")
	router.HandleFunc("/patient/list", listPatients).Methods("GET")
	router.HandleFunc("patient", deletePatient).Methods("DELETE")

	port := fmt.Sprintf(":%s", config.Port)
	if err := http.ListenAndServe(port, router); err != nil {
		panic(err)
	}
}
