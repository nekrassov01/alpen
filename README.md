alpen
=====

alpen is aws log parser/encoder.

Supported
---------

- Amazon S3
- Amazon CloudFront
- Application Load Balancer
- Network Load Balancer
- Classic Load Balancer

Usage
-----

```text
NAME:
   alpen - AWS access log parser

USAGE:
   alpen [global options] command [command options] [arguments...]

VERSION:
   0.0.0

DESCRIPTION:
   A cli application for parsing AWS access logs

COMMANDS:
   s3       Parses S3 access logs
   cf       Parses CloudFront access logs
   alb      Parses ALB access logs
   nlb      Parses NLB access logs
   clb      Parses CLB access logs
   help, h  Shows a list of commands or help for one command

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
   --buffer value, -b value                           input from buffer
   --file-path value, -f value                        input from file path
   --gzip-path value, -g value                        input from gzip file path
   --zip-path value, -z value                         input from zip file path
   --output value, -o value                           select output format: text|json|pretty-json (default: "text")
   --skip value, -s value [ --skip value, -s value ]  skip records by index
   --metadata, -m                                     enable metadata output (default: false)
   --glob-pattern value, -G value                     filter glob pattern: available for parsing zip only (default: "*")
   --help, -h                                         show help
```

Installation
------------

Download binary from the release page or install it with the following command:

```sh
go install github.com/nekrassov01/alpen
```

Author
------

[nekrassov01](https://github.com/nekrassov01)

License
-------

[MIT](https://github.com/nekrassov01/alpen/blob/main/LICENSE)
