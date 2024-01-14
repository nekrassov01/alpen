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
			name:    "clf+input",
			args:    []string{Name, "clf", "-i", `dummy`},
			wantErr: false,
		},
		{
			name:    "clf+file",
			args:    []string{Name, "clf", "-f", filepath.Join("testdata", "log", "sample_clf.log")},
			wantErr: false,
		},
		{
			name:    "clf+gzip",
			args:    []string{Name, "clf", "-g", filepath.Join("testdata", "gz", "sample_clf.log.gz")},
			wantErr: false,
		},
		{
			name:    "clf+zip",
			args:    []string{Name, "clf", "-z", filepath.Join("testdata", "zip", "sample_clf.log.zip")},
			wantErr: false,
		},
		{
			name:    "clfv+input",
			args:    []string{Name, "clfv", "-i", `dummy`},
			wantErr: false,
		},
		{
			name:    "clfv+file",
			args:    []string{Name, "clfv", "-f", filepath.Join("testdata", "log", "sample_clfv.log")},
			wantErr: false,
		},
		{
			name:    "clfv+gzip",
			args:    []string{Name, "clfv", "-g", filepath.Join("testdata", "gz", "sample_clfv.log.gz")},
			wantErr: false,
		},
		{
			name:    "clfv+zip",
			args:    []string{Name, "clfv", "-z", filepath.Join("testdata", "zip", "sample_clfv.log.zip")},
			wantErr: false,
		},
		{
			name:    "s3+input",
			args:    []string{Name, "s3", "-i", `dummy`},
			wantErr: false,
		},
		{
			name:    "s3+file",
			args:    []string{Name, "s3", "-f", filepath.Join("testdata", "log", "sample_s3.log")},
			wantErr: false,
		},
		{
			name:    "s3+gzip",
			args:    []string{Name, "s3", "-g", filepath.Join("testdata", "gz", "sample_s3.log.gz")},
			wantErr: false,
		},
		{
			name:    "s3+zip",
			args:    []string{Name, "s3", "-z", filepath.Join("testdata", "zip", "sample_s3.log.zip")},
			wantErr: false,
		},
		{
			name:    "cf+input",
			args:    []string{Name, "cf", "-i", `dummy`},
			wantErr: false,
		},
		{
			name:    "cf+file",
			args:    []string{Name, "cf", "-f", filepath.Join("testdata", "log", "sample_cf.log")},
			wantErr: false,
		},
		{
			name:    "cf+gzip",
			args:    []string{Name, "cf", "-g", filepath.Join("testdata", "gz", "sample_cf.log.gz")},
			wantErr: false,
		},
		{
			name:    "cf+zip",
			args:    []string{Name, "cf", "-z", filepath.Join("testdata", "zip", "sample_cf.log.zip")},
			wantErr: false,
		},
		{
			name:    "alb+input",
			args:    []string{Name, "alb", "-i", `dummy`},
			wantErr: false,
		},
		{
			name:    "alb+file",
			args:    []string{Name, "alb", "-f", filepath.Join("testdata", "log", "sample_alb.log")},
			wantErr: false,
		},
		{
			name:    "alb+gzip",
			args:    []string{Name, "alb", "-g", filepath.Join("testdata", "gz", "sample_alb.log.gz")},
			wantErr: false,
		},
		{
			name:    "alb+zip",
			args:    []string{Name, "alb", "-z", filepath.Join("testdata", "zip", "sample_alb.log.zip")},
			wantErr: false,
		},
		{
			name:    "nlb+input",
			args:    []string{Name, "nlb", "-i", `dummy`},
			wantErr: false,
		},
		{
			name:    "nlb+file",
			args:    []string{Name, "nlb", "-f", filepath.Join("testdata", "log", "sample_nlb.log")},
			wantErr: false,
		},
		{
			name:    "nlb+gzip",
			args:    []string{Name, "nlb", "-g", filepath.Join("testdata", "gz", "sample_nlb.log.gz")},
			wantErr: false,
		},
		{
			name:    "nlb+zip",
			args:    []string{Name, "nlb", "-z", filepath.Join("testdata", "zip", "sample_nlb.log.zip")},
			wantErr: false,
		},
		{
			name:    "clb+input",
			args:    []string{Name, "clb", "-i", `dummy`},
			wantErr: false,
		},
		{
			name:    "clb+file",
			args:    []string{Name, "clb", "-f", filepath.Join("testdata", "log", "sample_clb.log")},
			wantErr: false,
		},
		{
			name:    "clb+gzip",
			args:    []string{Name, "clb", "-g", filepath.Join("testdata", "gz", "sample_clb.log.gz")},
			wantErr: false,
		},
		{
			name:    "clb+zip",
			args:    []string{Name, "clb", "-z", filepath.Join("testdata", "zip", "sample_clb.log.zip")},
			wantErr: false,
		},
		{
			name:    "ltsv+input",
			args:    []string{Name, "ltsv", "-i", `dummy`},
			wantErr: false,
		},
		{
			name:    "ltsv+file",
			args:    []string{Name, "ltsv", "-f", filepath.Join("testdata", "log", "sample_ltsv.log")},
			wantErr: false,
		},
		{
			name:    "ltsv+gzip",
			args:    []string{Name, "ltsv", "-g", filepath.Join("testdata", "gz", "sample_ltsv.log.gz")},
			wantErr: false,
		},
		{
			name:    "ltsv+zip",
			args:    []string{Name, "ltsv", "-z", filepath.Join("testdata", "zip", "sample_ltsv.log.zip")},
			wantErr: false,
		},
		{
			name:    "metadata",
			args:    []string{Name, "clf", "-f", filepath.Join("testdata", "log", "sample_clf.log"), "-m"},
			wantErr: false,
		},
		{
			name:    "skip",
			args:    []string{Name, "clf", "-f", filepath.Join("testdata", "log", "sample_clf.log"), "-s", "1,2,3"},
			wantErr: false,
		},
		{
			name:    "skip out of range",
			args:    []string{Name, "clf", "-f", filepath.Join("testdata", "log", "sample_clf.log"), "-s", "65535"},
			wantErr: false,
		},
		{
			name:    "line_number",
			args:    []string{Name, "clf", "-f", filepath.Join("testdata", "log", "sample_clf.log"), "-l"},
			wantErr: false,
		},
		{
			name:    "glob_pattern",
			args:    []string{Name, "clf", "-z", filepath.Join("testdata", "zip", "sample_clf.log.zip"), "-G", "*"},
			wantErr: false,
		},
		{
			name:    "glob_pattern error",
			args:    []string{Name, "clf", "-z", filepath.Join("testdata", "zip", "sample_clf.log.zip"), "-G", "["},
			wantErr: true,
		},
		{
			name:    "out+json",
			args:    []string{Name, "clf", "-f", filepath.Join("testdata", "log", "sample_clf.log"), "-o", "json"},
			wantErr: false,
		},
		{
			name:    "out+pretty-json",
			args:    []string{Name, "clf", "-f", filepath.Join("testdata", "log", "sample_clf.log"), "-o", "pretty-json"},
			wantErr: false,
		},
		{
			name:    "out+text",
			args:    []string{Name, "clf", "-f", filepath.Join("testdata", "log", "sample_clf.log"), "-o", "text"},
			wantErr: false,
		},
		{
			name:    "out+ltsv",
			args:    []string{Name, "clf", "-f", filepath.Join("testdata", "log", "sample_clf.log"), "-o", "ltsv"},
			wantErr: false,
		},
		{
			name:    "out+tsv",
			args:    []string{Name, "clf", "-f", filepath.Join("testdata", "log", "sample_clf.log"), "-o", "tsv"},
			wantErr: false,
		},
		{
			name:    "out+tsv+header",
			args:    []string{Name, "clf", "-f", filepath.Join("testdata", "log", "sample_clf.log"), "-o", "tsv", "-H"},
			wantErr: false,
		},
		{
			name:    "invalid_header",
			args:    []string{Name, "clf", "-f", filepath.Join("testdata", "log", "sample_clf.log"), "-o", "ltsv", "-H"},
			wantErr: true,
		},
		{
			name:    "multiple_input",
			args:    []string{Name, "clf", "-i", "dummy", "-f", filepath.Join("testdata", "log", "sample_clf.log")},
			wantErr: true,
		},
		{
			name:    "out+unknown_format",
			args:    []string{Name, "clf", "-f", filepath.Join("testdata", "log", "sample_clf.log"), "-o", "dummy"},
			wantErr: true,
		},
		{
			name:    "completion bash",
			args:    []string{Name, "completion", "bash"},
			wantErr: false,
		},
		{
			name:    "completion zsh",
			args:    []string{Name, "completion", "zsh"},
			wantErr: false,
		},
		{
			name:    "completion pwsh",
			args:    []string{Name, "completion", "pwsh"},
			wantErr: false,
		},
		{
			name:    "completion unsupported",
			args:    []string{Name, "completion", "fish"},
			wantErr: true,
		},
		{
			name:    "unknown flag provided",
			args:    []string{Name, "-1"},
			wantErr: true,
		},
		{
			name:    "no flag provided",
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
