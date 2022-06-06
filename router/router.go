package router

import (
	"fmt"
	"net/http"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/controller"
	"github.com/SamuelSalas/2022Q2GO-Bootcamp/repository"
	"github.com/SamuelSalas/2022Q2GO-Bootcamp/service"
	"github.com/gorilla/mux"
)

func Router(router *mux.Router) {
	csvController := controller.NewCsvController(service.NewCsvService(repository.NewCharacterClientRepository()))
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running...")
	}).Methods("GET")
	router.HandleFunc("/generateCsv", csvController.GetApiDataCsv).Methods("GET")
	router.HandleFunc("/sendCSVFile", csvController.PostCSVFile).Methods("POST")
}
