package authcontrollers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/rakhazufar/go-jwt/pkg/config"
	"github.com/rakhazufar/go-jwt/pkg/models"
	"github.com/rakhazufar/go-jwt/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var userInput models.User
	//membaca json dari r.body
	decoder := json.NewDecoder(r.Body)
	//simpan data hasil decode ke userinput
	if err := decoder.Decode(&userInput); err != nil {
		message := map[string]string{"message":  "Gagal decode JSON"}
		utils.SendJSONResponse(w, http.StatusBadRequest, message)
		return
	}

	defer r.Body.Close()


	user, err := models.GetUserByUsername(userInput.Username);	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			message := map[string]string{"message": "Server error"}
			utils.SendJSONResponse(w, http.StatusInternalServerError, message)
            return
		}
	} else if user == nil {
		message := map[string]string{"message": "Username atau password salah"}
		utils.SendJSONResponse(w, http.StatusUnauthorized, message)
		return 
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		message := map[string]string{"message": "Username atau password salah"}
		utils.SendJSONResponse(w, http.StatusUnauthorized, message)
		return
	}

	expTimeToken := time.Now().Add(time.Hour * 24)

	claims := &config.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "go-jwt-postgres",
			ExpiresAt: jwt.NewNumericDate(expTimeToken),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		message := map[string]string{"message": "Server error"}
		utils.SendJSONResponse(w, http.StatusInternalServerError, message)
        return
	}

	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Path: "/",
		Value: token,
		HttpOnly: true,
	})
	message := map[string]string{"message": "success"}
	utils.SendJSONResponse(w, http.StatusOK, message)
}

func Register(w http.ResponseWriter, r *http.Request) {

	//buat variabel dengan struct sesuai data yang sudah dibuat
	var userInput models.User
	//membaca json dari r.body
	decoder := json.NewDecoder(r.Body)
	//simpan data hasil decode ke userinput
	if err := decoder.Decode(&userInput); err != nil {
		message := map[string]string{"message":  "Gagal decode JSON"}
		utils.SendJSONResponse(w, http.StatusBadRequest, message)
		return
	}

	defer r.Body.Close()

	//cek apakah username sudah ada di database
	if user, err := models.GetUserByUsername(userInput.Username); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			message := map[string]string{"message": "Server error"}
			utils.SendJSONResponse(w, http.StatusInternalServerError, message)
			return
		}
	} else if user != nil {
		message := map[string]string{"message": "User Already registered"}
		utils.SendJSONResponse(w, http.StatusConflict, message)
		return 
	}

	utils.HashPassword(&userInput)
	if err := models.CreateUser(&userInput); err != nil {
		message := map[string]string{"message": "Error creating user"}
		utils.SendJSONResponse(w, http.StatusInternalServerError, message)
        return
	}
	message := map[string]string{"message": "success"}
	utils.SendJSONResponse(w, http.StatusOK, message)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Path: "/",
		Value: "",
		HttpOnly: true,
		MaxAge: -1,
	})
	message := map[string]string{"message": "Logout Success"}
	utils.SendJSONResponse(w, http.StatusOK, message)
}