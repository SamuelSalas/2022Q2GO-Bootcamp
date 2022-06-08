package router

import (
	"fmt"
	"net/http"

	"github.com/SamuelSalas/2022Q2GO-Bootcamp/controller"
	"github.com/SamuelSalas/2022Q2GO-Bootcamp/repository"
	"github.com/SamuelSalas/2022Q2GO-Bootcamp/service"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func Router(router *mux.Router) {
	csvController := controller.NewCsvController(service.NewCsvService(repository.NewCsvRepository()))
	router.Handle("/swagger.yml", http.FileServer(http.Dir("./")))
	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yml"}
	sh := middleware.SwaggerUI(opts, nil)
	router.Handle("/docs", sh)
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running...")
	}).Methods("GET")
	router.HandleFunc("/generateCsv", csvController.GetApiDataCsv).Methods("GET")
	router.HandleFunc("/readCsvFile", csvController.GetCSVFileData).Methods("GET")
	router.HandleFunc("/readCsvFileWorkerPool", csvController.GetCSVFileDataWorkerPool).Methods("GET")
}
