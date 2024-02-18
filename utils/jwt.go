package utils

import (
	"errors"

	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

func GenerateToken(email string, userId *uint) (string, error) {
	//SigningMethodHS256 typ HMAC sprawdzany poniżej

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(), /*ten token jest wazny 2 h*/
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (*uint, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		//Sprawdzamy czy token jest takiego typu
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signin method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, errors.New("Could not parse token")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return nil, errors.New(("Invalid token"))
	}

	//Access to MapClaims ( above ) field
	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("Invalid token claims")
	}

	//te nawiasy na końcu to type check
	//email := claims["email"].(string)

	var test1 float64 = claims["userId"].(float64)
	userId := uint(test1)
	return &userId, nil
}
