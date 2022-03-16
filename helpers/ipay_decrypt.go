package helpers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"os"
)

func RSADecrypt(ciphertext string) (string, error) {
	//Open file
	mydir, _ := os.Getwd()
	file, err := os.Open(mydir + "/res/tiki_private_key.pem")
	if err != nil {
		return "", err
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
		privatePkcs8Key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return "", err
		}
		privateKey = privatePkcs8Key.(*rsa.PrivateKey)
	} else {
		privateKey = key
	}
	//Encrypt plaintext
	// md5Str := md5.New()
	//hashed := sha1.Sum(ciphertext)
	msg, err := base64.StdEncoding.DecodeString(ciphertext)
	data, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, msg)
	if err != nil {
		return "", err
	}

	//Return ciphertext
	plainText := string(data)
	return plainText, nil
}
