package Middleware

import (
	"carakan-apps/configuration"
	"carakan-apps/controller"
	"carakan-apps/util"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		//ambil key Authorization dari header
		header := request.Header.Get("Authorization")

		// case : jika didalam request tidak terdapat bearer token
		if header == "" && request.URL.Path != "/auth" {
			http.Error(response, "Forbidden", http.StatusForbidden)
		}

		// case : jika user melakukan request token
		if request.URL.Path == "/auth" {
			next.ServeHTTP(response, request)
		}

		// case : jika request terdapat bearer token
		if header != "" {

			//ambil bearer token
			getAuthHeader := request.Header.Get("Authorization")
			myToken := strings.Replace(getAuthHeader, "Bearer ", "", 1)

			key := []byte(configuration.JWT_SECRET_KEY)

			// ubah token ke claim object
			token, err := jwt.ParseWithClaims(myToken, &controller.MyClaim{}, func(token *jwt.Token) (interface{}, error) {

				// dan validasi algoritma yang dipakai untuk memparsing token sudah sesuai
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				return key, nil
			})

			// case : gagal parsing token ke claim
			if err != nil {
				fmt.Println(err)
				http.Error(response, "Bad Gateway", http.StatusBadGateway)
			}

			//validasi token apalah sudah valid
			if claim, ok := token.Claims.(*controller.MyClaim); ok && token.Valid {

				// periksa apakah user boleh mengakses resource dengan HTTP method tersebut
				authorization, err := util.IsAuthorize(claim.User.UserName, request.URL.Path, request.Method)

				// case : gagal memeriksa otorisasi
				if err != nil {
					http.Error(response, "Bad Gateway", http.StatusBadGateway)
				}

				// case : user memiliki otorisasi terhadap resource
				if authorization {
					next.ServeHTTP(response, request)
				}

				// case : user tidak memiliki otorisasi terhadap resource
				if !authorization {
					http.Error(response, "Unauthorized", http.StatusUnauthorized)
				}

			} else {

				// case : token tidak valid
				http.Error(response, "Forbidden", http.StatusForbidden)
			}
		}
	})
}