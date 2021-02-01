package api

import (
	"fmt"
	"net/http"
)

func registerPatient(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "register patient")
}

func getPatient(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "get patient")
}

func updatePatient(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "update patient")
}

func deletePatient(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "delete patient")
}

func listPatients(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "list patients")
}
