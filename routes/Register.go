package routes

import (
	"encoding/json"
	"net/http"

	"github.com/FelixMH/tuitapp/bd"
	"github.com/FelixMH/tuitapp/models"
)

/*Register Ruta de registro para el api  */
func Register(w http.ResponseWriter, r *http.Request) {
	var t models.Users

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos		"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "El password debe tener al menos 6 caracteres", 400)
	}

	_, OK, _ := bd.CheckUsersTaken(t.Email)
	if OK == true {
		http.Error(w, "El usuario con este email ya existe", 400)
		return
	}

	_, status, err := bd.InsertRegister(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro de usuario"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha podido realizar el registro del usuario", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
