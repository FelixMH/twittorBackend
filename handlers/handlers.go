package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/FelixMH/tuitapp/middlewares"
	"github.com/FelixMH/tuitapp/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores maneja las rutas de la api con MUX  */
func Manejadores() {

	router := mux.NewRouter()

	/* RUTAS */
	router.HandleFunc("/registro", middlewares.CheckDB(routes.Register)).Methods("POST")
	router.HandleFunc("/login", middlewares.CheckDB(routes.Login)).Methods("POST")
	router.HandleFunc("/ver-perfil", middlewares.CheckDB(middlewares.ValidateToken(routes.SeeProfile))).Methods("GET")
	router.HandleFunc("/modificar-perfil", middlewares.CheckDB(middlewares.ValidateToken(routes.ModifyProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlewares.CheckDB(middlewares.ValidateToken(routes.SaveTweet))).Methods("POST")
	router.HandleFunc("/ver-tweets", middlewares.CheckDB(middlewares.ValidateToken(routes.ReadTweets))).Methods("GET")
	router.HandleFunc("/delete-tweet", middlewares.CheckDB(middlewares.ValidateToken(routes.DeleteTweets))).Methods("DELETE")

	router.HandleFunc("/subir-avatar", middlewares.CheckDB(middlewares.ValidateToken(routes.UploadAvatar))).Methods("POST")
	router.HandleFunc("/obtener-avatar", middlewares.CheckDB(routes.GettingAvatar)).Methods("GET")
	router.Handle("/subir-banner", middlewares.CheckDB(middlewares.ValidateToken(routes.UploadBanner))).Methods("POST")
	router.HandleFunc("/obtener-banner", middlewares.CheckDB(routes.GettingBanner)).Methods("GET")

	router.HandleFunc("/alta-relacion", middlewares.CheckDB(middlewares.ValidateToken(routes.SaveRelation))).Methods("POST")
	router.HandleFunc("/baja-relacion", middlewares.CheckDB(middlewares.ValidateToken(routes.DeleteRelation))).Methods("DELETE")
	router.HandleFunc("/consultar-relacion", middlewares.CheckDB(middlewares.ValidateToken(routes.SearchRelation))).Methods("GET")

	router.HandleFunc("/usuarios", middlewares.CheckDB(middlewares.ValidateToken(routes.UsersList))).Methods("GET")
	router.HandleFunc("/leer-tweets-followers", middlewares.CheckDB(middlewares.ValidateToken(routes.ReadTweetsFollowers))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
