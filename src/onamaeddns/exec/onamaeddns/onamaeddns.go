package main

import (
	"log"
	"time"
)

import (
	"onamaeddns"
)

var (
	timeout = 10 * time.Minute
)

const (
	ONAMAE_SV = "ddnsclient.onamae.com:65010"
)

func main() {
	cl, err := onamaeddns.Dial(ONAMAE_SV, "username", "password", timeout)
	if err != nil {
		log.Println(err)
		return
	}
	defer cl.Close()

	if err := cl.UpdateIPv4("mytest", "example.com", "127.0.0.1"); err != nil {
		log.Println(err)
		return
	}
	log.Println("updated")
}
