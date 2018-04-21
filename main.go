package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"github.com/diachenko/backend/compute"
)

// Equation struct used for storing equation data
type Equation struct {
	ID        string `json:"id,omitempty"`
	EqStr     string `json:"eq,omitempty"`
	ResultStr string `json:"result,omitempty"`
}

// Eq Array of equations TODO: put into key-value store?
var Eq []Equation

// GetEquationEndpoint used for retriving old equation by ID
func GetEquationEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range Eq {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Equation{})
}

// GetEquationListEndpoint used for retriving old equations for FE
func GetEquationListEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(Eq)
}

// CreateEquationEndpoint used for creating new equation and getting result
func CreateEquationEndpoint(w http.ResponseWriter, req *http.Request) {
	var eq Equation
	_ = json.NewDecoder(req.Body).Decode(&eq)
	b := make([]byte, 16)
	rand.Read(b)
	var uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	eq.ID = uuid
	fmt.Println(eq.EqStr)
	var text = strings.Replace(eq.EqStr, " ", "", -1)
	fmt.Println(text)
	res, err := compute.Evaluate(text)
	eq.ResultStr = strconv.FormatFloat(res, 'f', 6, 64)
	fmt.Println(err)
	Eq = append(Eq, eq)
	json.NewEncoder(w).Encode(eq)
}

// DeleteEquationEndpoint used for deleting old equation by ID
func DeleteEquationEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range Eq {
		if item.ID == params["id"] {
			Eq = append(Eq[:index], Eq[index+1:]...)
			break
		}
	}
}
func main() {

	router := mux.NewRouter()
	router.HandleFunc("/calc", GetEquationListEndpoint).Methods("GET")
	router.HandleFunc("/calc/{id}", GetEquationEndpoint).Methods("GET")
	router.HandleFunc("/calc", CreateEquationEndpoint).Methods("POST")
	router.HandleFunc("/calc/{id}", DeleteEquationEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":1880", router))

}
