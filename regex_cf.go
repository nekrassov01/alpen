package main

import (
	"regexp"
	"strings"
)

const cfSep = "\t"

var (
	cfFields = []string{
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
	}

	cfPatternV1 = []string{
		`^([\d\-.:]+)`,
		`([\d\-.:]+)`,
		`([ -~]+)`,
		`([\d\-.]+)`,
		`([ -~]+)`,
		`([ -~]+)`,
		`([ -~]+)`,
		`([ -~]+)`,
		`(\d{1,3}|-)`,
		`([ -~]+)`,
		`([ -~]+)`,
		`([ -~]+)`,
		`(\S+)`,
		`([ -~]+)`,
		`([ -~]+)`,
		`([ -~]+)`,
		`([ -~]+)`,
		`([\d\-.]+)`,
		`([\d\-.]+)`,
		`([ -~]+)`,
		`([ -~]+)`,
		`([ -~]+)`,
		`([ -~]+)`,
		`([ -~]+)`,
		`([ -~]+)`,
		`(\S+)`,
		`([\d\-.]+)`,
		`([\d\-.]+)`,
		`([ -~]+)`,
		`([ -~]+)`,
		`([\d\-.]+)`,
		`([\d\-.]+)`,
		`([\d\-.]+)`,
	}

	cfPatterns = []*regexp.Regexp{
		regexp.MustCompile(strings.Join(cfPatternV1, cfSep)),
	}
)
