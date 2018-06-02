package main

import "testing"

func TestLoadPublicKey(t *testing.T) {
	pubkeyPath := "test/nokey.pem"
	pubkey, err := LoadPublicKey(pubkeyPath)
	if pubkey != nil {
		t.Errorf("loadPublicKey() returned a result - it should be returning nil due to a missing file")
	}
	if err == nil {
		t.Errorf("loadPublicKey() returned no error - it should be failing due to a missing file")
	}

	pubkeyPath = "test/cert1.pem"
	pubkey, err = LoadPublicKey(pubkeyPath)
	if pubkey == nil {
		t.Errorf("loadPublicKey() returned no pubkey")
	}
	if err != nil {
		t.Errorf("loadPublicKey() returned error - %s", err)
	}
}

//func TestValidateToken(t *testing.T) {
//	pubkeyPath := "test/cert1.pem"
//	pubkey, err := LoadPublicKey(pubkeyPath)
//
//	claims, err := ValidateToken(pubkey, token)
//}