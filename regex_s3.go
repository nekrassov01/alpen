package main

import (
	"regexp"
	"strings"
)

const s3Sep = " "

var (
	s3Fields = []string{
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
	}

	s3PatternV1 = []string{
		`^([!-~]+)`,
		`([!-~]+)`,
		`(\[[ -~]+ [0-9+]+\])`,
		`([!-~]+)`,
		`([!-~]+)`,
		`([!-~]+)`,
		`([!-~]+)`,
		`([!-~]+)`,
		`"([ -~]+)"`,
		`(\d{1,3})`,
		`([!-~]+)`,
		`([\d\-.]+)`,
		`([\d\-.]+)`,
		`([\d\-.]+)`,
		`([\d\-.]+)`,
		`"([ -~]+)"`,
		`"([ -~]+)"`,
		`([!-~]+)`,
	}

	s3PatternV2 = append(
		s3PatternV1,
		[]string{
			`([!-~]+)`,
			`([!-~]+)`,
			`([!-~]+)`,
			`([!-~]+)`,
			`([!-~]+)`,
		}...,
	)

	s3PatternV3 = append(
		s3PatternV2,
		[]string{
			`([!-~]+)`,
		}...,
	)

	s3PatternV4 = append(
		s3PatternV3,
		[]string{
			`([!-~]+)$`,
		}...,
	)

	s3PatternV5 = append(
		s3PatternV4,
		[]string{
			`([!-~]+)$`,
		}...,
	)

	s3Patterns = []*regexp.Regexp{
		regexp.MustCompile(strings.Join(s3PatternV5, s3Sep)),
		regexp.MustCompile(strings.Join(s3PatternV4, s3Sep)),
		regexp.MustCompile(strings.Join(s3PatternV3, s3Sep)),
		regexp.MustCompile(strings.Join(s3PatternV2, s3Sep)),
		regexp.MustCompile(strings.Join(s3PatternV1, s3Sep)),
	}
)
