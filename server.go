package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ipochi/api-mock-example/handler"
	"github.com/ipochi/api-mock-example/implement"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	//sub := router.PathPrefix("/api/v1").Subrouter()

	server := handler.New(&implement.Implementor{})
	router.Methods("GET").Path("/api/v1/companies").HandlerFunc(server.GetCompanies)
	//sub.Methods("POST").Path("/companies").HandlerFunc(handler.SaveCompany)
	//sub.Methods("GET").Path("/companies/{name}").HandlerFunc(handler.GetCompany)
	//sub.Methods("PUT").Path("/companies/{name}").HandlerFunc(handler.UpdateCompany)
	//	sub.Methods("DELETE").Path("/companies/{name}").HandlerFunc(handler.DeleteCompany)

	fmt.Println("Staring api server")
	go http.ListenAndServe(":3000", router)
	fmt.Println("API server listening on port 3000")

	select {}
	//Routers(router, server)
	//StartServer(router)
}

// func Routers(router *mux.Router, server *handler.Server) {
// 	router.Methods("GET").Path("/companies").HandlerFunc(server.GetCompanies)
// }

// func StartServer(router *mux.Router) {
// 	log.Fatal(http.ListenAndServe(":3000", router))

// }
