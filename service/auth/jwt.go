package auth

import (
	"net/http"
	"strconv"
	"time"

	"github.com/AbhishekPSingh07/ecom_go/config"
	"github.com/AbhishekPSingh07/ecom_go/types"
	"github.com/golang-jwt/jwt"
)

func CreateJWT(secret []byte, userID int) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationiInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "",err
	}
	return tokenString, nil
}

func WithJWTAuth(handleFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the token from the user request
		// validate the JWT
		//set context "UserID"
	}
}