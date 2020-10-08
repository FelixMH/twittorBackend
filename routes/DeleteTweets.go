package routes

import (
	"net/http"
	"github.com/FelixMH/tuitapp/bd"
)

/* DeleteTweets borra los tweets de un usuario. */
func DeleteTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	err := bd.DeleteTweets(ID, IDusuario)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar el tweet"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}
