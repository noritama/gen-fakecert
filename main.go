package main

import (
	"flag"
	"fmt"
	"os"
	 "./fakecert"
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

	fekecert.Generate(outPri, outCrt, country, organization)
}
