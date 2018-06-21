package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ipochi/api-mock-example/handler"
	"github.com/ipochi/api-mock-example/implement"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("/api/v1").Subrouter()
	//sub.Methods("POST").Path("/companies").HandlerFunc(handler.SaveCompany)
	sub.Methods("GET").Path("/companies/{name}").HandlerFunc(handler.GetCompany)
	//sub.Methods("PUT").Path("/companies/{name}").HandlerFunc(handler.UpdateCompany)
	sub.Methods("DELETE").Path("/companies/{name}").HandlerFunc(handler.DeleteCompany)

	server := handler.New(&implement.Implementor{})
	Routers(router, server)
	StartServer(router)
}

func Routers(router *mux.Router, server *handler.Server) {
	router.Methods("GET").Path("/companies").HandlerFunc(server.GetCompanies)
}

func StartServer(router *mux.Router) {
	log.Fatal(http.ListenAndServe(":3000", router))

}
