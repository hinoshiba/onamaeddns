// onamaeddns is library of ddns for 'onamae.com' service.
// 'onamae.com' service have a official ddns client, But for Windows only. // https://help.onamae.com/answer/7920

package onamaeddns

import (
	"fmt"
	"sync"
	"time"
	"crypto/tls"
	"regexp"
	"strings"
	"context"
)

import (
	"github.com/google/goexpect"
)

const (
	debug bool = false

	statusSucess string = "000"
	stautsFailed string = "001"

	OfficialAddress string = "ddnsclient.onamae.com:65010"
)

var (
	regexp_status *regexp.Regexp = regexp.MustCompile("\\d{3}")

	ErrNotConnect error = fmt.Errorf("tls not connected")
)

// Client represents a DDNS session on TLS connection.
type Client struct {
	exp    *expect.GExpect
	ctx    context.Context
	mtx    *sync.Mutex
}

// Dial connects to the given address(<host>:<port>) use 'user' and 'pass'.
func Dial(address string, user string, pass string, timeout time.Duration) (*Client, error) {
	return dial(context.Background(), address, user, pass, timeout)
}

// DialWithContext connects to the given address(<host>:<port>) use 'user' and 'pass'.
func DialWithContext(ctx context.Context, address string, user string, pass string, timeout time.Duration) (*Client, error) {
	return dial(ctx, address, user, pass, timeout)
}

func dial(ctx context.Context, address string, user string, pass string, timeout time.Duration) (*Client, error) {
	exp, err := createTlsExpect(ctx, address, timeout)
	if err != nil {
		return nil, err
	}

	self := &Client{exp:exp, mtx:new(sync.Mutex), ctx:ctx}
	if err := self.login(user, pass); err != nil {
		defer self.exp.Close()
		return nil, err
	}
	if self.isCanceledContext() {
		return nil, fmt.Errorf("closed context.")
	}
	return self, nil
}

// Close closes the session and connection.
func (self *Client) Close() error {
	self.lock()
	defer self.unlock()

	if err := self.logout(); err != nil {
		return err
	}
	return self.exp.Close()
}

func (self *Client) isCanceledContext() bool {
	select {
	case <-self.ctx.Done():
		return true
	default:
	}
	return false
}

// UpdateIPv4 is update value of A record on '<host>.<domain>' to '<ip>'.
func (self *Client) UpdateIPv4(host string, domain string, ip string) error {
	self.lock()
	defer self.unlock()

	if self.isCanceledContext() {
		return fmt.Errorf("closed context.")
	}
	return self.update(host, domain, ip)
}

// Expect extension function.
// Wait on the session if want use arbitrary command.
func (self *Client) Expect(rxp *regexp.Regexp, timeout time.Duration) (string, []string, error) {
	self.lock()
	defer self.unlock()

	if self.isCanceledContext() {
		return "", nil, fmt.Errorf("closed context.")
	}
	return self.expect(rxp, timeout)
}

// Send extension function.
// Sent to the session if want use arbitrary command.
func (self *Client) Send(s string, val ...interface{}) error {
	self.lock()
	defer self.unlock()

	if self.isCanceledContext() {
		return fmt.Errorf("closed context.")
	}
	msg := fmt.Sprintf(s, val...)
	return self.send(msg)
}

func (self *Client) login(user string, pass string) error {
	if err := self.send("LOGIN\n"); err != nil {
		return err
	}
	if err := self.send("USERID:" + user + "\n"); err != nil {
		return err
	}
	if err := self.send("PASSWORD:" + pass + "\n"); err != nil {
		return err
	}
	if err := self.send(".\n"); err != nil {
		return err
	}

	st, err := self.getStatus()
	if err != nil {
		return err
	}
	if !st.ok {
		return fmt.Errorf("ResponseError : %s", st.msg)
	}
	return nil
}

func (self *Client) logout() error {
	if err := self.send("LOGOUT\n"); err != nil {
		return err
	}
	if err := self.send(".\n"); err != nil {
		return err
	}

	st, err := self.getStatus()
	if err != nil {
		return err
	}
	if !st.ok {
		return fmt.Errorf("%s", st.msg)
	}
	return nil
}

func (self *Client) update(host string, domain string, ip string) error {
	if err := self.send("MODIP\n"); err != nil {
		return err
	}
	if err := self.send("HOSTNAME:" + host  +"\n"); err != nil {
		return err
	}
	if err := self.send("DOMNAME:" + domain  +"\n"); err != nil {
		return err
	}
	if err := self.send("IPV4:" + ip  +"\n"); err != nil {
		return err
	}
	if err := self.send(".\n"); err != nil {
		return err
	}

	st, err := self.getStatus()
	if err != nil {
		return err
	}
	if !st.ok {
		return fmt.Errorf("%s", st.msg)
	}
	return nil
}

func (self *Client) getStatus() (*status, error) {
	timeout := 10 * time.Second
	ret, _, err := self.expect(regexp_status, timeout)
	if err != nil {
		return nil, err
	}

	ret_s := strings.SplitN(ret, "\n", 2)
	if ret_s == nil {
		return nil, fmt.Errorf("cant parse response, %s", ret)
	}
	if len(ret_s) < 1 {
		return nil, fmt.Errorf("cant parse response, %s", ret)
	}
	st := strings.SplitN(ret_s[0], " ", 2)
	if len(st) < 2 {
		return nil, fmt.Errorf("cant parse response, %s", ret)
	}

	ok := (st[0] == stutsSucess)
	msg := st[1]
	return &status{ok:ok, msg:msg}, nil
}

func (self *Client) expect(rxp *regexp.Regexp, timeout time.Duration) (string, []string, error) {
	if self.exp == nil {
		return "", nil, ErrNotConnect
	}
	return self.exp.Expect(rxp, timeout)
}

func (self *Client) send(msg string) error {
	if self.exp == nil {
		return ErrNotConnect
	}

	return self.exp.Send(msg)
}

func (self *Client) lock() {
	self.mtx.Lock()
}

func (self *Client) unlock() {
	self.mtx.Unlock()
}

type status struct {
	ok  bool
	msg string
}

func createTlsExpect(ctx context.Context, address string, timeout time.Duration) (*expect.GExpect, error) {
	dr := new(tls.Dialer)
	client, err := dr.DialContext(ctx, "tcp", address)
	if err != nil {
		return nil, err
	}

	resCh := make(chan error)
	exp, _, err := expect.SpawnGeneric(&expect.GenOptions{
		In:  client,
		Out: client,
		Wait: func() error {
			return <-resCh
		},
		Close: func() error {
			close(resCh)
			return client.Close()
		},
		Check: func() bool {
			select {
			case <- ctx.Done():
				return false
			default:
			}
			return true
		},
	}, timeout, expect.Verbose(debug))

	if err != nil {
		return nil, err
	}
	return exp, nil
}
