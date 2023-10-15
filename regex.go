package main

import (
	"regexp"
	"strings"
)

func generateS3Patterns() (patterns []*regexp.Regexp) {
	sep := " "
	basePattern := []string{
		`^([!-~]+)`,            // bucket_owner
		`([!-~]+)`,             // bucket
		`(\[[ -~]+ [0-9+]+\])`, // time
		`([!-~]+)`,             // remote_ip
		`([!-~]+)`,             // requester
		`([!-~]+)`,             // request_id
		`([!-~]+)`,             // operation
		`([!-~]+)`,             // key
		`"([ -~]+)"`,           // request_uri
		`(\d{1,3})`,            // http_status
		`([!-~]+)`,             // error_code
		`([\d\-.]+)`,           // bytes_sent
		`([\d\-.]+)`,           // object_size
		`([\d\-.]+)`,           // total_time
		`([\d\-.]+)`,           // turn_around_time
		`"([ -~]+)"`,           // referer
		`"([ -~]+)"`,           // user_agent
		`([!-~]+)`,             // version_id

	}
	additions := [][]string{
		{
			`([!-~]+)`, // host_id
			`([!-~]+)`, // signature_version
			`([!-~]+)`, // cipher_suite
			`([!-~]+)`, // authentication_type
			`([!-~]+)`, // host_header
			`([!-~]+)`, // tls_version
			`([!-~]+)`, // access_point_arn
			`([!-~]+)`, // acl_required
		},
		{
			`([!-~]+)`, // host_id
			`([!-~]+)`, // signature_version
			`([!-~]+)`, // cipher_suite
			`([!-~]+)`, // authentication_type
			`([!-~]+)`, // host_header
			`([!-~]+)`, // tls_version
			`([!-~]+)`, // access_point_arn
		},

		{
			`([!-~]+)`, // host_id
			`([!-~]+)`, // signature_version
			`([!-~]+)`, // cipher_suite
			`([!-~]+)`, // authentication_type
			`([!-~]+)`, // host_header
			`([!-~]+)`, // tls_version
		},
		{
			`([!-~]+)`, // host_id
			`([!-~]+)`, // signature_version
			`([!-~]+)`, // cipher_suite
			`([!-~]+)`, // authentication_type
			`([!-~]+)`, // host_header
		},
		{}, // for basePattern
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
		`^([\d\-.:]+)`, // date
		`([\d\-.:]+)`,  // time
		`([ -~]+)`,     // x_edge_location
		`([\d\-.]+)`,   // sc_bytes
		`([ -~]+)`,     // c_ip
		`([ -~]+)`,     // cs_method
		`([ -~]+)`,     // cs_host
		`([ -~]+)`,     // cs_uri_stem
		`(\d{1,3}|-)`,  // sc_status
		`([ -~]+)`,     // cs_referer
		`([ -~]+)`,     // cs_user_agent
		`([ -~]+)`,     // cs_uri_query
		`(\S+)`,        // cs_cookie
		`([ -~]+)`,     // x_edge_result_type
		`([ -~]+)`,     // x_edge_request_id
		`([ -~]+)`,     // x_host_header
		`([ -~]+)`,     // cs_protocol
		`([\d\-.]+)`,   // cs_bytes
		`([\d\-.]+)`,   // time_taken
		`([ -~]+)`,     // x_forwarded_for
		`([ -~]+)`,     // ssl_protocol
		`([ -~]+)`,     // ssl_cipher
		`([ -~]+)`,     // x_edge_response_result_type
		`([ -~]+)`,     // cs_protocol_version
		`([ -~]+)`,     // fle_status
		`(\S+)`,        // fle_encrypted_fields
		`([\d\-.]+)`,   // c_port
		`([\d\-.]+)`,   // time_to_first_byte
		`([ -~]+)`,     // x_edge_detailed_result_type
		`([ -~]+)`,     // sc_content_type
		`([\d\-.]+)`,   // sc_content_len
		`([\d\-.]+)`,   // sc_range_start
		`([\d\-.]+)`,   // sc_range_end
	}
	return []*regexp.Regexp{
		regexp.MustCompile(strings.Join(basePattern, sep)),
	}
}

func generateALBPatterns() (patterns []*regexp.Regexp) {
	sep := " "
	basePattern := []string{
		`^([!-~]+)`,   // type
		`([!-~]+)`,    // time
		`([!-~]+)`,    // elb
		`([!-~]+)`,    // client_port
		`([!-~]+)`,    // target_port
		`([\d\-.]+)`,  // request_processing_time
		`([\d\-.]+)`,  // target_processing_time
		`([\d\-.]+)`,  // response_processing_time
		`(\d{1,3}|-)`, // elb_status_code
		`(\d{1,3}|-)`, // target_status_code
		`([\d\-.]+)`,  // received_bytes
		`([\d\-.]+)`,  // sent_bytes
		`"([ -~]+)"`,  // request
		`"([ -~]+)"`,  // user_agent
		`([!-~]+)`,    // ssl_ciphe
		`([!-~]+)`,    // ssl_protocol
		`([!-~]+)`,    // target_group_arn
		`"([ -~]+)"`,  // trace_id
		`"([ -~]+)"`,  // domain_name
		`"([ -~]+)"`,  // chosen_cert_arn
		`([!-~]+)`,    // matched_rule_priority
		`([!-~]+)`,    // request_creation_time
		`"([ -~]+)"`,  // actions_executed
		`"([ -~]+)"`,  // redirect_url
		`"([ -~]+)"`,  // error_reason
		`"([ -~]+)"`,  // target_port_list
		`"([ -~]+)"`,  // target_status_code_list
		`"([ -~]+)"`,  // classification
		`"([ -~]+)"`,  // classification_reason
	}

	return []*regexp.Regexp{
		regexp.MustCompile(strings.Join(basePattern, sep)),
	}
}

func generateNLBPatterns() (patterns []*regexp.Regexp) {
	sep := " "
	basePattern := []string{
		`^([!-~]+)`,  // type
		`([!-~]+)`,   // version
		`([!-~]+)`,   // time
		`([!-~]+)`,   // elb
		`([!-~]+)`,   // listener
		`([!-~]+)`,   // client:port
		`([!-~]+)`,   // destination:port
		`([\d\-.]+)`, // connection_time
		`([\d\-.]+)`, // tls_handshake_time
		`([!-~]+)`,   // received_bytes
		`([!-~]+)`,   // sent_bytes
		`([!-~]+)`,   // incoming_tls_alert
		`([!-~]+)`,   // chosen_cert_arn
		`([ -~]+)`,   // chosen_cert_serial
		`(\S+)`,      // tls_cipher
		`([!-~]+)`,   // tls_protocol_version
		`([!-~]+)`,   // tls_named_group
		`([!-~]+)`,   // domain_name
		`([!-~]+)`,   // alpn_fe_protocol
		`([!-~]+)`,   // alpn_be_protocol
		`([ -~]+)`,   // alpn_client_preference_list
		`([!-~]+)`,   // tls_connection_creation_time
	}

	return []*regexp.Regexp{
		regexp.MustCompile(strings.Join(basePattern, sep)),
	}
}

func generateCLBPatterns() (patterns []*regexp.Regexp) {
	sep := " "
	basePattern := []string{
		`^([!-~]+)`,   // time
		`([!-~]+)`,    // elb
		`([!-~]+)`,    // client_port
		`([!-~]+)`,    // backend_port
		`([\d\-.]+)`,  // request_processing_time
		`([\d\-.]+)`,  // backend_processing_time
		`([\d\-.]+)`,  // response_processing_time
		`(\d{1,3}|-)`, // elb_status_code
		`(\d{1,3}|-)`, // backend_status_code
		`([\d\-.]+)`,  // received_bytes
		`([\d\-.]+)`,  // sent_bytes
		`"([ -~]+)"`,  // request
	}
	additions := [][]string{
		{
			`"([ -~]+)"`, // user_agent
			`([!-~]+)`,   // ssl_cipher
			`([!-~]+)`,   // ssl_protocol
		},
		{}, // for basePattern
	}
	patterns = make([]*regexp.Regexp, len(additions))
	for i, addition := range additions {
		patterns[i] = regexp.MustCompile(strings.Join(append(basePattern, addition...), sep))
	}
	return patterns
}
