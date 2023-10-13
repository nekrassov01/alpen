package main

import (
	"regexp"
	"strings"
)

const albSep = " "

var (
	albFields = []string{
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
	}

	albPatternV1 = []string{
		`^([!-~]+)`,
		`([!-~]+)`,
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
		`"([ -~]+)"`,
		`([!-~]+)`,
		`([!-~]+)`,
		`([!-~]+)`,
		`"([ -~]+)"`,
		`"([ -~]+)"`,
		`"([ -~]+)"`,
		`([!-~]+)`,
		`([!-~]+)`,
		`"([ -~]+)"`,
		`"([ -~]+)"`,
		`"([ -~]+)"`,
		`"([ -~]+)"`,
		`"([ -~]+)"`,
		`"([ -~]+)"`,
		`"([ -~]+)"`,
	}

	albPatterns = []*regexp.Regexp{
		regexp.MustCompile(strings.Join(albPatternV1, albSep)),
	}
)
