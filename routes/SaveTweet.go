package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/FelixMH/tuitapp/bd"
	"github.com/FelixMH/tuitapp/models"
)

/* SaveTweet función que guarda el tweet en el backend. */
func SaveTweet(w http.ResponseWriter, r *http.Request) {
	var m models.Tweets

	err := json.NewDecoder(r.Body).Decode(&m)

	register := models.Tweet{
		UserID: IDusuario,
		Message: m.Message,
		Date: time.Now(),
	}
	_ , status, err := bd.InsertTweet(register)
	if err != nil {
		http.Error(w,"Ocurrió un error al intentar enviar el registro.", 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el tweet", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
