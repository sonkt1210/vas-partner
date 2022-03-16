package helpers

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func RSAEncrypt(plainText []byte) ([]byte, error) {
	//Open file
	mydir, _ := os.Getwd()
	// key, err := ioutil.ReadFile(mydir + "/res/private.pem")
	file, err := os.Open(mydir + "/res/tiki_private_key.pem")
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
		fmt.Println("TÉT")
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
	hashed := sha1.Sum(plainText)
	cipherText, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA1, hashed[:])
	if err != nil {
		return nil, err
	}
	//Return ciphertext
	return cipherText, nil
}
