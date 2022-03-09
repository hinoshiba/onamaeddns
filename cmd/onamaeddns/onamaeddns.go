package main

import (
	"os"
	"os/user"
	"fmt"
	"flag"
	"time"
	"bufio"
	"strings"
)

import (
	"github.com/hinoshiba/onamaeddns"
)

var (
	timeout = 10 * time.Second

	Pass      string
	User      string
	IpAddress string
	Host      string
	Domain    string
)

func ddns() error {
	cl, err := onamaeddns.Dial(onamaeddns.OfficialAddress, User, Pass, timeout)
	if err != nil {
		return err
	}
	defer cl.Close()

	if err := cl.UpdateIPv4(Host, Domain, IpAddress); err != nil {
		return err
	}
	return nil
}

func die(s string, msg ...interface{}) {
	fmt.Fprintf(os.Stderr, s + "\n" , msg...)
	os.Exit(1)
}

func loadConf(path string) (string, string, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		return "", "", err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	if !sc.Scan() {
		if err := sc.Err(); err != nil {
			return "", "", err
		}
		return "", "", fmt.Errorf("cant read line")
	}
	row := sc.Text()

	creds := strings.SplitN(row, ":", 2)
	if len(creds) < 2 {
		return "", "", fmt.Errorf("cant load credential")
	}
	return creds[0], creds[1], nil
}

func init() {
	var c_path string
	flag.StringVar(&c_path, "c", "", "path of credential")
	flag.Parse()

	if flag.NArg() < 3 {
		die("Usage : onamaeddns <host> <domain> <ip address>")
	}

	host := flag.Arg(0)
	if host == "" {
		die("empty hostname")
	}
	domain := flag.Arg(1)
	if domain == "" {
		die("empty domain")
	}
	ipaddr := flag.Arg(2)
	if ipaddr == "" {
		die("empty ip address")
	}

	if c_path == "" {
		usr, err := user.Current()
		if err != nil {
			die("%s", err)
		}
		c_path = usr.HomeDir + "/.onamaeddns"
	}
	user, pass, err := loadConf(c_path)
	if err != nil {
		die("%s", err)
	}

	User = user
	Pass = pass
	Domain = domain
	Host = host
	IpAddress = ipaddr
}

func main() {
	if err := ddns(); err != nil {
		die("%s", err)
	}
	fmt.Printf("updated.\n")
}
