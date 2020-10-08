package routes

import (
	"encoding/json"
	"net/http"

	"github.com/FelixMH/tuitapp/bd"
	"github.com/FelixMH/tuitapp/models"
)

func SearchRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relation
	t.UsuarioID = IDusuario
	t.RelationID = ID

	var resp models.ResponseRelation

	status, err := bd.SearchRelation(t)
	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
