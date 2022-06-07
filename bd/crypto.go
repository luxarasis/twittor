package bd

import "golang.org/x/crypto/bcrypt"

/* EncryptPassword es la rutina que me permite cifrar la password */
func EncryptPassword(pass string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)

	return string(bytes), err
}
