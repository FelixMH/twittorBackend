package routes

import (
	"net/http"
	"encoding/json"

	"github.com/FelixMH/tuitapp/bd"
)

/* SeeProfile permite obtener los datos del profile. */
func SeeProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el ID", http.StatusBadRequest)
		return
	}

	profile, err := bd.SearchProfile(ID)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar la busqueda..."+err.Error(), 400)
		return
	}
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(profile)
}
