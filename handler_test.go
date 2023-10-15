package main

import (
	"reflect"
	"testing"

	"github.com/nekrassov01/access-log-parser"
)

func Test_prettyJSONLineHandler(t *testing.T) {
	type args struct {
		matches []string
		fields  []string
		index   int
	}
	type want struct {
		got string
		err error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "basic",
			args: args{
				matches: []string{"", "value1", "value2"},
				fields:  []string{"field1", "field2"},
				index:   1,
			},
			want: want{
				got: `{
  "index": 1,
  "field1": "value1",
  "field2": "value2"
}`,
				err: nil,
			},
		},
		{
			name: "invalid json character",
			args: args{
				matches: []string{"", "value1", "val\"ue2"},
				fields:  []string{"field1", "field2"},
				index:   2,
			},
			want: want{
				got: `{
  "index": 2,
  "field1": "value1",
  "field2": "val\"ue2"
}`,
				err: nil,
			},
		},
		{
			name: "more matches than fields",
			args: args{
				matches: []string{"", "value1", "value2", "value3"},
				fields:  []string{"field1", "field2"},
				index:   3,
			},
			want: want{
				got: `{
  "index": 3,
  "field1": "value1",
  "field2": "value2"
}`,
				err: nil,
			},
		},
		{
			name: "more fields than matches",
			args: args{
				matches: []string{"", "value1"},
				fields:  []string{"field1", "field2"},
				index:   4,
			},
			want: want{
				got: `{
  "index": 4,
  "field1": "value1"
}`,
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := prettyJSONLineHandler(tt.args.matches, tt.args.fields, tt.args.index)
			if err != nil && err.Error() != tt.want.err.Error() {
				t.Errorf("got: %v, want: %v", err.Error(), tt.want.err.Error())
			}
			if !reflect.DeepEqual(got, tt.want.got) {
				t.Errorf("got: %v, want: %v", got, tt.want.got)
			}
		})
	}
}

func Test_prettyJSONMetadataHandler(t *testing.T) {
	type args struct {
		m *parser.Metadata
	}
	type want struct {
		got string
		err error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "basic",
			args: args{
				m: &parser.Metadata{
					Total:     5,
					Matched:   3,
					Unmatched: 1,
					Skipped:   1,
					Source:    "",
					Errors: []parser.ErrorRecord{
						{
							Index:  1,
							Record: "aaa bbb ccc",
						},
					},
				},
			},
			want: want{
				got: `{
  "total": 5,
  "matched": 3,
  "unmatched": 1,
  "skipped": 1,
  "source": "",
  "errors": [
    {
      "index": 1,
      "record": "aaa bbb ccc"
    }
  ]
}`,
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := prettyJSONMetadataHandler(tt.args.m)
			if err != nil && err.Error() != tt.want.err.Error() {
				t.Errorf("got: %v, want: %v", err.Error(), tt.want.err.Error())
			}
			if !reflect.DeepEqual(got, tt.want.got) {
				t.Errorf("got: %v, want: %v", got, tt.want.got)
			}
		})
	}
}

func Test_prettyJSON(t *testing.T) {
	type args struct {
		s string
	}
	type want struct {
		got string
		err error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "basic",
			args: args{
				s: `{"index":1,"field1":"value1","field2":"value2"}`,
			},
			want: want{
				got: `{
  "index": 1,
  "field1": "value1",
  "field2": "value2"
}`,
				err: nil,
			},
		},
		{
			name: "invalid json character",
			args: args{
				s: `{"index":2,"field1":"value1","field2":"val\"ue2"}`,
			},
			want: want{
				got: `{
  "index": 2,
  "field1": "value1",
  "field2": "val\"ue2"
}`,
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := prettyJSON(tt.args.s)
			if err != nil && err.Error() != tt.want.err.Error() {
				t.Errorf("got: %v, want: %v", err.Error(), tt.want.err.Error())
			}
			if !reflect.DeepEqual(got, tt.want.got) {
				t.Errorf("got: %v, want: %v", got, tt.want.got)
			}
		})
	}
}

func Test_textLineHandler(t *testing.T) {
	type args struct {
		matches []string
		fields  []string
		index   int
	}
	type want struct {
		got string
		err error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "basic",
			args: args{
				matches: []string{"", "value1", "value2"},
				fields:  []string{"field1", "field2"},
				index:   1,
			},
			want: want{
				got: `index=1 field1="value1" field2="value2"`,
				err: nil,
			},
		},
		{
			name: "invalid json character",
			args: args{
				matches: []string{"", "value1", "val\"ue2"},
				fields:  []string{"field1", "field2"},
				index:   2,
			},
			want: want{
				got: `index=2 field1="value1" field2="val\"ue2"`,
				err: nil,
			},
		},
		{
			name: "more matches than fields",
			args: args{
				matches: []string{"", "value1", "value2", "value3"},
				fields:  []string{"field1", "field2"},
				index:   3,
			},
			want: want{
				got: `index=3 field1="value1" field2="value2"`,
				err: nil,
			},
		},
		{
			name: "more fields than matches",
			args: args{
				matches: []string{"", "value1"},
				fields:  []string{"field1", "field2"},
				index:   4,
			},
			want: want{
				got: `index=4 field1="value1"`,
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := textLineHandler(tt.args.matches, tt.args.fields, tt.args.index)
			if err != nil && err.Error() != tt.want.err.Error() {
				t.Errorf("got: %v, want: %v", err.Error(), tt.want.err.Error())
			}
			if !reflect.DeepEqual(got, tt.want.got) {
				t.Errorf("got: %v, want: %v", got, tt.want.got)
			}
		})
	}
}

func Test_textMetadataHandler(t *testing.T) {
	type args struct {
		m *parser.Metadata
	}
	type want struct {
		got string
		err error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "basic",
			args: args{
				m: &parser.Metadata{
					Total:     5,
					Matched:   3,
					Unmatched: 1,
					Skipped:   1,
					Source:    "",
					Errors: []parser.ErrorRecord{
						{
							Index:  1,
							Record: "aaa bbb ccc",
						},
					},
				},
			},
			want: want{
				got: `total=5 matched=3 unmatched=1 skipped=1 source="" errors=[{"index":1,"record":"aaa bbb ccc"}]`,
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := textMetadataHandler(tt.args.m)
			if err != nil && err.Error() != tt.want.err.Error() {
				t.Errorf("got: %v, want: %v", err.Error(), tt.want.err.Error())
			}
			if !reflect.DeepEqual(got, tt.want.got) {
				t.Errorf("got: %v, want: %v", got, tt.want.got)
			}
		})
	}
}
