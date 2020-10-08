package routes

import (
	"net/http"

	"github.com/FelixMH/tuitapp/bd"
	"github.com/FelixMH/tuitapp/models"
)

/* DeleteRelation realiza el borrado de la relacion de los usuarios. */
func DeleteRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Obligatorio enviar un parámetro ID. ", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UsuarioID = IDusuario
	t.RelationID = ID

	status, err := bd.DeleteRelation(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al Eliminar ....", http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado borrar la relación", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
