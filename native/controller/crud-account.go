package controller

import (
	"carakan-apps/Model"
	"encoding/json"
	"log"
	"net/http"
)

type AccountRequestId struct {
	AccountId string `json:"account_id"`
}

type AccountInsert struct {
	FirstName 			string				`json:"first_name"`
	LastName			string				`json:"last_name"`
	Email       		string				`json:"email"`
	Password 			string				`json:"password"`
}

type AccountUpdate struct {
	AccountId 			string 				`json:"account_id"`
	FirstName 			string				`json:"first_name"`
	LastName			string				`json:"last_name"`
}

func GetAllAccount(response http.ResponseWriter, request *http.Request) {
	accounts, err := Model.GetAllAccount()
	if err != nil {
		log.Printf("[ERROR] : %s", err)
		response.WriteHeader(http.StatusBadGateway)
		response.Write([]byte("error"))
	}

	resp, err := json.MarshalIndent(accounts, "", "\t\t")
	if err != nil {
		log.Printf("[ERROR] : %s", err)
		response.WriteHeader(http.StatusBadGateway)
		response.Write([]byte("error"))
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(resp)
}

func GetAccountById(response http.ResponseWriter, request *http.Request) {
	contentLength := request.ContentLength

	//buat slice untuk menampung JSON data
	body := make([]byte, contentLength)
	request.Body.Read(body)

	//konversi JSON ke struct
	var user AccountRequestId
	err := json.Unmarshal(body, &user)
	if err != nil {
		log.Printf("[ERROR] : %s", err)
		response.WriteHeader(http.StatusBadGateway)
		response.Write([]byte("error"))
	}

	account, err := Model.GetAccountById(user.AccountId)
	if err != nil {
		log.Printf("[ERROR] : %s", err)
		response.WriteHeader(http.StatusBadGateway)
		response.Write([]byte("error"))
	}

	resp, err := json.MarshalIndent(*account, "", "\t\t")
	if err != nil {
		log.Printf("[ERROR] : %s", err)
		response.WriteHeader(http.StatusBadGateway)
		response.Write([]byte("error"))
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(resp)
}

func InsertAccount(response http.ResponseWriter, request *http.Request) {
	var message map[string]string = map[string]string{"message": "success"}
	contentLength := request.ContentLength

	//buat slice untuk menampung JSON data
	body := make([]byte, contentLength)
	request.Body.Read(body)

	//konversi JSON ke struct
	var user AccountInsert
	err := json.Unmarshal(body, &user)
	if err != nil {
		log.Printf("[ERROR] : %s", err)
		response.WriteHeader(http.StatusBadGateway)
		response.Write([]byte("error"))
	}

	account := Model.Account{
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		Password: user.Password,
	}

	err = Model.InsertAccount(&account)
	if err != nil {
		log.Printf("[ERROR] : %s", err)
		response.WriteHeader(http.StatusBadGateway)
		response.Write([]byte("error"))
	}

	resp, err := json.MarshalIndent(message, "", "\t\t")
	if err != nil {
		log.Printf("[ERROR] : %s", err)
		response.WriteHeader(http.StatusBadGateway)
		response.Write([]byte("error"))
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(resp)
}

func UpdateAccount(response http.ResponseWriter, request *http.Request) {
	var message map[string]string = map[string]string{"message": "success"}
	contentLength := request.ContentLength

	//buat slice untuk menampung JSON data
	body := make([]byte, contentLength)
	request.Body.Read(body)

	//konversi JSON ke struct
	var user AccountUpdate
	err := json.Unmarshal(body, &user)
	if err != nil {
		log.Printf("[ERROR] : %s", err)
		response.WriteHeader(http.StatusBadGateway)
		response.Write([]byte("error"))
	}

	err = Model.UpdateUsernameAccount(user.AccountId, user.FirstName, user.LastName)
	if err != nil {
		log.Printf("[ERROR] : %s", err)
		response.WriteHeader(http.StatusBadGateway)
		response.Write([]byte("error"))
	}

	resp, err := json.MarshalIndent(message, "", "\t\t")
	if err != nil {
		log.Printf("[ERROR] : %s", err)
		response.WriteHeader(http.StatusBadGateway)
		response.Write([]byte("error"))
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(resp)
}

func DeleteAccount(response http.ResponseWriter, request *http.Request){
	var message map[string]string = map[string]string{"message": "success"}
	contentLength := request.ContentLength

	//buat slice untuk menampung JSON data
	body := make([]byte, contentLength)
	request.Body.Read(body)

	//konversi JSON ke struct
	var user AccountRequestId
	err := json.Unmarshal(body, &user)
	if err != nil {
		log.Printf("[ERROR] : %s", err)
		response.WriteHeader(http.StatusBadGateway)
		response.Write([]byte("error"))
	}

	err = Model.DeleteAccount(user.AccountId)
	if err != nil {
		log.Printf("[ERROR] : %s", err)
		response.WriteHeader(http.StatusBadGateway)
		response.Write([]byte("error"))
	}

	resp, err := json.MarshalIndent(message, "", "\t\t")
	if err != nil {
		log.Printf("[ERROR] : %s", err)
		response.WriteHeader(http.StatusBadGateway)
		response.Write([]byte("error"))
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(resp)
}