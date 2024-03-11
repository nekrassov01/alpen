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
- TSV format

Usage
-----

```text
NAME:
   alpen - Access log parser/encoder CLI

USAGE:
   alpen [global options] command [command options] [arguments...]

VERSION:
   0.0.0

DESCRIPTION:
   A cli application for parsing various access logs

COMMANDS:
   completion  Generate completion scripts for specified shell: bash|zsh|pwsh
   clf         Parses apache common/combined log format
   clfv        Parses apache common/combined log format with vhost
   s3          Parses S3 access logs
   cf          Parses CloudFront access logs
   alb         Parses ALB access logs
   nlb         Parses NLB access logs
   clb         Parses CLB access logs
   ltsv        Parses LTSV format logs

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

```text
NAME:
   alpen s3 - Parses S3 access logs

USAGE:
   alpen s3

DESCRIPTION:
   Parses S3 access logs and converts them to structured formats

OPTIONS:
   --input value, -i value                                  select input type: default|gz|zip (default: "default")
   --output value, -o value                                 select output type: json|pretty-json|text|ltsv|tsv (default: "json")
   --result, -r                                             enable result output (default: false)
   --glob value, -g value                                   filter glob pattern: available for parsing zip only (default: "*")
   --labels value, -l value [ --labels value, -l value ]    select labels to output with labels
   --filters value, -f value [ --filters value, -f value ]  set filter expressions: allowed operator: >|>=|<|<=|==|!=|==*|!=*|=~|!~|=~*|!~*
   --skip value, -s value [ --skip value, -s value ]        skip lines by line number
   --prefix, -p                                             enable line prefix: PROCESSED|UNMATCHED (default: false)
   --unmatch, -u                                            enable output of unmatched lines (default: false)
   --num, -n                                                set line number at the beginning of the line (default: false)
   --help, -h                                               show help
```

Example
-------

```sh
# Read and convert s3 logs from file and convert to default NDJSON format
alpen s3 "sample_s3.log"

# Set line number at the beginning of line
alpen s3 -n "sample_s3.log"

# Read s3 log from file, convert to pretty NDJSON and output parsed results
alpen s3 -r  -o pretty-json "sample_s3.log"

# Can be combined with tail -f to process standard input
# Results are consistent, even if interrupted with CTRL+C
tail -f sample_s3.log | alpen s3 -r

# Convert LTSV format
alpen s3 -r -o ltsv "sample_s3.log"

# In TSV format, the header is set from the parsing result of the first line
alpen s3 -r -o tsv "sample_s3.log"

# Read CloudFront logs from gzip file and skip header lines
alpen cf -r -s 1,2 -i gz "sample_cloudfront.log.gz"

# Read ALB logs from zip file and convert all entries with `.log` extension
alpen alb -r -g "*.log" -i zip "sample_alb.log.zip"

# Unmatched lines can also be output raw and made explicit by line prefix
alpen s3 -u -p "sample_s3.log"

# Columns can be narrowed by specifying labels
alpen s3 -l bucket,method,request_uri,protocol "sample_s3.log"

# Filter expressions to narrow down rows
#   > >= == <= <  (arithmetic (float64))
#   == ==* != !=* (string comparison (string))
#   =~ !~ =~* !~* (regular expression (string))
# inspired from <https://github.com/sonots/lltsv>
alpen s3 -f "method == GET,operation =~ .*BUCKETPOLICY"

# Read apache common/combined format logs
# Matches both common/combined by default
# Use space or tab as delimiter
alpen clf "sample_clf.log"

# Read apache common/combined log format with virtual host
# Matches if virtual host is at the beginning
alpen clfv "sample_clf.log"

# LTSV uses labels as names, so it is not possible to decompose a request into
# methods, request_uri, or protocols.
alpen ltsv "sample_ltsv.log"
```

Installation
------------

Install with homebrew

```sh
brew install nekrassov01/tap/alpen
```

Install with go

```sh
go install github.com/nekrassov01/alpen
```

Or download binary from [releases](https://github.com/nekrassov01/alpen/releases)

Shell completion
----------------

Supported shells are as follows:

- bash
- zsh
- pwsh

```sh
alpen completion bash|zsh|pwsh
```

Author
------

[nekrassov01](https://github.com/nekrassov01)

License
-------

[MIT](https://github.com/nekrassov01/alpen/blob/main/LICENSE)
