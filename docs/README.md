go-onamaeddns
===

* Linux及び、macOSの、[onamae.com](https://help.onamae.com/answer/7920) DDNSクライアント と、そのライブラリ
	* Windowsで動作する [公式クライアント](https://help.onamae.com/answer/7920) のモノマネです
	* 有志が勝手に作っているので、ご利用は自己責任でお願いします
	* いくつかのサンプルやイメージのグローバルIPアドレス取得元は、`globalip.me` を活用しています

# Usage

本リポジトリの使い方は、4つの方法があります。昇順で簡単なので、お好みの使い方をしてください  

### 1. [docker](./usage-docker.md)
![https://img.shields.io/docker/v/hinoshiba/onamaeddns.svg]()
### 2. [linux/macOS CLI](./usage-cli.md)
### 3. Library of Go.

```
package main

import (
	"log"
	"time"
)

import (
	"github.com/hinoshiba/go-onamaeddns"
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
$ make
## build to '<repository>/bin/...' & docker images
$ vim <repository>/docker-compose.yml
# enable debug
$ docker-compose up
```


# References

* [LinuxやMacで お名前.com ダイナミックDNS の IPアドレスを更新する https://qiita.com/ats124/items/59ec0f444d00bbcea27d](https://qiita.com/ats124/items/59ec0f444d00bbcea27d)
