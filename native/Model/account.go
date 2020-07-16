package Model

import (
	"carakan-apps/util"
	"errors"
	"fmt"
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
	AccountCreateTime 	time.Time			`json:"-"`
	AccountUpdateTime	time.Time			`json:"-"`
	IsActiveAccount		bool				`json:"-"`
	ActivationCode		string				`json:"-"`
}

func GetAllAccount() ([]Account, error) {
	var accounts []Account

	result, err := Db.Query("select account_id, first_name, last_name, email from account")
	if err != nil {
		log.Printf("[ERROR] : %s", err)
		return nil, err
	}

	for result.Next() {
		var account Account
		err = result.Scan(&account.AccountId, &account.FirstName, &account.LastName, &account.Email)
		if err != nil {
			log.Printf("[ERROR] : %s", err)
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func GetAccountById(accountId string) (*Account, error) {
	var result Account

	err := Db.QueryRow("select first_name, last_name, email, is_active_account from account where account_id = ?", accountId).Scan(&result.FirstName, &result.LastName, &result.Email, &result.IsActiveAccount)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("cannot get data from database")
	}

	result.AccountId = accountId
	return &result, nil
}

func GetAccountByEmail(account *Account) (*Account, error) {
	var result Account
	err := Db.QueryRow("select (first_name, last_name, email) from account where account_id = ?", account.AccountId).Scan(&result.FirstName, &result.LastName, &result.Email)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("cannot get data from database")
	}

	result.AccountId = account.AccountId
	return &result, nil
}

func InsertAccount(account *Account) (error) {
	userId := "B" + generateUserId(10)
	activationCode := generateActivationCode(6)
	_, err := Db.Exec("insert into account (account_id, first_name, last_name, email, password, account_create_time, account_update_time, is_active_account, activation_code) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", userId, account.FirstName, account.LastName, account.Email, account.Password , time.Now(), time.Now(), 0, activationCode)

	if err != nil {
		log.Printf("[ERROR] : %s", err)
		return errors.New("cannot insert data")
	}

	emailTo := []string{account.Email}
	message := "your activation code " + activationCode
	username := account.FirstName + " " + account.LastName

	err = util.SendMail(emailTo, username,"activation code", message)

	return nil
}

func UpdateUsernameAccount(accountId, firstName, lastName string) (error) {
	_, err := Db.Exec("update account set first_name = ?, last_name = ?, account_update_time = ? where account_id = ?", firstName, lastName, time.Now(), accountId)
	if err != nil {
		log.Printf("[ERROR] %s", err)
		return err
	}

	return nil
}

func DeleteAccount(accountId string) (error) {
	_, err := Db.Exec("delete from account where account_id = ?", accountId)
	if err != nil {
		log.Printf("[ERROR] %s", err)
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

func test() {
/*	account, _ := GetAllAccount()
	fmt.Println(account)*/

	/*	account, err := GetAccountById("BNIJTLUWBQW")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(*account)*/

/*	err := util.SendMail([]string{"ahmad.asrori08@gmail.com"}, "ahmad asrori", "hello", "1234567");
	if err != nil {
		fmt.Println(err)
	}*/

/*		user := Account{
		FirstName: "ahmad",
		LastName: "asrori",
		Email: "a@gmail.com",
		Password: "123",
		AccountCreateTime: time.Now(),
		AccountUpdateTime: time.Now(),
		ActivationCode: "123456",
		IsActiveAccount: false,
	}

	err := InsertAccount(&user)
	if err != nil {
		log.Println(err)
	}*/

	/*updateUser := Account{
		AccountId: "BWIZBMSQMWG",
		FirstName: "kevin",
		LastName: "sanjaya",
	}

	err := UpdateUsernameAccount(&updateUser)
	if err != nil {
		log.Println(err)
	}*/

	/*err := DeleteAccount("BWIZBMSQMWG")
	if err != nil {
		log.Println(err)
	}*/
}