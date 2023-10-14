package handler

import (
	"be/dto"
	"encoding/json"
	"net/http"
	"strings"
)

func RetrieveUsers(w http.ResponseWriter, r *http.Request, users dto.Users) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid Method",
		})
	}

	data := make([]interface{}, len(users))
	for i, v := range users {
		data[i] = v
	}

	json.NewEncoder(w).Encode(dto.Response{
		Data:    data,
		Status:  http.StatusOK,
		Message: "success",
	})
}

func ValidateUser(w http.ResponseWriter, r *http.Request, users dto.Users) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid Method",
		})
	}

	var creds dto.Login
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid JSON",
		})
		return
	}

	if strings.TrimSpace(creds.Password) == "" || strings.TrimSpace(creds.Email) == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.Response{
			Status:  http.StatusBadRequest,
			Message: "Email/Password cannot be empty",
		})
		return
	}

	response, err := users.Find(creds.Email, "email")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.Response{
			Status:  http.StatusInternalServerError,
			Message: "Email is not registered!",
		})
		return
	}

	data := []interface{}{response}

	json.NewEncoder(w).Encode(dto.Response{
		Data:    data,
		Status:  http.StatusOK,
		Message: "success",
	})
}
