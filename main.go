package main

import (
	"fmt"
	"net/http"

	"github.com/afzal-khan-89/vts-web/routes"
	"github.com/afzal-khan-89/vts-web/utils"
)

func main() {
	fmt.Println("hello go ")
	routes.RoutTest()

	utils.LoadTemplate()
	r := routes.NewRouter()
	http.Handle("/", r)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":8080", nil)
}
