package routes

import (
	"errors"
	"strings"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/FelixMH/tuitapp/bd"
	"github.com/FelixMH/tuitapp/models"
)

/* Email es el email del usuario usado en todos los ENDPOINTS */
var Email string

/* IDusuario es el ID del usuario que recibe los endpoints */
var IDusuario string

/* TokenProcess Se procesa el token para extraer todos sus valores. */
func TokenProcess(token string) (*models.Claim, bool, string, error) {
	key := []byte("tuitapp")
	claims := &models.Claim{

	}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	token = strings.TrimSpace(splitToken[1])

	kf := func(tkn *jwt.Token)(interface{}, error) {
		return key, nil
	}

	newToken, err := jwt.ParseWithClaims(token,claims, kf)

	if err == nil {
		_, OK, _ := bd.CheckUsersTaken(claims.Email)
		if OK == true {
			Email = claims.Email
			IDusuario = claims.ID.Hex()
		}
		return claims, OK, IDusuario, nil
	}

	if !newToken.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err
}