alpen
=====

[![CI](https://github.com/nekrassov01/alpen/actions/workflows/ci.yml/badge.svg)](https://github.com/nekrassov01/alpen/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/nekrassov01/alpen)](https://goreportcard.com/report/github.com/nekrassov01/alpen)
![GitHub](https://img.shields.io/github/license/nekrassov01/alpen)
![GitHub](https://img.shields.io/github/v/release/nekrassov01/alpen)

alpen is a CLI application for parsing and encoding various access logs

Supported
---------

- Apache common/combined log format
- Apache common/combined log format with virtual host
- Amazon S3 access log format
- Amazon CloudFront access log format
- AWS Application Load Balancer access log format
- AWS Network Load Balancer access log format
- AWS Classic Load Balancer access log format
- LTSV format

Usage
-----

```text
NAME:
   alpen - Access log parser/encoder CLI

USAGE:
   alpen [global options] command [command options] [arguments...]

VERSION:
   0.0.16

DESCRIPTION:
   A cli application for parsing various access logs

COMMANDS:
   clf   Parses apache common/combined log format
   clfv  Parses apache common/combined log format with vhost
   s3    Parses S3 access logs
   cf    Parses CloudFront access logs
   alb   Parses ALB access logs
   nlb   Parses NLB access logs
   clb   Parses CLB access logs
   ltsv  Parses LTSV format logs

GLOBAL OPTIONS:
   --completion value, -c value  select a shell to display completion scripts: bash|zsh|pwsh
   --help, -h                    show help
   --version, -v                 print the version
```

```text
NAME:
   alpen clf - Parses apache common/combined log format

USAGE:
   alpen clf

DESCRIPTION:
   Parses apache common/combined log format and converts them to structured formats

OPTIONS:
   --input value, -i value                            input from string
   --file-path value, -f value                        input from file path
   --gzip-path value, -g value                        input from gzip file path
   --zip-path value, -z value                         input from zip file path
   --output value, -o value                           select output format: json|pretty-json|text|ltsv (default: "json")
   --skip value, -s value [ --skip value, -s value ]  skip records by index
   --metadata, -m                                     enable metadata output (default: false)
   --line-number, -l                                  set line number at the beginning of the line (default: false)
   --glob-pattern value, -G value                     filter glob pattern: available for parsing zip only (default: "*")
   --help, -h                                         show help
```

Example
-------

```sh
# Read and convert s3 logs from file and convert to default NDJSON format
alpen s3 -f "sample_s3.log"

# Set line number at the beginning of line, like "index": "n"
alpen s3 -f "sample_s3.log" -l

# Read s3 logs from file and convert to pretty NDJSON, also output metadata
alpen s3 -f "sample_s3.log" -o pretty-json -m

# Convert LTSV format
alpen s3 -f "sample_s3.log" -o ltsv -m

# Read CloudFront logs from gzip file and skip header lines
alpen cf -g "sample_cloudfront.log.gz" -s 1,2

# Read ALB logs from zip file and convert all entries with `.log` extension
alpen alb -z "sample_alb.log.zip" -G "*.log"

# Read apache common/combined format logs
# Matches both common/combined by default
# Use space or tab as delimiter
alpen clf -f "sample_clf.log"

# Read apache common/combined log format with virtual host
# Matches if virtual host is at the beginning
alpen clfv -f "sample_clf.log"

# LTSV uses labels as names, so it is not possible to decompose a request into
# methods, request_uri, or protocols.
alpen ltsv -f "sample_ltsv.log"
```

Installation
------------

Download binary from the release page or install it with the following command:

```sh
go install github.com/nekrassov01/alpen@latest
```

Shell completion
----------------

Supported shells are as follows:

- bash
- zsh
- pwsh

```sh
alpen --completion bash|zsh|pwsh
```

Author
------

[nekrassov01](https://github.com/nekrassov01)

License
-------

[MIT](https://github.com/nekrassov01/alpen/blob/main/LICENSE)
