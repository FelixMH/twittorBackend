package routes

import (
	"net/http"
	"github.com/FelixMH/tuitapp/bd"
	"github.com/FelixMH/tuitapp/models"
)

/* SaveRelation se realiza el guardado de la relacion de los usuarios. */
func SaveRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w,"El parametro ID debe de ser obligatorio.", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UsuarioID = IDusuario
	t.RelationID = ID

	status, err := bd.InsertRelation(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al guardar ... ", http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar la relación", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}