package controller

import (
	"carakan-apps/Model"
	"carakan-apps/configuration"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
)

type MyClaim struct {
	User 				Model.UsersDTO	`json:"user"`
	jwt.StandardClaims
}

type authUser struct {
	UserName	string 	`json:"user_name"`
	Password	string	`json:"password"`
}

func Authentication(response http.ResponseWriter, request *http.Request) {
	contentLength := request.ContentLength

	//buat slice untuk menampung JSON data
	body := make([]byte, contentLength)
	request.Body.Read(body)
	request.Body.Close()

	//konversi JSON ke struct
	var user authUser
	err := json.Unmarshal(body, &user)

	if err != nil {
		log.Printf("[ERROR] : %s", err)
		response.WriteHeader(http.StatusBadGateway)
		response.Write([]byte("error"))
	}

	authUser, err := Model.GetAuthUser(&user.UserName, &user.Password)
	if err != nil {
		log.Printf("[ERROR] : %s", err)
		response.WriteHeader(http.StatusBadGateway)
		response.Write([]byte("error"))
	}

	claims := MyClaim{
		*authUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 0, 1).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key, err := token.SignedString([]byte(configuration.JWT_SECRET_KEY))
	if err != nil {
		log.Printf("[ERROR] : %s", err)
		response.WriteHeader(http.StatusBadGateway)
		response.Write([]byte("error"))
	}

	var message map[string]string = map[string]string{"token": key}
	resp, err := json.MarshalIndent(message, "", "\t\t")
	if err != nil {
		log.Printf("[ERROR] : %s", err)
		response.WriteHeader(http.StatusBadGateway)
		response.Write([]byte("error"))
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(resp)
}
