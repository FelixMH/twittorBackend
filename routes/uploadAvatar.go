package routes

import (
	"github.com/FelixMH/tuitapp/bd"
	"github.com/FelixMH/tuitapp/models"

	"io"
	"net/http"
	"os"
	"strings"
)

/* UploadAvatar se usa para subir un avatar a la bd (para cada usuario) */
func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")
	if err != nil {
		http.Error(w,"Algun error ocurrió " +err.Error(), http.StatusBadRequest)
		return
	}
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/avatars/" + IDusuario + "." + extension
	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.Users
	var status bool

	user.Avatar = IDusuario + "." + extension
	status, err = bd.ModifyRegister(user, IDusuario)
	if err != nil || status == false {
		http.Error(w, "Error al grabar el avatar en la BD ! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}

