package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ipochi/api-mock-example/db"
	"github.com/ipochi/api-mock-example/model"
)

type Server struct {
	impl model.Functions
}

func New(impl model.Functions) *Server {

	return &Server{
		impl: impl,
	}
}
func (s *Server) GetCompanies(w http.ResponseWriter, _ *http.Request) {

	fmt.Println("Hey we got a call")
	companies, err := s.impl.GetCompanies()
	fmt.Println("Hey hey --- ", companies, "and err ---", err)
	bytes, err := json.Marshal(companies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	writeJsonResponse(w, bytes)
}

func GetCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	com, ok := db.FindBy(name)
	if !ok {
		http.NotFound(w, r)
		return
	}

	bytes, err := json.Marshal(com)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	writeJsonResponse(w, bytes)
}

// func SaveCompany(w http.ResponseWriter, r *http.Request) {
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// 	com := new(model.Company)
// 	err = json.Unmarshal(body, com)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// 	db.Save(com.Name, com)

// 	w.Header().Set("Location", r.URL.Path+"/"+com.Name)
// 	w.WriteHeader(http.StatusCreated)
// }

// func UpdateCompany(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	name := vars["name"]

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// 	com := new(model.Company)
// 	err = json.Unmarshal(body, com)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// 	db.Save(name, com)
// }

func DeleteCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	db.Remove(name)
	w.WriteHeader(http.StatusNoContent)
}

func writeJsonResponse(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
}
