package utils

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"products/core/models"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CreateToken function creates and return an access
// token based in user id with 15 minute expiration.
func CreateToken(userID string) (*models.TokenDetails, error) {
	expTime := time.Now().Add(time.Minute * 15).Unix()

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = expTime

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))

	if err != nil {
		return nil, err
	}

	return &models.TokenDetails{
		AccessToken: tokenString,
	}, nil
}

// VerifyToken function validates token signature checking
// the sign method and secret used and return the token or
// an error.
func VerifyToken(request *http.Request) (*jwt.Token, error) {
	token := extractToken(request)

	if len(token) == 0 {
		return nil, errors.New("unauthorized")
	}

	verifyResult, err := jwt.Parse(token, func(result *jwt.Token) (interface{}, error) {
		if _, ok := result.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("wrong singning method: %v", result.Header["alg"])
		}

		return []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	return verifyResult, nil
}

func extractToken(request *http.Request) string {
	authorization := request.Header.Get("Authorization")

	strArr := strings.Split(authorization, " ")

	if len(strArr) > 1 {
		return strArr[1]
	}

	return ""
}

// TokenIsValid function verify the token passed in
// authorization request header and return only an
// error if the token is invalid or nil if it's valid.
func TokenIsValid(request *http.Request) (string, error) {
	token, err := VerifyToken(request)

	if err != nil {
		return "", err
	}

	data, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		userID := data["user_id"].(string)

		return userID, nil
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return "", err
	}

	return "", err
}
