package main

import (
	"regexp"
	"log"
	"time"
	"crypto/tls"
	"github.com/google/goexpect"
)

var (
	timeout = 10 * time.Minute
)

const (
	TEST_SERVER = "www.i.hinoshiba.com:443"
)

func main() {
	conn, err := tls.Dial("tcp", TEST_SERVER, nil)
	if err != nil {
		log.Println(err)
		return
	}

	resCh := make(chan error)

	exp, _, err := expect.SpawnGeneric(&expect.GenOptions{
		In:  conn,
		Out: conn,
		Wait: func() error {
			return <-resCh
		},
		Close: func() error {
			close(resCh)
			return conn.Close()
		},
		Check: func() bool { return true },
	}, timeout, expect.Verbose(true))
	if err != nil {
		log.Println(err)
		return
	}
	defer exp.Close()

	promptRE := regexp.MustCompile("HTTP")
	exp.Send("" + "\n")
	exp.Send("" + "\n")
	exp.Send("." + "\n")
	log.Println(exp.Expect(promptRE, timeout))
}
