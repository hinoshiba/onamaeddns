Usage:  Linux/macOS CLI
===

* `onaamaddns`ファイルには、グローバルIPアドレスを通知する機能はなく、単なるDDNSクライアントとして、指定されたIPアドレスを更新します
* グローバルIPアドレスを指定するスクリプトは手動で設定する必要があります

# Usage
```
[user@host]$ onamaeddns [-c <path of credential>] <host> <domain> <ip address>
Usage of bin/onamaeddns:
  -c string
        path of credential. default "~/.onamaeddns"
```

# Setup

1. 以下から、バイナリをダウンロードし、実行権限の付与、PATHを通します
	* https://github.com/hinoshiba/onamaeddns/releases
2. `~/.onamaeddns` に、認証情報を設置します
	* [sample: .onamaeddns](../sample/userhome/.onamaeddns)
3. お好みで、実行スクリプトを作成し、実行します
	* グローバルIPアドレスを取得する例
		* [sample: ddns_clnt.sh](../sample/usr/local/bin/ddns_clnt.sh)
