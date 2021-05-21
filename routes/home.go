package routes

import (
	"net/http"
	"utils"

	"github.com/afzal-khan-89/vts-web/utils"
	"github.com/gorilla/mux"
)

func homeRout(r *mux.Router) {
	r.HandleFunc("/", welcomeHome).Methods("GET")
}

func welcomeHome(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "index.html", nil)

}
