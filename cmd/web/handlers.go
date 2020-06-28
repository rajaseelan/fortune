package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/rajaseelan/fortune/pkg"
)

// Cookie spits out a random fortune + host details for Get Requests
type Cookie struct {
	Fortune *pkg.Fortune
}

func (c Cookie) home(w http.ResponseWriter, r *http.Request) {

	jsonBytes := c.encodeAsJSON()
	fmt.Println(string(*jsonBytes))
}

func (c Cookie) encodeAsJSON() *[]byte {
	cookieString := c.randomQuote()
	timeString := c.getCurrentTime()
	hostName := c.getHostName()

	answer := pkg.Answer{
		Time:     *timeString,
		Fortune:  *cookieString,
		Hostname: *hostName,
		Set:      *c.Fortune.SetName,
	}

	jsonBytes, err := json.MarshalIndent(&answer, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return &jsonBytes
}

func (c Cookie) getHostName() *string {
	hostName, err := os.Hostname()
	if err != nil {
		hostName = err.Error()
	}
	return &hostName
}

func (c Cookie) getCurrentTime() *string {
	t := time.Now().String()
	return &t
}

func (c Cookie) randomQuote() *string {

	randomQuoteIndex := rand.Intn(len(c.Fortune.Cookies))
	return c.Fortune.Cookies[randomQuoteIndex]
}

// NewCookie returns a Cookie HTTP Handler
func NewCookie(f *pkg.Fortune) Cookie {
	return Cookie{
		Fortune: f,
	}
}
