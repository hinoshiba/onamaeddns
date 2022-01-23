go-onamaeddns
===

* [onamae.com](https://help.onamae.com/answer/7920) ddns client on Linux And macOS.
	* like a [this client](https://help.onamae.com/answer/7920) at Windows.

# Usage

### 1. macOS GUI
tba
### 2. [linux/macOS CLI](./usage-cli.md)

### 3. docker
tba
### 4. Library of Go.

```
package main

import (
	"log"
	"time"
)

import (
	"github.com/hinoshiba/go-onamaeddns/"
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


# Maintenance

## build

```
$ cd <repository>/docker/
$ make
## use binary on '<repository>/bin/...'
```


## References

* [LinuxやMacで お名前.com ダイナミックDNS の IPアドレスを更新する https://qiita.com/ats124/items/59ec0f444d00bbcea27d](https://qiita.com/ats124/items/59ec0f444d00bbcea27d)
