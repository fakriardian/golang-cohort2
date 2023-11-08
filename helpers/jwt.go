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

func loadPrivateKey(privateKeyBytes []byte) (*rsa.PrivateKey, error) {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing private key: %s", err)
	}

	return privateKey, nil
}

func loadPublicKey(publicKeyBytes []byte) (*rsa.PublicKey, error) {
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing public key: %s", err)
	}

	return publicKey, nil
}

func GenerateAccessToken(userID string, privateKey []byte) (string, error) {
	privateKeyPem, _ := loadPrivateKey(privateKey)
	exp, _ := strconv.Atoi(os.Getenv("AUTH_EXP"))
	expiredTime := time.Duration(exp) * time.Minute

	accessTokenExp := time.Now().Add(expiredTime).Unix()

	accessClaim := jwt.MapClaims{
		"sub": userID,
		"exp": accessTokenExp,
	}

	accessJwt := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), accessClaim)

	return accessJwt.SignedString(privateKeyPem)
}

func VerifyToken(tokenString string, publicKey []byte) (string, error) {
	publicKeyPem, _ := loadPublicKey(publicKey)

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("Unauthorized")
		}
		return publicKeyPem, nil
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
