package helpers

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func loadPrivateKey() (*rsa.PrivateKey, error) {
	privateKeyBytes, err := os.ReadFile("./private")
	if err != nil {
		return nil, fmt.Errorf("error reading private key: %s", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing private key: %s", err)
	}

	return privateKey, nil
}

func loadPublicKey() (*rsa.PublicKey, error) {
	publicKeyBytes, err := os.ReadFile("./public")
	if err != nil {
		return nil, fmt.Errorf("error reading public key: %s", err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing public key: %s", err)
	}

	return publicKey, nil
}

func GenerateAccessToken(userID string) (string, error) {
	privateKey, _ := loadPrivateKey()
	exp, _ := strconv.Atoi(os.Getenv("AUTH_EXP"))
	expiredTime := time.Duration(exp) * time.Minute

	accessTokenExp := time.Now().Add(expiredTime).Unix()

	accessClaim := jwt.MapClaims{
		"sub": userID,
		"exp": accessTokenExp,
	}

	accessJwt := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), accessClaim)

	return accessJwt.SignedString(privateKey)
}

func VerifyToken(tokenString string) (string, error) {
	publicKey, _ := loadPublicKey()

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("Unauthorized")
		}
		return publicKey, nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token")
	}

	expValue := claims["exp"]
	if reflect.TypeOf(expValue).Kind() == reflect.Float64 {
		exp := int64(expValue.(float64))
		currentUnixTime := time.Now().Unix()
		if currentUnixTime > exp {
			return "", errors.New("token has expired")
		}
	} else {
		return "", errors.New("invalid token")
	}

	userID := claims["sub"].(string)

	return userID, nil
}
