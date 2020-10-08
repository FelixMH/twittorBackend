package middlewares

import (
	"net/http"

	"github.com/FelixMH/tuitapp/bd"
)

/*CheckDB Permite conocer el estado de la DB  */
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "conexión perdida con mi base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
