package main

import (
	"path/filepath"
	"testing"
)

func Test_cli(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "clf+stdin",
			args:    []string{"echo", "aaa", "|", Name, "clf"},
			wantErr: false,
		},
		{
			name:    "clf+default",
			args:    []string{Name, "clf", filepath.Join("testdata", "log", "sample_clf.log")},
			wantErr: false,
		},
		{
			name:    "clf+gz",
			args:    []string{Name, "clf", "-i", "gz", filepath.Join("testdata", "gz", "sample_clf.log.gz")},
			wantErr: false,
		},
		{
			name:    "clf+zip",
			args:    []string{Name, "clf", "-i", "zip", filepath.Join("testdata", "zip", "sample_clf.log.zip")},
			wantErr: false,
		},
		{
			name:    "clfv+default",
			args:    []string{Name, "clfv", filepath.Join("testdata", "log", "sample_clfv.log")},
			wantErr: false,
		},
		{
			name:    "clfv+gzip",
			args:    []string{Name, "clfv", "-i", "gz", filepath.Join("testdata", "gz", "sample_clfv.log.gz")},
			wantErr: false,
		},
		{
			name:    "clfv+zip",
			args:    []string{Name, "clfv", "-i", "zip", filepath.Join("testdata", "zip", "sample_clfv.log.zip")},
			wantErr: false,
		},
		{
			name:    "s3+default",
			args:    []string{Name, "s3", filepath.Join("testdata", "log", "sample_s3.log")},
			wantErr: false,
		},
		{
			name:    "s3+gzip",
			args:    []string{Name, "s3", "-i", "gz", filepath.Join("testdata", "gz", "sample_s3.log.gz")},
			wantErr: false,
		},
		{
			name:    "s3+zip",
			args:    []string{Name, "s3", "-i", "zip", filepath.Join("testdata", "zip", "sample_s3.log.zip")},
			wantErr: false,
		},
		{
			name:    "cf+default",
			args:    []string{Name, "cf", filepath.Join("testdata", "log", "sample_cf.log")},
			wantErr: false,
		},
		{
			name:    "cf+gzip",
			args:    []string{Name, "cf", "-i", "gz", filepath.Join("testdata", "gz", "sample_cf.log.gz")},
			wantErr: false,
		},
		{
			name:    "cf+zip",
			args:    []string{Name, "cf", "-i", "zip", filepath.Join("testdata", "zip", "sample_cf.log.zip")},
			wantErr: false,
		},
		{
			name:    "alb+default",
			args:    []string{Name, "alb", filepath.Join("testdata", "log", "sample_alb.log")},
			wantErr: false,
		},
		{
			name:    "alb+gzip",
			args:    []string{Name, "alb", "-i", "gz", filepath.Join("testdata", "gz", "sample_alb.log.gz")},
			wantErr: false,
		},
		{
			name:    "alb+zip",
			args:    []string{Name, "alb", "-i", "zip", filepath.Join("testdata", "zip", "sample_alb.log.zip")},
			wantErr: false,
		},
		{
			name:    "nlb+default",
			args:    []string{Name, "nlb", filepath.Join("testdata", "log", "sample_nlb.log")},
			wantErr: false,
		},
		{
			name:    "nlb+gzip",
			args:    []string{Name, "nlb", "-i", "gz", filepath.Join("testdata", "gz", "sample_nlb.log.gz")},
			wantErr: false,
		},
		{
			name:    "nlb+zip",
			args:    []string{Name, "nlb", "-i", "zip", filepath.Join("testdata", "zip", "sample_nlb.log.zip")},
			wantErr: false,
		},
		{
			name:    "clb+default",
			args:    []string{Name, "clb", filepath.Join("testdata", "log", "sample_clb.log")},
			wantErr: false,
		},
		{
			name:    "clb+gzip",
			args:    []string{Name, "clb", "-i", "gz", filepath.Join("testdata", "gz", "sample_clb.log.gz")},
			wantErr: false,
		},
		{
			name:    "clb+zip",
			args:    []string{Name, "clb", "-i", "zip", filepath.Join("testdata", "zip", "sample_clb.log.zip")},
			wantErr: false,
		},
		{
			name:    "ltsv+default",
			args:    []string{Name, "ltsv", filepath.Join("testdata", "log", "sample_ltsv.log")},
			wantErr: false,
		},
		{
			name:    "ltsv+gzip",
			args:    []string{Name, "ltsv", "-i", "gz", filepath.Join("testdata", "gz", "sample_ltsv.log.gz")},
			wantErr: false,
		},
		{
			name:    "ltsv+zip",
			args:    []string{Name, "ltsv", "-i", "zip", filepath.Join("testdata", "zip", "sample_ltsv.log.zip")},
			wantErr: false,
		},
		{
			name:    "result",
			args:    []string{Name, "clf", "-r", filepath.Join("testdata", "log", "sample_clf.log")},
			wantErr: false,
		},
		{
			name:    "labels",
			args:    []string{Name, "clf", "-l", "remote_host,method,request_uri,protocol", filepath.Join("testdata", "log", "sample_clf.log")},
			wantErr: false,
		},
		{
			name:    "filters",
			args:    []string{Name, "clf", "-f", "size < 100,method == GET, remote_host =~ ^192.168", filepath.Join("testdata", "log", "sample_clf.log")},
			wantErr: false,
		},
		{
			name:    "skip",
			args:    []string{Name, "clf", "-s", "1,2,3", filepath.Join("testdata", "log", "sample_clf.log")},
			wantErr: false,
		},
		{
			name:    "skip_out_of_range",
			args:    []string{Name, "clf", "-s", "65535", filepath.Join("testdata", "log", "sample_clf.log")},
			wantErr: false,
		},
		{
			name:    "line_number",
			args:    []string{Name, "clf", "-n", filepath.Join("testdata", "log", "sample_clf.log")},
			wantErr: false,
		},
		{
			name:    "in+default",
			args:    []string{Name, "clf", "-i", "default", filepath.Join("testdata", "log", "sample_clf.log")},
			wantErr: false,
		},
		{
			name:    "in+gz",
			args:    []string{Name, "clf", "-i", "gz", filepath.Join("testdata", "gz", "sample_clf.log.gz")},
			wantErr: false,
		},
		{
			name:    "in+zip",
			args:    []string{Name, "clf", "-i", "zip", "-g", "*", filepath.Join("testdata", "zip", "sample_clf.log.zip")},
			wantErr: false,
		},
		{
			name:    "in+zip+glob_pattern_error",
			args:    []string{Name, "clf", "-i", "zip", "-g", "[", filepath.Join("testdata", "zip", "sample_clf.log.zip")},
			wantErr: true,
		},
		{
			name:    "in+invalid_type",
			args:    []string{Name, "clf", "-i", "", filepath.Join("testdata", "zip", "sample_clf.log.zip")},
			wantErr: true,
		},
		{
			name:    "out+json",
			args:    []string{Name, "clf", "-o", "json", filepath.Join("testdata", "log", "sample_clf.log")},
			wantErr: false,
		},
		{
			name:    "out+pretty-json",
			args:    []string{Name, "clf", "-o", "pretty-json", filepath.Join("testdata", "log", "sample_clf.log")},
			wantErr: false,
		},
		{
			name:    "out+text",
			args:    []string{Name, "clf", "-o", "text", filepath.Join("testdata", "log", "sample_clf.log")},
			wantErr: false,
		},
		{
			name:    "out+ltsv",
			args:    []string{Name, "clf", "-o", "ltsv", filepath.Join("testdata", "log", "sample_clf.log")},
			wantErr: false,
		},
		{
			name:    "out+tsv",
			args:    []string{Name, "clf", "-o", "tsv", filepath.Join("testdata", "log", "sample_clf.log")},
			wantErr: false,
		},
		{
			name:    "out+invalid_type",
			args:    []string{Name, "clf", "-o", "", filepath.Join("testdata", "zip", "sample_clf.log.zip")},
			wantErr: true,
		},
		{
			name:    "no_zip+glob",
			args:    []string{Name, "clf", "-i", "gz", "-g", "*", filepath.Join("testdata", "zip", "sample_clf.log.zip")},
			wantErr: true,
		},
		{
			name:    "completion_bash",
			args:    []string{Name, "completion", "bash"},
			wantErr: false,
		},
		{
			name:    "completion_zsh",
			args:    []string{Name, "completion", "zsh"},
			wantErr: false,
		},
		{
			name:    "completion_pwsh",
			args:    []string{Name, "completion", "pwsh"},
			wantErr: false,
		},
		{
			name:    "completion_unsupported",
			args:    []string{Name, "completion", "fish"},
			wantErr: true,
		},
		{
			name:    "unknown_flag_provided",
			args:    []string{Name, "-1"},
			wantErr: true,
		},
		{
			name:    "no_flag_provided",
			args:    []string{Name},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := newApp().cli.Run(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
