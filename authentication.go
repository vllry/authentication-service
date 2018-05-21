package main

import (
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"io/ioutil"
	"crypto/rsa"
)

func loadPublicKey(keyPath string) (*rsa.PublicKey, error) {
	verifyBytes, err := ioutil.ReadFile("/home/vallery/Development/Go/src/jwtissuer/app.rsa.pub")
	if err != nil {
		return nil, err
	}
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return nil, err
	}

	return verifyKey, nil
}

func validateToken(verifyKey *rsa.PublicKey, token string) (bool, error) {
	//verifyBytes, err := ioutil.ReadFile("/home/vallery/Development/Go/src/jwtissuer/app.rsa.pub")
	//if err != nil {
	//	panic(err)
	//}
	//verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	//if err != nil {
	//	panic(err)
	//}

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}
	claims := MyCustomClaims{}
	processedToken, err := jwt.ParseWithClaims(
		token,
		&claims,
		func(token *jwt.Token) (interface{}, error) {
			fmt.Println("Callback")
			return verifyKey, nil
		},
	)

	fmt.Println(processedToken.Claims.(*MyCustomClaims).Foo)

	if err == nil && processedToken.Valid {
		fmt.Println("Your processedToken is valid.  I like your style.")
		return true, err
	} else {
		fmt.Println("This processedToken is terrible!  I cannot accept this.")
		return false, err
	}
}