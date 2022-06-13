package jwt

import (
	"errors"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/luxarasis/twittor/bd"
	"github.com/luxarasis/twittor/models"
)

var Email string
var IDUser string

var myKey = []byte("secret")

func GenerateJWT(user models.User) (string, error) {

	payload := jwt.MapClaims{
		"email":     user.Email,
		"name":      user.Name,
		"lastname":  user.LastName,
		"birthday":  user.Birthday,
		"biography": user.Biography,
		"location":  user.Location,
		"website":   user.WebSite,
		"_id":       user.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}

func ProccesJWTToken(token string) (*models.Claim, bool, string, error) {
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("JWT Token invalid")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		_, found, _ := bd.FindUserByEmail(claims.Email)

		if found {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}

		return claims, found, IDUser, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("JWT Token invalid")
	}

	return claims, false, string(""), err
}
