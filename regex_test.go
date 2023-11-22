package main

import (
	"reflect"
	"regexp"
	"testing"
)

func Test_generateApacheCLFPatterns(t *testing.T) {
	tests := []struct {
		name string
		want []*regexp.Regexp
	}{
		{
			name: "basic",
			want: []*regexp.Regexp{
				regexp.MustCompile(`^(?P<remote_host>\S+) (?P<remote_logname>\S+) (?P<remote_user>[\S ]+) (?P<datetime>\[[^\]]+\]) \"(?P<method>[A-Z]+) (?P<request_uri>[^ \"]+) (?P<protocol>HTTP/[0-9.]+)\" (?P<status>[0-9]{3}) (?P<size>[0-9]+|-) "(?P<referer>[^\"]*)" "(?P<user_agent>[^\"]*)"`),
				regexp.MustCompile(`^(?P<remote_host>\S+) (?P<remote_logname>\S+) (?P<remote_user>[\S ]+) (?P<datetime>\[[^\]]+\]) \"(?P<method>[A-Z]+) (?P<request_uri>[^ \"]+) (?P<protocol>HTTP/[0-9.]+)\" (?P<status>[0-9]{3}) (?P<size>[0-9]+|-)`),
				regexp.MustCompile(`^(?P<remote_host>\S+)	(?P<remote_logname>\S+)	(?P<remote_user>[\S ]+)	(?P<datetime>\[[^\]]+\])	\"(?P<method>[A-Z]+) (?P<request_uri>[^ \"]+) (?P<protocol>HTTP/[0-9.]+)\"	(?P<status>[0-9]{3})	(?P<size>[0-9]+|-)	"(?P<referer>[^\"]*)"	"(?P<user_agent>[^\"]*)"`),
				regexp.MustCompile(`^(?P<remote_host>\S+)	(?P<remote_logname>\S+)	(?P<remote_user>[\S ]+)	(?P<datetime>\[[^\]]+\])	\"(?P<method>[A-Z]+) (?P<request_uri>[^ \"]+) (?P<protocol>HTTP/[0-9.]+)\"	(?P<status>[0-9]{3})	(?P<size>[0-9]+|-)`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateApacheCLFPatterns(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func Test_generateApacheCLFWithVHostPatterns(t *testing.T) {
	tests := []struct {
		name string
		want []*regexp.Regexp
	}{
		{
			name: "basic",
			want: []*regexp.Regexp{
				regexp.MustCompile(`^(?P<virtual_host>\S+) (?P<remote_host>\S+) (?P<remote_logname>\S+) (?P<remote_user>[\S ]+) (?P<datetime>\[[^\]]+\]) \"(?P<method>[A-Z]+) (?P<request_uri>[^ \"]+) (?P<protocol>HTTP/[0-9.]+)\" (?P<status>[0-9]{3}) (?P<size>[0-9]+|-) "(?P<referer>[^\"]*)" "(?P<user_agent>[^\"]*)"`),
				regexp.MustCompile(`^(?P<virtual_host>\S+) (?P<remote_host>\S+) (?P<remote_logname>\S+) (?P<remote_user>[\S ]+) (?P<datetime>\[[^\]]+\]) \"(?P<method>[A-Z]+) (?P<request_uri>[^ \"]+) (?P<protocol>HTTP/[0-9.]+)\" (?P<status>[0-9]{3}) (?P<size>[0-9]+|-)`),
				regexp.MustCompile(`^(?P<virtual_host>\S+)	(?P<remote_host>\S+)	(?P<remote_logname>\S+)	(?P<remote_user>[\S ]+)	(?P<datetime>\[[^\]]+\])	\"(?P<method>[A-Z]+) (?P<request_uri>[^ \"]+) (?P<protocol>HTTP/[0-9.]+)\"	(?P<status>[0-9]{3})	(?P<size>[0-9]+|-)	"(?P<referer>[^\"]*)"	"(?P<user_agent>[^\"]*)"`),
				regexp.MustCompile(`^(?P<virtual_host>\S+)	(?P<remote_host>\S+)	(?P<remote_logname>\S+)	(?P<remote_user>[\S ]+)	(?P<datetime>\[[^\]]+\])	\"(?P<method>[A-Z]+) (?P<request_uri>[^ \"]+) (?P<protocol>HTTP/[0-9.]+)\"	(?P<status>[0-9]{3})	(?P<size>[0-9]+|-)`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateApacheCLFWithVHostPatterns(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func Test_generateS3Patterns(t *testing.T) {
	tests := []struct {
		name string
		want []*regexp.Regexp
	}{
		{
			name: "basic",
			want: []*regexp.Regexp{
				regexp.MustCompile(`^(?P<bucket_owner>[!-~]+) (?P<bucket>[!-~]+) (?P<time>\[[ -~]+ [0-9+]+\]) (?P<remote_ip>[!-~]+) (?P<requester>[!-~]+) (?P<request_id>[!-~]+) (?P<operation>[!-~]+) (?P<key>[!-~]+) "(?P<request_uri>[ -~]+)" (?P<http_status>\d{1,3}) (?P<error_code>[!-~]+) (?P<bytes_sent>[\d\-.]+) (?P<object_size>[\d\-.]+) (?P<total_time>[\d\-.]+) (?P<turn_around_time>[\d\-.]+) "(?P<referer>[^\"]*)" "(?P<user_agent>[^\"]*)" (?P<version_id>[!-~]+) (?P<host_id>[!-~]+) (?P<signature_version>[!-~]+) (?P<cipher_suite>[!-~]+) (?P<authentication_type>[!-~]+) (?P<host_header>[!-~]+) (?P<tls_version>[!-~]+) (?P<access_point_arn>[!-~]+) (?P<acl_required>[!-~]+)`),
				regexp.MustCompile(`^(?P<bucket_owner>[!-~]+) (?P<bucket>[!-~]+) (?P<time>\[[ -~]+ [0-9+]+\]) (?P<remote_ip>[!-~]+) (?P<requester>[!-~]+) (?P<request_id>[!-~]+) (?P<operation>[!-~]+) (?P<key>[!-~]+) "(?P<request_uri>[ -~]+)" (?P<http_status>\d{1,3}) (?P<error_code>[!-~]+) (?P<bytes_sent>[\d\-.]+) (?P<object_size>[\d\-.]+) (?P<total_time>[\d\-.]+) (?P<turn_around_time>[\d\-.]+) "(?P<referer>[^\"]*)" "(?P<user_agent>[^\"]*)" (?P<version_id>[!-~]+) (?P<host_id>[!-~]+) (?P<signature_version>[!-~]+) (?P<cipher_suite>[!-~]+) (?P<authentication_type>[!-~]+) (?P<host_header>[!-~]+) (?P<tls_version>[!-~]+) (?P<access_point_arn>[!-~]+)`),
				regexp.MustCompile(`^(?P<bucket_owner>[!-~]+) (?P<bucket>[!-~]+) (?P<time>\[[ -~]+ [0-9+]+\]) (?P<remote_ip>[!-~]+) (?P<requester>[!-~]+) (?P<request_id>[!-~]+) (?P<operation>[!-~]+) (?P<key>[!-~]+) "(?P<request_uri>[ -~]+)" (?P<http_status>\d{1,3}) (?P<error_code>[!-~]+) (?P<bytes_sent>[\d\-.]+) (?P<object_size>[\d\-.]+) (?P<total_time>[\d\-.]+) (?P<turn_around_time>[\d\-.]+) "(?P<referer>[^\"]*)" "(?P<user_agent>[^\"]*)" (?P<version_id>[!-~]+) (?P<host_id>[!-~]+) (?P<signature_version>[!-~]+) (?P<cipher_suite>[!-~]+) (?P<authentication_type>[!-~]+) (?P<host_header>[!-~]+) (?P<tls_version>[!-~]+)`),
				regexp.MustCompile(`^(?P<bucket_owner>[!-~]+) (?P<bucket>[!-~]+) (?P<time>\[[ -~]+ [0-9+]+\]) (?P<remote_ip>[!-~]+) (?P<requester>[!-~]+) (?P<request_id>[!-~]+) (?P<operation>[!-~]+) (?P<key>[!-~]+) "(?P<request_uri>[ -~]+)" (?P<http_status>\d{1,3}) (?P<error_code>[!-~]+) (?P<bytes_sent>[\d\-.]+) (?P<object_size>[\d\-.]+) (?P<total_time>[\d\-.]+) (?P<turn_around_time>[\d\-.]+) "(?P<referer>[^\"]*)" "(?P<user_agent>[^\"]*)" (?P<version_id>[!-~]+) (?P<host_id>[!-~]+) (?P<signature_version>[!-~]+) (?P<cipher_suite>[!-~]+) (?P<authentication_type>[!-~]+) (?P<host_header>[!-~]+)`),
				regexp.MustCompile(`^(?P<bucket_owner>[!-~]+) (?P<bucket>[!-~]+) (?P<time>\[[ -~]+ [0-9+]+\]) (?P<remote_ip>[!-~]+) (?P<requester>[!-~]+) (?P<request_id>[!-~]+) (?P<operation>[!-~]+) (?P<key>[!-~]+) "(?P<request_uri>[ -~]+)" (?P<http_status>\d{1,3}) (?P<error_code>[!-~]+) (?P<bytes_sent>[\d\-.]+) (?P<object_size>[\d\-.]+) (?P<total_time>[\d\-.]+) (?P<turn_around_time>[\d\-.]+) "(?P<referer>[^\"]*)" "(?P<user_agent>[^\"]*)" (?P<version_id>[!-~]+)`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateS3Patterns(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func Test_generateCFPatterns(t *testing.T) {
	tests := []struct {
		name string
		want []*regexp.Regexp
	}{
		{
			name: "basic",
			want: []*regexp.Regexp{
				regexp.MustCompile(`^(?P<date>[\d\-.:]+)	(?P<time>[\d\-.:]+)	(?P<x_edge_location>[ -~]+)	(?P<sc_bytes>[\d\-.]+)	(?P<c_ip>[ -~]+)	(?P<cs_method>[ -~]+)	(?P<cs_host>[ -~]+)	(?P<cs_uri_stem>[ -~]+)	(?P<sc_status>\d{1,3}|-)	(?P<cs_referer>[^\"]*)	(?P<cs_user_agent>[^\"]*)	(?P<cs_uri_query>[ -~]+)	(?P<cs_cookie>\S+)	(?P<x_edge_result_type>[ -~]+)	(?P<x_edge_request_id>[ -~]+)	(?P<x_host_header>[ -~]+)	(?P<cs_protocol>[ -~]+)	(?P<cs_bytes>[\d\-.]+)	(?P<time_taken>[\d\-.]+)	(?P<x_forwarded_for>[ -~]+)	(?P<ssl_protocol>[ -~]+)	(?P<ssl_cipher>[ -~]+)	(?P<x_edge_response_result_type>[ -~]+)	(?P<cs_protocol_version>[ -~]+)	(?P<fle_status>[ -~]+)	(?P<fle_encrypted_fields>\S+)	(?P<c_port>[\d\-.]+)	(?P<time_to_first_byte>[\d\-.]+)	(?P<x_edge_detailed_result_type>[ -~]+)	(?P<sc_content_type>[ -~]+)	(?P<sc_content_len>[\d\-.]+)	(?P<sc_range_start>[\d\-.]+)	(?P<sc_range_end>[\d\-.]+)`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateCFPatterns(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func Test_generateALBPatterns(t *testing.T) {
	tests := []struct {
		name string
		want []*regexp.Regexp
	}{
		{
			name: "basic",
			want: []*regexp.Regexp{
				regexp.MustCompile(`^(?P<type>[!-~]+) (?P<time>[!-~]+) (?P<elb>[!-~]+) (?P<client_port>[!-~]+) (?P<target_port>[!-~]+) (?P<request_processing_time>[\d\-.]+) (?P<target_processing_time>[\d\-.]+) (?P<response_processing_time>[\d\-.]+) (?P<elb_status_code>\d{1,3}|-) (?P<target_status_code>\d{1,3}|-) (?P<received_bytes>[\d\-.]+) (?P<sent_bytes>[\d\-.]+) "(?P<request>[ -~]+)" "(?P<user_agent>[^\"]*)" (?P<ssl_cipher>[!-~]+) (?P<ssl_protocol>[!-~]+) (?P<target_group_arn>[!-~]+) "(?P<trace_id>[ -~]+)" "(?P<domain_name>[ -~]+)" "(?P<chosen_cert_arn>[ -~]+)" (?P<matched_rule_priority>[!-~]+) (?P<request_creation_time>[!-~]+) "(?P<actions_executed>[ -~]+)" "(?P<redirect_url>[ -~]+)" "(?P<error_reason>[ -~]+)" (?P<target_port_list>[ -~]+) (?P<target_status_code_list>[ -~]+) (?P<classification>[ -~]+) (?P<classification_reason>[ -~]+)`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateALBPatterns(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func Test_generateNLBPatterns(t *testing.T) {
	tests := []struct {
		name string
		want []*regexp.Regexp
	}{
		{
			name: "basic",
			want: []*regexp.Regexp{
				regexp.MustCompile(`^(?P<type>[!-~]+) (?P<version>[!-~]+) (?P<time>[!-~]+) (?P<elb>[!-~]+) (?P<listener>[!-~]+) (?P<client_port>[!-~]+) (?P<destination_port>[!-~]+) (?P<connection_time>[\d\-.]+) (?P<tls_handshake_time>[\d\-.]+) (?P<received_bytes>[!-~]+) (?P<sent_bytes>[!-~]+) (?P<incoming_tls_alert>[!-~]+) (?P<chosen_cert_arn>[!-~]+) (?P<chosen_cert_serial>[ -~]+) (?P<tls_cipher>\S+) (?P<tls_protocol_version>[!-~]+) (?P<tls_named_group>[!-~]+) (?P<domain_name>[!-~]+) (?P<alpn_fe_protocol>[!-~]+) (?P<alpn_be_protocol>[!-~]+) (?P<alpn_client_preference_list>[ -~]+) (?P<tls_connection_creation_time>[!-~]+)`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateNLBPatterns(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func Test_generateCLBPatterns(t *testing.T) {
	tests := []struct {
		name string
		want []*regexp.Regexp
	}{
		{
			name: "basic",
			want: []*regexp.Regexp{
				regexp.MustCompile(`^(?P<time>[!-~]+) (?P<elb>[!-~]+) (?P<client_port>[!-~]+) (?P<backend_port>[!-~]+) (?P<request_processing_time>[\d\-.]+) (?P<backend_processing_time>[\d\-.]+) (?P<response_processing_time>[\d\-.]+) (?P<elb_status_code>\d{1,3}|-) (?P<backend_status_code>\d{1,3}|-) (?P<received_bytes>[\d\-.]+) (?P<sent_bytes>[\d\-.]+) "(?P<request>[ -~]+)" "(?P<user_agent>[^\"]*)" (?P<ssl_cipher>[!-~]+) (?P<ssl_protocol>[!-~]+)`),
				regexp.MustCompile(`^(?P<time>[!-~]+) (?P<elb>[!-~]+) (?P<client_port>[!-~]+) (?P<backend_port>[!-~]+) (?P<request_processing_time>[\d\-.]+) (?P<backend_processing_time>[\d\-.]+) (?P<response_processing_time>[\d\-.]+) (?P<elb_status_code>\d{1,3}|-) (?P<backend_status_code>\d{1,3}|-) (?P<received_bytes>[\d\-.]+) (?P<sent_bytes>[\d\-.]+) "(?P<request>[ -~]+)"`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateCLBPatterns(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
