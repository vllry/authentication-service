package main

import (
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
)

type CustomClaims struct {
	Foo string `json:"foo"`
	jwt.StandardClaims
}

func LoadPublicKey(pubkeyPath string) (*rsa.PublicKey, error) {
	verifyBytes, err := ioutil.ReadFile(pubkeyPath)
	if err != nil {
		return nil, err
	}
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return nil, err
	}

	return verifyKey, nil
}

func ValidateToken(verifyKey *rsa.PublicKey, token string) (CustomClaims, error) {
	claims := CustomClaims{}
	processedToken, err := jwt.ParseWithClaims(
		token,
		&claims,
		func(token *jwt.Token) (interface{}, error) {
			fmt.Println("Callback")
			return verifyKey, nil
		},
	)

	fmt.Println(processedToken.Claims.(*CustomClaims).Foo)

	if err == nil && processedToken.Valid {
		fmt.Println("Your processedToken is valid.  I like your style.")
		return nil, err
	} else {
		fmt.Println("This processedToken is terrible!  I cannot accept this.")
		return nil, err
	}

	return processedToken.Claims.(CustomClaims), nil
}
