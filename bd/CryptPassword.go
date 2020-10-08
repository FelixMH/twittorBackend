package bd

import (
	"golang.org/x/crypto/bcrypt"
)

/*CryptPassword encripta la password de los registros...  */
func CryptPassword(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)

	if err != nil {
		return "", err
	}

	return string(bytes), err
}
