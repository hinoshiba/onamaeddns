go-onamaeddns
===

* [onamae.com](https://help.onamae.com/answer/7920) ddns client.
	* like a [this client](https://help.onamae.com/answer/7920).

## usage

```
[user@host]$ onamaeddns [-c <path of credential>] <host> <domain> <ip address>
Usage of bin/onamaeddns:
  -c string
        path of credential. default "~/.onamaeddns"
```

## as a package

```
package main

import (
	"log"
	"time"
)

import (
	"github.com/hinoshiba/go-onamaeddns/src/onamaeddns"
)

func main() {
	cl, err := onamaeddns.Dial("ddnsclient.onamae.com:65010", "username", "password", time.Minute)
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
```


## References

* [LinuxやMacで お名前.com ダイナミックDNS の IPアドレスを更新する https://qiita.com/ats124/items/59ec0f444d00bbcea27d](https://qiita.com/ats124/items/59ec0f444d00bbcea27d)
