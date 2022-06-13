package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/luxarasis/twittor/bd"
	"github.com/luxarasis/twittor/jwt"
	"github.com/luxarasis/twittor/models"
)

/* SingIn es la funcion para crear en la base de datos el registro del usuario */
func SingIn(w http.ResponseWriter, r *http.Request) {
	var request models.User

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Bad request: "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(request.Email) == 0 {
		http.Error(w, "Bad request: Email mustn't be empty", http.StatusBadRequest)
		return
	}
	if len(request.Password) < 6 {
		http.Error(w, "Bad request: Password length must be at least 6", http.StatusBadRequest)
		return
	}

	_, found, _ := bd.FindUserByEmail(request.Email)
	if found {
		http.Error(w, "Bad request: User already exist", http.StatusBadRequest)
		return
	}

	_, status, err := bd.SaveUser(request)
	if err != nil {
		http.Error(w, "Error saving User in database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(w, "Error: couldn't save User on database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var request models.User

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "User or passeord invalid: "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(request.Email) == 0 {
		http.Error(w, "Bad request: Email mustn't be empty", http.StatusBadRequest)
		return
	}

	document, found := bd.Login(request.Email, request.Password)
	if !found {
		http.Error(w, "User or passeord invalid", http.StatusBadRequest)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "Error: Genering JWT Token > "+err.Error(), http.StatusBadRequest)
		return
	}

	response := models.LoginResponseDTO{
		Token: jwtKey,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
