package routes

import (
	"encoding/json"
	"github.com/FelixMH/tuitapp/bd"
	"net/http"
	"strconv"
)

/* ReadTweets permite leer los tweets hechos por los usuarios de la app. */
func ReadTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parametro page", http.StatusBadRequest)
		return
	}

	page , err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Debe enviar el page con un valor mayor a cero", http.StatusBadRequest)
		return
	}

	pag := int64(page)

	resp, correct := bd.ReadTweet(ID, pag)
	if correct == false {
		http.Error(w, "Error al leer o cargar los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}