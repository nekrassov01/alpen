package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/nekrassov01/access-log-parser"
)

func prettyJSONLineHandler(matches []string, fields []string, index int) (string, error) {
	s, err := parser.DefaultLineHandler(matches, fields, index)
	if err != nil {
		return "", err
	}
	return prettyJSON(s)
}

func prettyJSONMetadataHandler(m *parser.Metadata) (string, error) {
	s, err := parser.DefaultMetadataHandler(m)
	if err != nil {
		return "", err
	}
	return prettyJSON(s)
}

func prettyJSON(s string) (string, error) {
	var buf bytes.Buffer
	if err := json.Indent(&buf, []byte(s), "", "  "); err != nil {
		return "", fmt.Errorf("cannot format string as json: %w", err)
	}
	return buf.String(), nil
}

func textLineHandler(matches []string, fields []string, index int) (string, error) {
	var builder strings.Builder
	builder.WriteString("index=")
	builder.WriteString(strconv.Itoa(index))
	for i, match := range matches {
		if i < len(fields) {
			builder.WriteString(" ")
			builder.WriteString(fields[i])
			builder.WriteString(`="`)
			if match == "\"-\"" {
				builder.WriteString("-")
			} else {
				for _, m := range match {
					if m == '"' {
						builder.WriteString(`\"`)
					} else {
						builder.WriteRune(m)
					}
				}
			}
			builder.WriteString(`"`)
		}
	}
	return builder.String(), nil
}

func textMetadataHandler(m *parser.Metadata) (string, error) {
	e, err := json.Marshal(m.Errors)
	if err != nil {
		return "", fmt.Errorf("cannot marshal errors as json: %w", err)
	}
	return fmt.Sprintf(
		"total=%d matched=%d unmatched=%d skipped=%d source=\"%s\" errors=%s",
		m.Total, m.Matched, m.Unmatched, m.Skipped, m.Source, e,
	), nil
}
