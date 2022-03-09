release
===

```bash
$ cd <repository root>
git tag -a v[0-9].[0-9].[0-9] -m ''
git push origin v[0-9].[0-9].[0-9]
## auto run: ./github/workflows/release.yml
# Edit release of draft on https://github.com/hinoshiba/onamaeddns/releases
# And, Check to dockerhub image: https://hub.docker.com/repository/docker/hinoshiba/onamaeddns
```
