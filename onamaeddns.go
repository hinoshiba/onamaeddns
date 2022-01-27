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
	DEBUG bool = false

	StutsSucess string = "000"
	StutsFailed string = "001"
)

var (
	STATUS_REGEXP *regexp.Regexp = regexp.MustCompile("\\d{3}")
	STATUS_WAIT   = 10 * time.Second

	ErrNotConected error = fmt.Errorf("tls not connected")
)

type Client struct {
	exp    *expect.GExpect
	ctx    context.Context
	mtx    *sync.Mutex
}

func Dial(sv string, user string, pass string, timeout time.Duration) (*Client, error) {
	return dial(context.Background(), sv, user, pass, timeout)
}

func DialWithContext(ctx context.Context, sv string, user string, pass string, timeout time.Duration) (*Client, error) {
	return dial(ctx, sv, user, pass, timeout)
}

func dial(ctx context.Context, sv string, user string, pass string, timeout time.Duration) (*Client, error) {
	exp, err := createTlsExpect(ctx, sv, timeout)
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

func (self *Client) UpdateIPv4(host string, dom string, ip string) error {
	self.lock()
	defer self.unlock()

	if self.isCanceledContext() {
		return fmt.Errorf("closed context.")
	}
	return self.update(host, dom, ip)
}

func (self *Client) Expect(rxp *regexp.Regexp, timeout time.Duration) (string, []string, error) {
	self.lock()
	defer self.unlock()

	if self.isCanceledContext() {
		return "", nil, fmt.Errorf("closed context.")
	}
	return self.expect(rxp, timeout)
}

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

func (self *Client) update(host string, dom string, ip string) error {
	if err := self.send("MODIP\n"); err != nil {
		return err
	}
	if err := self.send("HOSTNAME:" + host  +"\n"); err != nil {
		return err
	}
	if err := self.send("DOMNAME:" + dom  +"\n"); err != nil {
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
	ret, _, err := self.expect(STATUS_REGEXP, STATUS_WAIT)
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

	ok := (st[0] == StutsSucess)
	msg := st[1]
	return &status{ok:ok, msg:msg}, nil
}

func (self *Client) expect(rxp *regexp.Regexp, timeout time.Duration) (string, []string, error) {
	if self.exp == nil {
		return "", nil, ErrNotConected
	}
	return self.exp.Expect(rxp, timeout)
}

func (self *Client) send(msg string) error {
	if self.exp == nil {
		return ErrNotConected
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

func createTlsExpect(ctx context.Context, sv string, timeout time.Duration) (*expect.GExpect, error) {
	dr := new(tls.Dialer)
	client, err := dr.DialContext(ctx, "tcp", sv)
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
	}, timeout, expect.Verbose(DEBUG))

	if err != nil {
		return nil, err
	}
	return exp, nil
}
