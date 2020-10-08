package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/FelixMH/tuitapp/models"
)
/*GenerateJWT genera el encriptado con JWT*/
func GenerateJWT( t models.Users ) (string, error) {
	key := []byte("tuitapp")

	payload := jwt.MapClaims{
		"name":		t.Name,
		"lastName": t.LastName,
		"birthday":	t.Birthday,
		"email":	t.Email,
		"webSite":	t.WebSite,
		"avatars":	t.Avatar,
		"_id":		t.ID.Hex(),
		"exp":		time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}
