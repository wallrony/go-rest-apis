package tokenutils

import (
	models "auth_token/models"

	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

// CreateToken Function create an authorization token with
// userID.
func CreateToken(userID uint64) (*models.TokenDetails, error) {
	var err error

	td := &models.TokenDetails{}

	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUUID = uuid.NewV4().String()

	// Creating the Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["user_id"] = userID
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	td.AccessToken, err = accessToken.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))

	if err != nil {
		return nil, err
	}

	return td, nil
}

// ExtractToken function extract authorization token from request Header.
func ExtractToken(request *http.Request) string {
	tokenString := request.Header.Get("Authorization")

	strArr := strings.Split(tokenString, " ")

	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}

// VerifyToken function validates token signature checking
// your method and secret and return the token.
func VerifyToken(request *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(request)

	if len(tokenString) == 0 {
		return nil, errors.New("need authorization token")
	}

	token, err := jwt.Parse(tokenString, func(tokenResult *jwt.Token) (interface{}, error) {
		if _, ok := tokenResult.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", tokenResult.Header["alg"])
		}

		return []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

// TokenIsValid function verify if token passed by user
// is valid, returning nil if is and an error if not.
func TokenIsValid(request *http.Request) error {
	token, err := VerifyToken(request)

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}

	return nil
}
