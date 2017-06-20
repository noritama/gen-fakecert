package fekecert

import (
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"crypto/rsa"
	"math/big"
	"time"
	"encoding/pem"
	"io/ioutil"
	"fmt"
)

func Generate(key *string, crt *string, country *string, organization *string) {
	template := &x509.Certificate{
		IsCA: true,
		BasicConstraintsValid: true,
		SubjectKeyId:          []byte{1, 2, 3},
		SerialNumber:          big.NewInt(1234),
		Subject: pkix.Name{
			Country:      []string{*country},
			Organization: []string{*organization},
		},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().AddDate(10, 0, 0),
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	cert, err := x509.CreateCertificate(rand.Reader, template, template, &privateKey.PublicKey, privateKey)
	if err != nil {
		panic(err)
	}
	priStr := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	crtStr := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert,
	})

	ioutil.WriteFile(*key, priStr, 0600)
	ioutil.WriteFile(*crt, crtStr, 0644)
	fmt.Println("Output private-key file:", key)
	fmt.Println("Output cert file:", crt)
}