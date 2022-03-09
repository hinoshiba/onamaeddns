build
===

# binary
```bash
$ cd <repository root>
$ make build-bin ## build to '<repository>/bin/...'
```

# docker-image
```bash
$ cd <repository root>
$ make build-docker
```

## check on docker compose
```bash
$ cd <repository root>
$ vim <repository>/docker-compose.yml # edit to enable debug
$ docker-compose up
```
