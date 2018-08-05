package main

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"testing"
)

func loadPrivateKey() (*rsa.PrivateKey, error) {
	secretKeyBytes, err := ioutil.ReadFile("test/key1.pem")
	if err != nil {
		return nil, err
	}
	secretKey, err := jwt.ParseRSAPrivateKeyFromPEM(secretKeyBytes)
	if err != nil {
		return nil, err
	}

	return secretKey, nil
}

func generateToken(userId string) (string, error) {
	secretKey, err := loadPrivateKey()
	if err != nil {
		return "", err
	}

	secretToken := jwt.NewWithClaims(
		jwt.SigningMethodRS256,
		jwt.MapClaims{
			"userId": userId,
		})

	secretTokenString, err := secretToken.SignedString(secretKey)

	return secretTokenString, err
}

func TestLoadPublicKey(t *testing.T) {
	pubkeyPath := "test/nokey.pem"
	pubkey, err := loadPublicKey(pubkeyPath)
	if pubkey != nil {
		t.Errorf("loadPublicKey() returned a result - it should be returning nil due to a missing file")
	}
	if err == nil {
		t.Errorf("loadPublicKey() returned no error - it should be failing due to a missing file")
	}

	pubkeyPath = "test/cert1.pem"
	pubkey, err = loadPublicKey(pubkeyPath)
	if pubkey == nil {
		t.Errorf("loadPublicKey() returned no pubkey")
	}
	if err != nil {
		t.Errorf("loadPublicKey() returned error - %s", err)
	}
}

func TestValidateToken(t *testing.T) {
	testUser := "testuser"

	token, err := generateToken(testUser)
	if err != nil {
		t.Errorf("Generation error: %s", err)
	}

	pubkeyPath := "test/cert1.pem"
	pubkey, _ := loadPublicKey(pubkeyPath)
	claims, err := validateToken(pubkey, token)
	if err != nil {
		t.Errorf("Validation error: %s", err)
	}

	if claims.UserId != testUser {
		t.Errorf("Wrong userId - got %s", claims.UserId)
	}

	badToken := "evil"
	_, err = validateToken(pubkey, badToken)
	if err == nil {
		t.Errorf("Expected invalid token to fail.")
	}
}
