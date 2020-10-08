package routes

import (
	"io"
	"net/http"
	"os"
	"github.com/FelixMH/tuitapp/bd"
)

/* GettingBanner obtiene el Banner de cada usuario. */
func GettingBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w,"Debe enviar el parametro ID ", http.StatusBadRequest)
		return
	}

	profile, err := bd.SearchProfile(ID)
	if err != nil {
		http.Error(w, "Usuario no encontrado "+err.Error(), http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/banners/"+profile.Banner)
	if err != nil {
		http.Error(w, "Imagen no encontrada "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error al copiar la imagen "+err.Error(), http.StatusBadRequest)
		return
	}

}