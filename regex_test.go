package main

import (
	"reflect"
	"regexp"
	"testing"
)

func Test_generateS3Patterns(t *testing.T) {
	tests := []struct {
		name string
		want []*regexp.Regexp
	}{
		{
			name: "basic",
			want: []*regexp.Regexp{
				regexp.MustCompile(`^([!-~]+) ([!-~]+) (\[[ -~]+ [0-9+]+\]) ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) "([ -~]+)" (\d{1,3}) ([!-~]+) ([\d\-.]+) ([\d\-.]+) ([\d\-.]+) ([\d\-.]+) "([ -~]+)" "([ -~]+)" ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+)`),
				regexp.MustCompile(`^([!-~]+) ([!-~]+) (\[[ -~]+ [0-9+]+\]) ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) "([ -~]+)" (\d{1,3}) ([!-~]+) ([\d\-.]+) ([\d\-.]+) ([\d\-.]+) ([\d\-.]+) "([ -~]+)" "([ -~]+)" ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+)`),
				regexp.MustCompile(`^([!-~]+) ([!-~]+) (\[[ -~]+ [0-9+]+\]) ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) "([ -~]+)" (\d{1,3}) ([!-~]+) ([\d\-.]+) ([\d\-.]+) ([\d\-.]+) ([\d\-.]+) "([ -~]+)" "([ -~]+)" ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+)`),
				regexp.MustCompile(`^([!-~]+) ([!-~]+) (\[[ -~]+ [0-9+]+\]) ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) "([ -~]+)" (\d{1,3}) ([!-~]+) ([\d\-.]+) ([\d\-.]+) ([\d\-.]+) ([\d\-.]+) "([ -~]+)" "([ -~]+)" ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+)`),
				regexp.MustCompile(`^([!-~]+) ([!-~]+) (\[[ -~]+ [0-9+]+\]) ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) "([ -~]+)" (\d{1,3}) ([!-~]+) ([\d\-.]+) ([\d\-.]+) ([\d\-.]+) ([\d\-.]+) "([ -~]+)" "([ -~]+)" ([!-~]+)`),
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
				regexp.MustCompile(`^([\d\-.:]+)	([\d\-.:]+)	([ -~]+)	([\d\-.]+)	([ -~]+)	([ -~]+)	([ -~]+)	([ -~]+)	(\d{1,3}|-)	([ -~]+)	([ -~]+)	([ -~]+)	(\S+)	([ -~]+)	([ -~]+)	([ -~]+)	([ -~]+)	([\d\-.]+)	([\d\-.]+)	([ -~]+)	([ -~]+)	([ -~]+)	([ -~]+)	([ -~]+)	([ -~]+)	(\S+)	([\d\-.]+)	([\d\-.]+)	([ -~]+)	([ -~]+)	([\d\-.]+)	([\d\-.]+)	([\d\-.]+)`),
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
				regexp.MustCompile(`^([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) ([\d\-.]+) ([\d\-.]+) ([\d\-.]+) (\d{1,3}|-) (\d{1,3}|-) ([\d\-.]+) ([\d\-.]+) "([ -~]+)" "([ -~]+)" ([!-~]+) ([!-~]+) ([!-~]+) "([ -~]+)" "([ -~]+)" "([ -~]+)" ([!-~]+) ([!-~]+) "([ -~]+)" "([ -~]+)" "([ -~]+)" "([ -~]+)" "([ -~]+)" "([ -~]+)" "([ -~]+)"`),
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
				regexp.MustCompile(`^([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) ([\d\-.]+) ([\d\-.]+) ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) ([ -~]+) (\S+) ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) ([ -~]+) ([!-~]+)`),
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
				regexp.MustCompile(`^([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) ([\d\-.]+) ([\d\-.]+) ([\d\-.]+) (\d{1,3}|-) (\d{1,3}|-) ([\d\-.]+) ([\d\-.]+) "([ -~]+)" "([ -~]+)" ([!-~]+) ([!-~]+)`),
				regexp.MustCompile(`^([!-~]+) ([!-~]+) ([!-~]+) ([!-~]+) ([\d\-.]+) ([\d\-.]+) ([\d\-.]+) (\d{1,3}|-) (\d{1,3}|-) ([\d\-.]+) ([\d\-.]+) "([ -~]+)"`),
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
