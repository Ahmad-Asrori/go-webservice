package controller

import (
	"github.com/labstack/echo"
	"krama/model"
	"net/http"
)

func GetAllAccount(c echo.Context) error {
	accounts, err := model.GetAllAccount()
	if err != nil {
		return c.JSON(http.StatusBadGateway, "error")
	}

	return c.JSON(http.StatusOK, accounts)
}

func GetAccountById(c echo.Context) error {
	var accountId model.AccountId
	c.Bind(&accountId)

	account, err := model.GetAccountById(&accountId.AccountId)
	if err != nil {
		return c.JSON(http.StatusBadGateway, "error")
	}

	return c.JSON(http.StatusOK, account)
}

func InsertAccount(c echo.Context) error {
	var account model.AccountInsert
	c.Bind(&account)

	err := model.InsertAccount(&account)
	if err != nil {
		return c.JSON(http.StatusBadGateway, err)
	}

	resp := map[string]string{"message": "OK"}
	return c.JSON(http.StatusOK, resp)
}

func UpdateAccountById(c echo.Context) error {
	var account model.AccountDTO
	c.Bind(&account)

	err := model.UpdateUsernameAccount(&account.AccountId, &account.FirstName, &account.LastName)
	if err != nil {
		return c.JSON(http.StatusBadGateway, "error")
	}

	resp := map[string]string{"message": "OK"}
	return c.JSON(http.StatusOK, resp)
}

func DeleteAccountById(c echo.Context) error {
	var account model.AccountId
	c.Bind(&account)

	err := model.DeleteAccount(&account.AccountId)
	if err != nil {
		return c.JSON(http.StatusBadGateway, "error")
	}

	resp := map[string]string{"message": "OK"}
	return c.JSON(http.StatusOK, resp)
}