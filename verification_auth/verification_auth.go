package verification_auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func ExtractTokenMetadata(r *http.Request, accessSecret *string, parse func(jwt.MapClaims) interface{}) (interface{}, error) {
	token, err := VerifyToken(r, accessSecret)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		if accessSecret != nil && !token.Valid {
			return nil, errors.New("access token is invalid")
		}
		return parse(claims), nil
	}
	return nil, errors.New("claims is invalid")
}

func VerifyToken(r *http.Request, accessSecret *string) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secret := []byte("")
		if accessSecret != nil {
			secret = []byte(*accessSecret)
		}
		return secret, nil
	})
	if accessSecret != nil && err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

type User interface {
	Role() string
}
