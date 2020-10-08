package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/FelixMH/tuitapp/bd"
)

func UsersList(w http.ResponseWriter, r *http.Request) {

	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagtmp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Se debe enviar el parametro con un valor mayor a 0 "+err.Error(), http.StatusBadRequest)
		return
	}

	pag := int64(pagtmp)

	results, status := bd.ReadUsersAll(IDusuario, pag, search, typeUser)
	if status == false {
		http.Error(w, "Error al leer los usuarios "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}
