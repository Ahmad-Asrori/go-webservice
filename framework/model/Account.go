package model

import (
	"krama/util"
	"log"
	"math/rand"
	"time"
)

type Account struct {
	AccountId 			string				`json:"account_id"`
	FirstName 			string				`json:"first_name"`
	LastName			string				`json:"last_name"`
	Email       		string				`json:"email"`
	Password 			string				`json:"-"`
	PasswordResetCode	string				`json:"-"`
	PasswordResetExpiry time.Time			`json:"-"`
	AccountCreateTime 	time.Time			`json:"-"`
	AccountUpdateTime	time.Time			`json:"-"`
	IsActiveAccount		bool				`json:"-"`
	ActivationCode		string				`json:"-"`
}

type AccountDTO struct {
	AccountId 	string 		`json:"account_id"`
	FirstName 	string 		`json:"first_name"`
	LastName  	string 		`json:"last_name"`
	Email     	string 		`json:"email"`
}

type AccountId struct {
	AccountId string		`json:"account_id"`
}

type AccountInsert struct {
	FirstName 			string				`json:"first_name"`
	LastName			string				`json:"last_name"`
	Email       		string				`json:"email"`
	Password 			string				`json:"password"`
}

func GetAllAccount() ([]AccountDTO, error) {
	var accounts []AccountDTO

	result, err := Db.Query("select account_id, first_name, last_name, email from account")
	if err != nil {
		log.Printf("[ERROR] : %s", err)
		return nil, err
	}

	for result.Next() {
		var account AccountDTO
		err = result.Scan(&account.AccountId, &account.FirstName, &account.LastName, &account.Email)
		if err != nil {
			log.Printf("[ERROR] : %s", err)
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func GetAccountById(accountId *string) (*AccountDTO, error) {
	var result AccountDTO

	err := Db.QueryRow("select first_name, last_name, email from account where account_id = ?", accountId).Scan(&result.FirstName, &result.LastName, &result.Email)
	if err != nil {
		return nil, err
	}

	result.AccountId = *accountId
	return &result, nil
}

func InsertAccount(account *AccountInsert) (error) {
	userId := "B" + generateUserId(10)
	activationCode := generateActivationCode(6)
	_, err := Db.Exec("INSERT INTO account (account_id, first_name, last_name, email, password, password_reset_code, password_reset_expiry, account_create_time, account_update_time, is_active_account, activation_code) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", userId, account.FirstName, account.LastName, account.Email, account.Password, nil, 0, time.Now(), time.Now(), 0, activationCode)
	if err != nil {
		return err
	}

	err = util.SendEmail(account.Email, account.FirstName+ " " + account.LastName, "Account Activation Code", activationCode)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUsernameAccount(accountId, firstName, lastName *string) (error) {
	_, err := Db.Exec("update account set first_name = ?, last_name = ?, account_update_time = ? where account_id = ?", *firstName, *lastName, time.Now(), *accountId)
	if err != nil {
		return err
	}

	return nil
}

func DeleteAccount(accountId *string) (error) {
	_, err := Db.Exec("delete from account where account_id = ?", *accountId)
	if err != nil {
		return err
	}

	return nil
}

func generateUserId(n int) string {
	rand.Seed(time.Now().UnixNano())

	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func generateActivationCode(n int) string {
	rand.Seed(time.Now().UnixNano())

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}