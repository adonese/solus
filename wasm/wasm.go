package main

import (
	"syscall/js"
	// "github.com/google/uuid"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
)

var c chan bool
var document js.Value

func init() {
	document = js.Global().Get("document")
	c = make(chan bool)
}

// function definition
func add(this js.Value, i []js.Value) interface{} {
	return js.ValueOf(i[0].Int() + i[1].Int())
}

// exposing to JS

func main() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("encrypt", js.FuncOf(encrypt))
	<-c
}

// encrypt list of items are as follows:
// pubkey, pin, uuid
func encrypt(this js.Value, i []js.Value) interface{} {

	pubkey, pin, uuid := js.Value(i[0]).String(), js.Value(i[1]).String(), js.Value(i[2]).String()
	pinBlock, err := rsaEncrypt(pubkey, pin, uuid)
	if err != nil {
		panic(err)
	}
	return pinBlock
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
