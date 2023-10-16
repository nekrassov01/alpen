package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/nekrassov01/access-log-parser"
	"github.com/urfave/cli/v2"
)

type OutputFormat int

const (
	Text OutputFormat = iota
	JSON
	PrettyJSON
)

func (o OutputFormat) String() string {
	if o >= 0 && int(o) < len(outputFormat) {
		return outputFormat[o]
	}
	return ""
}

var (
	outputFormat = []string{
		"text",
		"json",
		"pretty-json",
	}

	bufferFlag = &cli.StringFlag{
		Name:    "buffer",
		Aliases: []string{"b"},
		Usage:   "input from buffer",
	}

	fileFlag = &cli.StringFlag{
		Name:    "file-path",
		Aliases: []string{"f"},
		Usage:   "input from file path",
	}

	gzipFlag = &cli.StringFlag{
		Name:    "gzip-path",
		Aliases: []string{"g"},
		Usage:   "input from gzip file path",
	}

	zipFlag = &cli.StringFlag{
		Name:    "zip-path",
		Aliases: []string{"z"},
		Usage:   "input from zip file path",
	}

	outputFlag = &cli.StringFlag{
		Name:    "output",
		Aliases: []string{"o"},
		Usage:   fmt.Sprintf("select output format: %s", pipeJoin(outputFormat)),
		Value:   Text.String(),
	}

	skipFlag = &cli.IntSliceFlag{
		Name:    "skip",
		Aliases: []string{"s"},
		Usage:   "skip records by index",
	}

	metadataFlag = &cli.BoolFlag{
		Name:    "metadata",
		Aliases: []string{"m"},
		Usage:   "enable metadata output",
	}

	globFlag = &cli.StringFlag{
		Name:    "glob-pattern",
		Aliases: []string{"G"},
		Usage:   "filter glob pattern: available for parsing zip only",
		Value:   "*",
	}
)

func NewApp() *cli.App {
	return &cli.App{
		Name:        Name,
		Usage:       "AWS log parser/encoder",
		Version:     Version,
		Description: "A cli application for parsing AWS access logs",
		Action:      func(c *cli.Context) error { return nil },
		Commands: []*cli.Command{
			{
				Name:        "s3",
				Description: "Parses S3 access logs and converts them to structured formats",
				Usage:       "Parses S3 access logs",
				UsageText:   fmt.Sprintf("%s s3", Name),
				Action:      doS3Action,
				Flags:       []cli.Flag{bufferFlag, fileFlag, gzipFlag, zipFlag, outputFlag, skipFlag, metadataFlag, globFlag},
				Before:      validateFlags,
			},
			{
				Name:        "cf",
				Description: "Parses CloudFront access logs and converts them to structured formats",
				Usage:       "Parses CloudFront access logs",
				UsageText:   fmt.Sprintf("%s cf", Name),
				Action:      doCFAction,
				Flags:       []cli.Flag{bufferFlag, fileFlag, gzipFlag, zipFlag, outputFlag, skipFlag, metadataFlag, globFlag},
				Before:      validateFlags,
			},
			{
				Name:        "alb",
				Description: "Parses ALB access logs and converts them to structured formats",
				Usage:       "Parses ALB access logs",
				UsageText:   fmt.Sprintf("%s alb", Name),
				Action:      doALBAction,
				Flags:       []cli.Flag{bufferFlag, fileFlag, gzipFlag, zipFlag, outputFlag, skipFlag, metadataFlag, globFlag},
				Before:      validateFlags,
			},
			{
				Name:        "nlb",
				Description: "Parses NLB access logs and converts them to structured formats",
				Usage:       "Parses NLB access logs",
				UsageText:   fmt.Sprintf("%s nlb", Name),
				Action:      doNLBAction,
				Flags:       []cli.Flag{bufferFlag, fileFlag, gzipFlag, zipFlag, outputFlag, skipFlag, metadataFlag, globFlag},
				Before:      validateFlags,
			},
			{
				Name:        "clb",
				Description: "Parses CLB access logs and converts them to structured formats",
				Usage:       "Parses CLB access logs",
				UsageText:   fmt.Sprintf("%s clb", Name),
				Action:      doCLBAction,
				Flags:       []cli.Flag{bufferFlag, fileFlag, gzipFlag, zipFlag, outputFlag, skipFlag, metadataFlag, globFlag},
				Before:      validateFlags,
			},
		},
	}
}

func doS3Action(c *cli.Context) error {
	return doAction(c, []string{
		"bucket_owner",
		"bucket",
		"time",
		"remote_ip",
		"requester",
		"request_id",
		"operation",
		"key",
		"request_uri",
		"http_status",
		"error_code",
		"bytes_sent",
		"object_size",
		"total_time",
		"turn_around_time",
		"referer",
		"user_agent",
		"version_id",
		"host_id",
		"signature_version",
		"cipher_suite",
		"authentication_type",
		"host_header",
		"tls_version",
		"access_point_arn",
		"acl_required",
	}, generateS3Patterns())
}

func doCFAction(c *cli.Context) error {
	return doAction(c, []string{
		"date",
		"time",
		"x_edge_location",
		"sc_bytes",
		"c_ip",
		"cs_method",
		"cs_host",
		"cs_uri_stem",
		"sc_status",
		"cs_referer",
		"cs_user_agent",
		"cs_uri_query",
		"cs_cookie",
		"x_edge_result_type",
		"x_edge_request_id",
		"x_host_header",
		"cs_protocol",
		"cs_bytes",
		"time_taken",
		"x_forwarded_for",
		"ssl_protocol",
		"ssl_cipher",
		"x_edge_response_result_type",
		"cs_protocol_version",
		"fle_status",
		"fle_encrypted_fields",
		"c_port",
		"time_to_first_byte",
		"x_edge_detailed_result_type",
		"sc_content_type",
		"sc_content_len",
		"sc_range_start",
		"sc_range_end",
	}, generateCFPatterns())
}

func doALBAction(c *cli.Context) error {
	return doAction(c, []string{
		"type",
		"time",
		"elb",
		"client_port",
		"target_port",
		"request_processing_time",
		"target_processing_time",
		"response_processing_time",
		"elb_status_code",
		"target_status_code",
		"received_bytes",
		"sent_bytes",
		"request",
		"user_agent",
		"ssl_ciphe",
		"ssl_protocol",
		"target_group_arn",
		"trace_id",
		"domain_name",
		"chosen_cert_arn",
		"matched_rule_priority",
		"request_creation_time",
		"actions_executed",
		"redirect_url",
		"error_reason",
		"target_port_list",
		"target_status_code_list",
		"classification",
		"classification_reason",
	}, generateALBPatterns())
}

func doNLBAction(c *cli.Context) error {
	return doAction(c, []string{
		"type",
		"version",
		"time",
		"elb",
		"listener",
		"client:port",
		"destination:port",
		"connection_time",
		"tls_handshake_time",
		"received_bytes",
		"sent_bytes",
		"incoming_tls_alert",
		"chosen_cert_arn",
		"chosen_cert_serial",
		"tls_cipher",
		"tls_protocol_version",
		"tls_named_group",
		"domain_name",
		"alpn_fe_protocol",
		"alpn_be_protocol",
		"alpn_client_preference_list",
		"tls_connection_creation_time",
	}, generateNLBPatterns())
}

func doCLBAction(c *cli.Context) error {
	return doAction(c, []string{
		"time",
		"elb",
		"client_port",
		"backend_port",
		"request_processing_time",
		"backend_processing_time",
		"response_processing_time",
		"elb_status_code",
		"backend_status_code",
		"received_bytes",
		"sent_bytes",
		"request",
		"user_agent",
		"ssl_cipher",
		"ssl_protocol",
	}, generateCLBPatterns())
}

func doAction(c *cli.Context, fields []string, patterns []*regexp.Regexp) error {
	p, err := load(c, fields, patterns)
	if err != nil {
		return err
	}
	result, results, err := dispatch(c, p)
	if err != nil {
		return err
	}
	printResult(c, result, results)
	return nil
}

func load(c *cli.Context, fields []string, patterns []*regexp.Regexp) (*parser.Parser, error) {
	switch c.String(outputFlag.Name) {
	case Text.String():
		return parser.New(fields, patterns, textLineHandler, textMetadataHandler), nil
	case JSON.String():
		return parser.New(fields, patterns, nil, nil), nil
	case PrettyJSON.String():
		return parser.New(fields, patterns, prettyJSONLineHandler, prettyJSONMetadataHandler), nil
	default:
		return nil, fmt.Errorf("cannot parse command line flags: invalid output format: allowed values: %s", pipeJoin(outputFormat))
	}
}

func dispatch(c *cli.Context, p *parser.Parser) (result *parser.Result, results []*parser.Result, err error) {
	switch {
	case c.IsSet(bufferFlag.Name):
		result, err = p.ParseString(c.String(bufferFlag.Name), c.IntSlice(skipFlag.Name))
		return result, nil, err
	case c.IsSet(fileFlag.Name):
		result, err = p.ParseFile(c.String(fileFlag.Name), c.IntSlice(skipFlag.Name))
		return result, nil, err
	case c.IsSet(gzipFlag.Name):
		result, err = p.ParseGzip(c.String(gzipFlag.Name), c.IntSlice(skipFlag.Name))
		return result, nil, err
	case c.IsSet(zipFlag.Name):
		results, err = p.ParseZipEntries(c.String(zipFlag.Name), c.IntSlice(skipFlag.Name), c.String(globFlag.Name))
		return nil, results, err
	default:
		return nil, nil, fmt.Errorf("cannot parse command line flags: no valid input provided: %s", pipeJoin([]string{bufferFlag.Name, fileFlag.Name, gzipFlag.Name, zipFlag.Name}))
	}
}

func printResult(c *cli.Context, result *parser.Result, results []*parser.Result) {
	w := func(c *cli.Context, r *parser.Result) {
		fmt.Println(strings.Join(r.Data, "\n"))
		if c.Bool(metadataFlag.Name) {
			fmt.Println(r.Metadata)
		}
	}
	switch {
	case result != nil && results == nil:
		w(c, result)
	case result == nil && results != nil:
		for _, r := range results {
			w(c, r)
		}
	default:
	}
}

func validateFlags(c *cli.Context) error {
	if err := isSingle(c, bufferFlag.Name, fileFlag.Name, gzipFlag.Name, zipFlag.Name); err != nil {
		return err
	}
	return isValidPair(c, zipFlag.Name, globFlag.Name)
}

func isSingle(c *cli.Context, flags ...string) error {
	i := 0
	for _, flag := range flags {
		if c.IsSet(flag) {
			i++
		}
	}
	if i > 1 {
		return fmt.Errorf("cannot parse command line flags: only one flag can be used: %s", pipeJoin(flags))
	}
	return nil
}

func isValidPair(c *cli.Context, flagA, flagB string) error {
	if !c.IsSet(flagA) && c.IsSet(flagB) {
		return fmt.Errorf("cannot parse command line flags: `%s` is available for `%s` only", flagB, flagA)
	}
	return nil
}

func pipeJoin(s []string) string {
	return strings.Join(s, "|")
}
