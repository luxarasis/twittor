package routers

import (
	"encoding/json"
	"net/http"

	"github.com/luxarasis/twittor/bd"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Bad request: id RequestParam is requeride", http.StatusBadRequest)
		return
	}

	user, err := bd.FindUserByID(ID)
	if err != nil {
		http.Error(w, "User not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
