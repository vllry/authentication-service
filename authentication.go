package main

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
)

type CustomClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

func loadPublicKey(pubkeyPath string) (*rsa.PublicKey, error) {
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

func validateToken(verifyKey *rsa.PublicKey, token string) (*CustomClaims, error) {
	claims := CustomClaims{}
	processedToken, err := jwt.ParseWithClaims(
		token,
		&claims,
		func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		},
	)

	if err != nil {
		fmt.Println("Error during token validation", err)
		return nil, err
	} else if !processedToken.Valid {
		fmt.Println("token not valid")
		return nil, errors.New("processed token not valid")
	}

	processedClaims := processedToken.Claims.(*CustomClaims)
	return processedClaims, nil
}
