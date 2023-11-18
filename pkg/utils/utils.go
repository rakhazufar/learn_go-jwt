package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rakhazufar/go-jwt/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(input *models.User)  {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal("Gagal hashsing password")
	}

	input.Password = string(hashPass)
}

func SendJSONResponse(w http.ResponseWriter, status int, payload interface{}) {
	response, _ := json.Marshal(payload)
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(status)
    w.Write(response)
}

