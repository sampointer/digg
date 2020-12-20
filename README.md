# digg [![GoDoc](https://godoc.org/github.com/sampointer/digg?status.svg)](https://godoc.org/github.com/sampointer/digg) [![Go Report Card](https://goreportcard.com/badge/github.com/sampointer/digg)](https://goreportcard.com/report/github.com/sampointer/digg)

Look-up region and other information for any Google-owned IP address:

```bash
$ digg $(dig king.com +short)
prefix: 34.64.0.0/10 scope:  service:
prefix: 34.120.0.0/16 scope: global service: Google Cloud
```

```bash
$ digg 8.8.8.8 2a00:1450:4009:814::200e
prefix: 8.8.8.0/24 scope:  service:
prefix: 2a00:1450::/32 scope:  service:
```

## Installation

### Homebrew

```bash
brew tap sampointer/digg
brew install digg
```

### Packages
Debian and RPM packages can be found on the [releases][1] page.

### Docker

```bash
git clone https://github.com/sampointer/digg; cd digg
docker build -t digg .
docker run --rm -it digg $(dig king.com +short)
```

## Similar tools

| Company  | Tool        |
|----------|-------------|
| AWS      | [digaws][a] |
| Google   | [digg][g]   |

[1]: https://github.com/sampointer/digg/releases/

[a]: https://github.com/sampointer/digaws
[g]: https://github.com/sampointer/digg
