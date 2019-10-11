package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var certiPrt string
var certoPrt string
var extension string

func main() {

	// Possible Arguments (Arg-Name, Arg-Default, Help-Text)
	certiPrt := flag.String("in", "", "<certfile> in")
	certoPrt := flag.String("out", "", "[certfile] out")
	keyiPrt := flag.String("ki", "", "<certfile> in")
	keyoPrt := flag.String("ko", "", "<certfile> out")
	cvrtfPrt := flag.String("f", "", "format: {x509|pkcs7|pkcs12|der}")

	// Parsing the Flags
	flag.Parse()

	// Check Args
	if *certiPrt == "" || *cvrtfPrt == "" {
		flag.PrintDefaults()
		os.Exit(1)
		// If outfile is not set, set it according to format
	} else if *certoPrt == "" {
		certoPrt := strings.Split(*certiPrt, ".")
		fileslice := certoPrt[:len(certoPrt)-1]

		switch *cvrtfPrt {
		case "x509":
			extension = "crt"
			ConvPem(*certiPrt)
		case "pkcs7":
			extension = "pk"
			ConvP7(*certiPrt)
		case "pkcs12":
			extension = "p12"
			ConvP12(*certiPrt)
		case "der":
			extension = "der"
			ConvDer(*certiPrt)
		default:
			fmt.Printf("Unsupported format: %s", *cvrtfPrt)
			flag.PrintDefaults()
			os.Exit(1)
		}

		fileslice = append(fileslice, extension)
		outfile := strings.Join(fileslice, ".")

		fmt.Printf("OUT: %s\n", outfile)
	}

	if *keyiPrt != "" || *keyoPrt != "" {
		fmt.Printf("Keys - Out: %s, In: %s", *keyoPrt, *keyiPrt)
	}

	fmt.Printf("CERT in: %s, FORMAT: %s\n", *certiPrt, *cvrtfPrt)
}

//ConvPem - Does the conversion with x509 certs
func ConvPem(certiPrt string) {
	fmt.Printf("inside x509 func\n")
}

//ConvP7 - Does the conversion with pkcs7 certs
func ConvP7(certiPrt string) {
	fmt.Printf("inside pkcs7 func\n")
}

//ConvP12 - Does the conversion with pkcs12 certs
func ConvP12(certiPrt string) {
	fmt.Printf("inside pkcs12 func\n")
}

//ConvDer - Does the conversion with der certs
func ConvDer(certiPrt string) {
	fmt.Printf("inside der func\n")
}
