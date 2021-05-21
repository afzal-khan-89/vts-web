package routes

import (
	"fmt"
	"net/http"

	"github.com/afzal-khan-89/vts-web/utils"
	"github.com/gorilla/mux"
)

func RoutTest() {
	fmt.Println("rout package is working ....")
}

func reportsRouter(r *mux.Router) {
	r.HandleFunc("/reports", rawDataReports).Methods("GET")
	r.HandleFunc("/triphistory", tripHistoryReporst).Methods("GET")
}

// Handlers

func rawDataReports(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "reports.html", nil)
}
func tripHistoryReporst(w http.ResponseWriter, r *http.Request) {

}
