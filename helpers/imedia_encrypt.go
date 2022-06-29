package helpers

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func IMediaRSAEncrypt(plainText []byte) ([]byte, error) {
	//Open file
	mydir, _ := os.Getwd()
	// key, err := ioutil.ReadFile(mydir + "/res/private.pem")
	file, err := os.Open(mydir + "/res/imedia_tiki_private.pem")
	if err != nil {
		return nil, err
		// panic(err)
	}
	defer file.Close()
	//Read the contents of the file
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)

	//PEM decoding
	block, _ := pem.Decode(buf)
	//X509 decoding
	var privateKey *rsa.PrivateKey
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("TEST")
		privatePkcs8Key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		privateKey = privatePkcs8Key.(*rsa.PrivateKey)
	} else {
		privateKey = key
	}
	//Encrypt plaintext
	// md5Str := md5.New()
	hashed := sha256.Sum256(plainText)
	cipherText, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return nil, err
	}
	//Return ciphertext
	return cipherText, nil
}
