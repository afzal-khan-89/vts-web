package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func userRouter(r *mux.Router) {
	r.HandleFunc("/adduser", addUser).Methods("GET")
}
func addUser(w http.ResponseWriter, r *http.Request) {

}
