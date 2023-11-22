package main

import (
	"regexp"
	"strings"
)

func generateApacheCLFPatterns() []*regexp.Regexp {
	seps := []string{" ", "\t"}
	basePattern := []string{
		`^(?P<remote_host>\S+)`,
		`(?P<remote_logname>\S+)`,
		`(?P<remote_user>[\S ]+)`,
		`(?P<datetime>\[[^\]]+\])`,
		`\"(?P<method>[A-Z]+) (?P<request_uri>[^ \"]+) (?P<protocol>HTTP/[0-9.]+)\"`,
		`(?P<status>[0-9]{3})`,
		`(?P<size>[0-9]+|-)`,
	}
	combines := [][]string{
		{
			`"(?P<referer>[^\"]*)"`,
			`"(?P<user_agent>[^\"]*)"`,
		},
		{}, // for basePattern
	}
	patterns := make([]*regexp.Regexp, 0, len(combines)*len(seps))
	for _, sep := range seps {
		for _, combine := range combines {
			patterns = append(patterns, regexp.MustCompile(strings.Join(append(basePattern, combine...), sep)))
		}
	}
	return patterns
}

func generateApacheCLFWithVHostPatterns() []*regexp.Regexp {
	seps := []string{" ", "\t"}
	basePattern := []string{
		`^(?P<virtual_host>\S+)`,
		`(?P<remote_host>\S+)`,
		`(?P<remote_logname>\S+)`,
		`(?P<remote_user>[\S ]+)`,
		`(?P<datetime>\[[^\]]+\])`,
		`\"(?P<method>[A-Z]+) (?P<request_uri>[^ \"]+) (?P<protocol>HTTP/[0-9.]+)\"`,
		`(?P<status>[0-9]{3})`,
		`(?P<size>[0-9]+|-)`,
	}
	combines := [][]string{
		{
			`"(?P<referer>[^\"]*)"`,
			`"(?P<user_agent>[^\"]*)"`,
		},
		{}, // for basePattern
	}
	patterns := make([]*regexp.Regexp, 0, len(combines)*len(seps))
	for _, sep := range seps {
		for _, combine := range combines {
			patterns = append(patterns, regexp.MustCompile(strings.Join(append(basePattern, combine...), sep)))
		}
	}
	return patterns
}

func generateS3Patterns() (patterns []*regexp.Regexp) {
	sep := " "
	basePattern := []string{
		`^(?P<bucket_owner>[!-~]+)`,
		`(?P<bucket>[!-~]+)`,
		`(?P<time>\[[ -~]+ [0-9+]+\])`,
		`(?P<remote_ip>[!-~]+)`,
		`(?P<requester>[!-~]+)`,
		`(?P<request_id>[!-~]+)`,
		`(?P<operation>[!-~]+)`,
		`(?P<key>[!-~]+)`,
		`"(?P<request_uri>[ -~]+)"`,
		`(?P<http_status>\d{1,3})`,
		`(?P<error_code>[!-~]+)`,
		`(?P<bytes_sent>[\d\-.]+)`,
		`(?P<object_size>[\d\-.]+)`,
		`(?P<total_time>[\d\-.]+)`,
		`(?P<turn_around_time>[\d\-.]+)`,
		`"(?P<referer>[^\"]*)"`,
		`"(?P<user_agent>[^\"]*)"`,
		`(?P<version_id>[!-~]+)`,
	}
	additions := [][]string{
		{
			`(?P<host_id>[!-~]+)`,
			`(?P<signature_version>[!-~]+)`,
			`(?P<cipher_suite>[!-~]+)`,
			`(?P<authentication_type>[!-~]+)`,
			`(?P<host_header>[!-~]+)`,
			`(?P<tls_version>[!-~]+)`,
			`(?P<access_point_arn>[!-~]+)`,
			`(?P<acl_required>[!-~]+)`,
		},
		{
			`(?P<host_id>[!-~]+)`,
			`(?P<signature_version>[!-~]+)`,
			`(?P<cipher_suite>[!-~]+)`,
			`(?P<authentication_type>[!-~]+)`,
			`(?P<host_header>[!-~]+)`,
			`(?P<tls_version>[!-~]+)`,
			`(?P<access_point_arn>[!-~]+)`,
		},
		{
			`(?P<host_id>[!-~]+)`,
			`(?P<signature_version>[!-~]+)`,
			`(?P<cipher_suite>[!-~]+)`,
			`(?P<authentication_type>[!-~]+)`,
			`(?P<host_header>[!-~]+)`,
			`(?P<tls_version>[!-~]+)`,
		},
		{
			`(?P<host_id>[!-~]+)`,
			`(?P<signature_version>[!-~]+)`,
			`(?P<cipher_suite>[!-~]+)`,
			`(?P<authentication_type>[!-~]+)`,
			`(?P<host_header>[!-~]+)`,
		},
		{},
	}
	patterns = make([]*regexp.Regexp, len(additions))
	for i, addition := range additions {
		patterns[i] = regexp.MustCompile(strings.Join(append(basePattern, addition...), sep))
	}
	return patterns
}

func generateCFPatterns() (patterns []*regexp.Regexp) {
	sep := "\t"
	basePattern := []string{
		`^(?P<date>[\d\-.:]+)`,
		`(?P<time>[\d\-.:]+)`,
		`(?P<x_edge_location>[ -~]+)`,
		`(?P<sc_bytes>[\d\-.]+)`,
		`(?P<c_ip>[ -~]+)`,
		`(?P<cs_method>[ -~]+)`,
		`(?P<cs_host>[ -~]+)`,
		`(?P<cs_uri_stem>[ -~]+)`,
		`(?P<sc_status>\d{1,3}|-)`,
		`(?P<cs_referer>[^\"]*)`,
		`(?P<cs_user_agent>[^\"]*)`,
		`(?P<cs_uri_query>[ -~]+)`,
		`(?P<cs_cookie>\S+)`,
		`(?P<x_edge_result_type>[ -~]+)`,
		`(?P<x_edge_request_id>[ -~]+)`,
		`(?P<x_host_header>[ -~]+)`,
		`(?P<cs_protocol>[ -~]+)`,
		`(?P<cs_bytes>[\d\-.]+)`,
		`(?P<time_taken>[\d\-.]+)`,
		`(?P<x_forwarded_for>[ -~]+)`,
		`(?P<ssl_protocol>[ -~]+)`,
		`(?P<ssl_cipher>[ -~]+)`,
		`(?P<x_edge_response_result_type>[ -~]+)`,
		`(?P<cs_protocol_version>[ -~]+)`,
		`(?P<fle_status>[ -~]+)`,
		`(?P<fle_encrypted_fields>\S+)`,
		`(?P<c_port>[\d\-.]+)`,
		`(?P<time_to_first_byte>[\d\-.]+)`,
		`(?P<x_edge_detailed_result_type>[ -~]+)`,
		`(?P<sc_content_type>[ -~]+)`,
		`(?P<sc_content_len>[\d\-.]+)`,
		`(?P<sc_range_start>[\d\-.]+)`,
		`(?P<sc_range_end>[\d\-.]+)`,
	}
	return []*regexp.Regexp{
		regexp.MustCompile(strings.Join(basePattern, sep)),
	}
}

func generateALBPatterns() []*regexp.Regexp {
	sep := " "
	basePattern := []string{
		`^(?P<type>[!-~]+)`,
		`(?P<time>[!-~]+)`,
		`(?P<elb>[!-~]+)`,
		`(?P<client_port>[!-~]+)`,
		`(?P<target_port>[!-~]+)`,
		`(?P<request_processing_time>[\d\-.]+)`,
		`(?P<target_processing_time>[\d\-.]+)`,
		`(?P<response_processing_time>[\d\-.]+)`,
		`(?P<elb_status_code>\d{1,3}|-)`,
		`(?P<target_status_code>\d{1,3}|-)`,
		`(?P<received_bytes>[\d\-.]+)`,
		`(?P<sent_bytes>[\d\-.]+)`,
		`"(?P<request>[ -~]+)"`,
		`"(?P<user_agent>[^\"]*)"`,
		`(?P<ssl_cipher>[!-~]+)`,
		`(?P<ssl_protocol>[!-~]+)`,
		`(?P<target_group_arn>[!-~]+)`,
		`"(?P<trace_id>[ -~]+)"`,
		`"(?P<domain_name>[ -~]+)"`,
		`"(?P<chosen_cert_arn>[ -~]+)"`,
		`(?P<matched_rule_priority>[!-~]+)`,
		`(?P<request_creation_time>[!-~]+)`,
		`"(?P<actions_executed>[ -~]+)"`,
		`"(?P<redirect_url>[ -~]+)"`,
		`"(?P<error_reason>[ -~]+)"`,
		`(?P<target_port_list>[ -~]+)`,
		`(?P<target_status_code_list>[ -~]+)`,
		`(?P<classification>[ -~]+)`,
		`(?P<classification_reason>[ -~]+)`,
	}
	return []*regexp.Regexp{
		regexp.MustCompile(strings.Join(basePattern, sep)),
	}
}

func generateNLBPatterns() []*regexp.Regexp {
	sep := " "
	basePattern := []string{
		`^(?P<type>[!-~]+)`,
		`(?P<version>[!-~]+)`,
		`(?P<time>[!-~]+)`,
		`(?P<elb>[!-~]+)`,
		`(?P<listener>[!-~]+)`,
		`(?P<client_port>[!-~]+)`,
		`(?P<destination_port>[!-~]+)`,
		`(?P<connection_time>[\d\-.]+)`,
		`(?P<tls_handshake_time>[\d\-.]+)`,
		`(?P<received_bytes>[!-~]+)`,
		`(?P<sent_bytes>[!-~]+)`,
		`(?P<incoming_tls_alert>[!-~]+)`,
		`(?P<chosen_cert_arn>[!-~]+)`,
		`(?P<chosen_cert_serial>[ -~]+)`,
		`(?P<tls_cipher>\S+)`,
		`(?P<tls_protocol_version>[!-~]+)`,
		`(?P<tls_named_group>[!-~]+)`,
		`(?P<domain_name>[!-~]+)`,
		`(?P<alpn_fe_protocol>[!-~]+)`,
		`(?P<alpn_be_protocol>[!-~]+)`,
		`(?P<alpn_client_preference_list>[ -~]+)`,
		`(?P<tls_connection_creation_time>[!-~]+)`,
	}
	return []*regexp.Regexp{
		regexp.MustCompile(strings.Join(basePattern, sep)),
	}
}

func generateCLBPatterns() []*regexp.Regexp {
	sep := " "
	basePattern := []string{
		`^(?P<time>[!-~]+)`,
		`(?P<elb>[!-~]+)`,
		`(?P<client_port>[!-~]+)`,
		`(?P<backend_port>[!-~]+)`,
		`(?P<request_processing_time>[\d\-.]+)`,
		`(?P<backend_processing_time>[\d\-.]+)`,
		`(?P<response_processing_time>[\d\-.]+)`,
		`(?P<elb_status_code>\d{1,3}|-)`,
		`(?P<backend_status_code>\d{1,3}|-)`,
		`(?P<received_bytes>[\d\-.]+)`,
		`(?P<sent_bytes>[\d\-.]+)`,
		`"(?P<request>[ -~]+)"`,
	}
	additions := [][]string{
		{
			`"(?P<user_agent>[^\"]*)"`,
			`(?P<ssl_cipher>[!-~]+)`,
			`(?P<ssl_protocol>[!-~]+)`,
		},
		{}, // for basePattern
	}
	patterns := make([]*regexp.Regexp, len(additions))
	for i, addition := range additions {
		patterns[i] = regexp.MustCompile(strings.Join(append(basePattern, addition...), sep))
	}
	return patterns
}
