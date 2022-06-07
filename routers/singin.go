package routers

import (
	"encoding/json"
	"net/http"

	"github.com/luxarasis/twittor/bd"
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
	}
	if len(request.Password) < 6 {
		http.Error(w, "Bad request: Password length must be at least 6", http.StatusBadRequest)
	}

	_, found, _ := bd.FindUserByEmail(request.Email)
	if found {
		http.Error(w, "Bad request: User already exist", http.StatusBadRequest)
	}

	_, status, err := bd.SaveUser(request)
	if err != nil {
		http.Error(w, "Error saving User in database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if status == false {
		http.Error(w, "Error: couldn't save User on database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
