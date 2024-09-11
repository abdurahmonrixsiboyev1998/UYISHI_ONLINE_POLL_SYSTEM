package handlers

import (
	"encoding/json"
	"net/http"
	"online/internal/storage"
	"online/internal/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var creds map[string]string
	json.NewDecoder(r.Body).Decode(&creds)

	email, password := creds["email"], creds["password"]
	storage.Users[email] = password

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Registered successfully"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds map[string]string
	json.NewDecoder(r.Body).Decode(&creds)

	email, password := creds["email"], creds["password"]

	if storedPassword, exists := storage.Users[email]; exists && storedPassword == password {
		token, _ := utils.GenerateJWT(email)
		json.NewEncoder(w).Encode(map[string]string{"token": token})
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}
