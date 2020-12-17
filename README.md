# digg [![GoDoc](https://godoc.org/github.com/sampointer/digg?status.svg)](https://godoc.org/github.com/sampointer/digg) [![Go Report Card](https://goreportcard.com/badge/github.com/sampointer/digg)](https://goreportcard.com/report/github.com/sampointer/digg)

Look-up region and other information for any AWS-owned IP address:

```bash
$ digg $(dig netflix.com +short)
prefix: 52.208.0.0/13 region: eu-west-1 service: AMAZON, network_border_group: eu-west-1
prefix: 52.208.0.0/13 region: eu-west-1 service: EC2, network_border_group: eu-west-1
prefix: 52.18.0.0/15 region: eu-west-1 service: AMAZON, network_border_group: eu-west-1
...
```

```bash
$ digg 52.94.76.5 2a05:d07a:a0ff:ffff:ffff:ffff:ffff:aaaa
prefix: 52.94.76.0/22 region: us-west-2 service: AMAZON, network_border_group: us-west-2
prefix: 2a05:d07a:a000::/40 region: eu-south-1 service: AMAZON, network_border_group: eu-south-1
prefix: 2a05:d07a:a000::/40 region: eu-south-1 service: S3, network_border_group: eu-south-1
```

## Installation

### Homebrew

```bash
brew tap sampointer/digg
brew install digg
```

### Packages
Debian and RPM packages can be found on the [releases][3] page.

[1]: https://ip-ranges.amazonaws.com/ip-ranges.json
[2]: https://docs.aws.amazon.com/general/latest/gr/aws-ip-ranges.html
[3]: https://github.com/sampointer/digg/releases/

### Docker

```bash
git clone https://github.com/sampointer/digg; cd digg
docker build -t digg .
docker run --rm -it digg $(dig netflix.com +short)
```
