package main

import (
	"be/handler"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	route := mux.NewRouter()
	route.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		handler.RetrieveUsers(w, r, Users)
	})
	route.HandleFunc("/user/validate", func(w http.ResponseWriter, r *http.Request) {
		handler.ValidateUser(w, r, Users)
	})

	corsHandler := cors.Default().Handler(route)

	server := &http.Server{
		Addr:    PORT,
		Handler: corsHandler,
	}

	fmt.Printf("🚀 Application is running on: http://localhost%v \n", PORT)

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
