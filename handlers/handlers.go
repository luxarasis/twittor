package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/luxarasis/twittor/middlew"
	"github.com/luxarasis/twittor/routers"
	"github.com/rs/cors"
)

/* Managers setea el puerto y pone a escuchar el Servidor */
func Managers() {
	router := mux.NewRouter()

	router.HandleFunc("/singin", middlew.CheckDB(routers.SingIn)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
