package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"time"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	///
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `
Usage of %s:
   %s [OPTIONS] ARGS...
Options`, os.Args[0], os.Args[0])

		flag.PrintDefaults()
	}

	outPri := flag.String("key", cwd+"/server.key", "output private-key file path")
	outCrt := flag.String("crt", cwd+"/server.crt", "output cert file path")
	country := flag.String("country", "Japan", "your country")
	organization := flag.String("organization", "Noritama", "your organization")

	flag.Parse()

	///
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

	ioutil.WriteFile(*outPri, priStr, 0600)
	ioutil.WriteFile(*outCrt, crtStr, 0644)
	fmt.Println("Output private-key file:", outPri)
	fmt.Println("Output cert file:", outCrt)
}
