package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"github.com/gin-gonic/gin"
)

func getEngine() *gin.Engine {
	r := gin.Default()
	return r
}

func rsaEncrypt(pubkey string, pin string, uuid string) (string, error) {
	block, err := base64.StdEncoding.DecodeString(pubkey)
	if err != nil {
		return "", err
	}
	pub, err := x509.ParsePKIXPublicKey(block)
	if err != nil {
		return "", err
	}
	rsaPub, _ := pub.(*rsa.PublicKey)

	rsakey, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPub, []byte(uuid+pin))
	if err != nil {
		return "", err
	}
	encodedKey := base64.StdEncoding.EncodeToString(rsakey)
	return encodedKey, nil
}
