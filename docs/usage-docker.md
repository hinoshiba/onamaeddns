Usage: docker
===

* 本リポジトリrootで、`docker-compose up` で利用できます

## 動作準備

1. `/var/service/onamaeddns/etc/cred`
	* `username:password` 形式で、お名前.comの認証情報を設置します
		* docker-compose up時に、roマウントされて利用されます
2. 環境変数の設置
	* `TARGET_HOST`
		* 変更対象のホスト名を指定します
	* `TARGET_DOMAIN`
		* 変更対象のドメイン名を指定します

## 便利な活用方法: systemd 経由での実行

1. `/var/service/onamaeddns` を作成します
1. 本リポジトリの中身全てを、コピーします
1. `/var/service/onamaeddns/etc/cred` を作成します
	* `username:password` 形式で、お名前.comの認証情報を設置します
		* docker-compose up時に、roマウントされて利用されます
1. `本リポジトリ/sample/etc/system/systemd/onamaeddns.service*`を、`/etc/system/systemd/`へコピーします
1. `/etc/system/systemd/onamaeddns.service.d/env.conf`の環境変数を設定します
	* `TARGET_HOST`
		* 変更対象のホスト名を指定します
	* `TARGET_DOMAIN`
		* 変更対象のドメイン名を指定します
1. `systemd daemon-reload`
1. `systemd restart onamaeddns.service`
	* `/var/log/syslog` などに、ログが出力されます
