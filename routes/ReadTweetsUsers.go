package routes
import (
"encoding/json"
"net/http"
"strconv"

"github.com/FelixMH/tuitapp/bd"
)

/*LeoTweetsSeguidores lee los tweets de todos nuestros seguidores */
func ReadTweetsFollowers(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página como entero mayor a 0", http.StatusBadRequest)
		return
	}

	resp, OK := bd.ReadTweetsFollowers(IDusuario, page)
	if OK == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}