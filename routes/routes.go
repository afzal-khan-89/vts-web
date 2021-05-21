package routes

import "github.com/gorilla/mux"

func RoutSetup() {

}
func NewRouter() *mux.Router {
	r := mux.NewRouter()

	//fileServer(r) // Fileserver to serve static files
	homeRout(r)
	reportsRouter(r) // Index handler
	userRouter(r)    // Other domain/business logic scoped handler

	return r
}
