package routes

import (
	"github.com/FelixMH/tuitapp/bd"
	"github.com/FelixMH/tuitapp/models"
	"io"
	"net/http"
	"os"
	"strings"
)

/* UploadBanner se usa para subir un banner a la bd (para cada usuario) */
func UploadBanner(w http.ResponseWriter, r *http.Request) {
	file, handler , err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo = "uploads/banners/"+ IDusuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w,"Error al subir la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w,"Error al copiar la imagen ! ", http.StatusBadRequest)
		return
	}

	var user models.Users
	var status bool

	user.Banner = IDusuario + "." + extension
	status, err = bd.ModifyRegister(user, IDusuario)

	if err != nil || status == false {
		http.Error(w, "Error al grabar el banner en la BD ! ", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}

