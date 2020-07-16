package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"krama/configuration"
	"krama/model"
	"net/http"
	"time"
)

type JWTAccountRequest struct {
	Username 	string		`json:"username"`
	Password	string		`json:"password"`
}

type MyClaim struct {
	model.UsersDTO
	jwt.StandardClaims
}

type JWTTokenResponse struct {
	Token 		string		`json:"token"`
}

func Auth(c echo.Context) error {
	// siapkan struct untuk menampung data
	var user JWTAccountRequest

	// binding data JSON ke struct
	err := c.Bind(&user)
	if err != nil {
		return err
	}

	// cek ke database apakah user ada
	data, err := model.GetAuthUser(&user.Username, &user.Password)
	if err != nil {
		return err
	}

	// build claim
	myCLaim := MyClaim{
		*data, jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 0, 1).Unix(),
		},
	}

	// beritahu apabila token akan digenerate menggunakan algoritma HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myCLaim)

	// generate token disertai dengan JWT Secret
	tokenResult, err := token.SignedString([]byte(configuration.JWT_SECRET))
	if err != nil {
		return err
	}

	resp := JWTTokenResponse{
		Token: tokenResult,
	}

	return c.JSON(200, resp)
}

func Verify(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*MyClaim)
	name := claims.UserName

	return c.String(http.StatusOK, "Welcome "+name+"!")
}