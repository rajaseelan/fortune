// microservice that spits out random fortune cookies
// fortune material acquired from fortune program in ubuntu 20.04
package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/rajaseelan/fortune/pkg"
)

func main() {

	fortuneFile := parseInput()

	port := getAppPort()

	// Load up linux Fortunes
	fortune := pkg.LoadFortunes(*fortuneFile)
	cookie := NewCookie(fortune)

	mux := http.NewServeMux()
	mux.HandleFunc("/", cookie.home)

	err := http.ListenAndServe(":"+port, mux)
	log.Fatal(err)
}

func parseInput() *string {
	fortuneFile := flag.String("fortune", "", "path to fortune file")
	flag.Parse()

	return fortuneFile
}

func getAppPort() string {
	port, portExist := os.LookupEnv("FORTUNE_PORT")
	if !portExist {
		port = "3000"
	}

	return port
}
