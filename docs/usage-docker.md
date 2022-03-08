Usage: docker
===

[onamaeddns](https://hub.docker.com/repository/docker/hinoshiba/onamaeddns) ![dokcerimage-lastversion](https://img.shields.io/docker/v/hinoshiba/onamaeddns.svg)  


本イメージでは、[exec_ddns.sh](../docker-in/exec_ddns.sh) が実行されます。  
300秒おきに、globalip.me へ実行場所のグローバルIPアドレスの取得を行い、過去の取得内容と相違がある場合に、`${TARGET_HOST}.${TARGET_DOMAIN}`を更新します。  
なお、docker container起動時は、そのcontainer上には前回が存在しないので、必ずお名前.comへ更新を行います。  

# dokcer

## 動作準備

* お名前.com認証情報の、`<username>:<password>` の内容のファイルをどこかに設置します
	* 例えば、`/home/hinoshiba/.onamaeddns-cred`等

## 実行

```bash
docker run -it --rm --mount type=bind,src=<credfile>,dst=/etc/onamaeddns/cred,ro -e TARGET_HOST="<yourhost>" -e TARGET_DOMAIN="<yourdomain>" hinoshiba/onamaeddns:<version>
```

### example
```bash
docker run -it --rm --mount type=bind,src=/home/hinoshiba/.onamaeddns-cred,dst=/etc/onamaeddns/cred,ro -e TARGET_HOST="superhost" -e TARGET_DOMAIN="example.com" hinoshiba/onamaeddns:v1.0.1
```


# docker-compose

## 動作準備

1. `/var/service/onamaeddns/etc/cred`
	* `<username>:<password>` 形式で、お名前.comの認証情報を設置します
		* docker-compose up時に、roマウントされて利用されます
2. 環境変数の設置
	* `TARGET_HOST`
		* 変更対象のホスト名を指定します
	* `TARGET_DOMAIN`
		* 変更対象のドメイン名を指定します

## 実行

本リポジトリのrootで、以下
```bash
docker-compose up
```

## 便利な活用方法: systemd 経由での実行

1. `/var/service/onamaeddns` を作成します
1. 本リポジトリの中身全てを、コピーします
1. `/var/service/onamaeddns/etc/cred` を作成します
	* `<username>:<password>` 形式で、お名前.comの認証情報を設置します
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
