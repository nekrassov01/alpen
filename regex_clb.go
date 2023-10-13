package main

import (
	"regexp"
	"strings"
)

const clbSep = " "

var (
	clbFields = []string{
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
	}

	clbPatternV1 = []string{
		`^([!-~]+)`,
		`([!-~]+)`,
		`([!-~]+)`,
		`([!-~]+)`,
		`([\d\-.]+)`,
		`([\d\-.]+)`,
		`([\d\-.]+)`,
		`(\d{1,3}|-)`,
		`(\d{1,3}|-)`,
		`([\d\-.]+)`,
		`([\d\-.]+)`,
		`"([ -~]+)"`,
	}

	clbPatternV2 = append(
		clbPatternV1,
		[]string{
			`"([ -~]+)"`,
			`([!-~]+)`,
			`([!-~]+)`,
		}...,
	)

	clbPatterns = []*regexp.Regexp{
		regexp.MustCompile(strings.Join(clbPatternV1, clbSep)),
		regexp.MustCompile(strings.Join(clbPatternV2, clbSep)),
	}
)
