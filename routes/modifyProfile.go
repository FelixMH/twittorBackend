package routes

import (
	"encoding/json"
	"net/http"

	"github.com/FelixMH/tuitapp/bd"
	"github.com/FelixMH/tuitapp/models"
)

/* ModifyProfile modifica el perfil del usuario enviado al endpoint */
func ModifyProfile (w http.ResponseWriter, r *http.Request) {
	var t models.Users

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(), 400)
		return
	}
	var status bool
	status, err = bd.ModifyRegister(t, IDusuario)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar modificar tus datos. Reintenta nuevamente"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se han podido modificar los datos"+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}