package routes

import (
	"encoding/json"
	"github.com/FelixMH/tuitapp/jwt"
	"net/http"
	"time"

	"github.com/FelixMH/tuitapp/bd"
	"github.com/FelixMH/tuitapp/models"
)

/*Login Realiza el login */
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.Users

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o contraseña válido"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w,"El email es requerido", 400)
		return
	}

	document, exists := bd.CheckLogin(t.Email, t.Password)

	if exists == false {
		http.Error(w,"Usuario y contraseña inválidos", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "Ocurrió un error : "+err.Error() , 400)
		return
	}

	resp := models.LoginResponse {
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w,&http.Cookie{
		Name: "token",
		Value: jwtKey,
		Expires: expirationTime,
	})

}