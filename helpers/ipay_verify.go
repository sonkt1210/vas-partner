package helpers

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"os"
)

func RSAVerify(msg string, sig string) error {
	//Open file
	mydir, _ := os.Getwd()
	// key, err := ioutil.ReadFile(mydir + "/res/private.pem")
	file, err := os.Open(mydir + "/res/ipay_pub.pem")
	if err != nil {
		return err
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
	var publicKey *rsa.PublicKey
	publicPkixKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	publicKey = publicPkixKey.(*rsa.PublicKey)

	message := []byte(msg)
	h := sha1.New()
	h.Write(message)
	digest := h.Sum(nil)
	ds, _ := base64.StdEncoding.DecodeString(sig)

	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA1, digest, ds)
	if err != nil {
		return err
	}
	//Return ciphertext
	return nil
}
