package main

import (
	"regexp"
	"strings"
)

const nlbSep = " "

var (
	nlbFields = []string{
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
	}

	nlbPatternV1 = []string{
		`^([!-~]+)`,
		`([!-~]+)`,
		`([!-~]+)`,
		`([!-~]+)`,
		`([!-~]+)`,
		`([!-~]+)`,
		`([!-~]+)`,
		`([\d\-.]+)`,
		`([\d\-.]+)`,
		`([!-~]+)`,
		`([!-~]+)`,
		`([!-~]+)`,
		`([!-~]+)`,
		`([ -~]+)`,
		`(\S+)`,
		`([!-~]+)`,
		`([!-~]+)`,
		`([!-~]+)`,
		`([!-~]+)`,
		`([!-~]+)`,
		`([ -~]+)`,
		`([!-~]+)`,
	}

	nlbPatterns = []*regexp.Regexp{
		regexp.MustCompile(strings.Join(nlbPatternV1, nlbSep)),
	}
)
