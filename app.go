package main

import (
	"fmt"
	"strings"

	"github.com/nekrassov01/access-log-parser"
	"github.com/urfave/cli/v2"
)

type OutputFormat int

const (
	Text OutputFormat = iota
	JSON
	PrettyJSON
)

func (o OutputFormat) String() string {
	if o >= 0 && int(o) < len(outputFormat) {
		return outputFormat[o]
	}
	return ""
}

var (
	outputFormat = []string{
		"text",
		"json",
		"pretty-json",
	}

	bufferFlag = &cli.StringFlag{
		Name:    "buffer",
		Aliases: []string{"b"},
		Usage:   "input from buffer",
	}

	fileFlag = &cli.StringFlag{
		Name:    "file-path",
		Aliases: []string{"f"},
		Usage:   "input from file path",
	}

	gzipFlag = &cli.StringFlag{
		Name:    "gzip-path",
		Aliases: []string{"g"},
		Usage:   "input from gzip file path",
	}

	zipFlag = &cli.StringFlag{
		Name:    "zip-path",
		Aliases: []string{"z"},
		Usage:   "input from zip file path",
	}

	outputFlag = &cli.StringFlag{
		Name:    "output",
		Aliases: []string{"o"},
		Usage:   fmt.Sprintf("select output format: %s", pipeJoin(outputFormat)),
		Value:   Text.String(),
	}

	skipFlag = &cli.IntSliceFlag{
		Name:    "skip",
		Aliases: []string{"s"},
		Usage:   "skip records by index",
	}

	metadataFlag = &cli.BoolFlag{
		Name:    "metadata",
		Aliases: []string{"m"},
		Usage:   "enable metadata output",
	}

	globFlag = &cli.StringFlag{
		Name:    "glob-pattern",
		Aliases: []string{"G"},
		Usage:   "filter glob pattern: available for parsing zip only",
		Value:   "*",
	}
)

func NewApp() *cli.App {
	return &cli.App{
		Name:        Name,
		Usage:       "AWS access log parser",
		Version:     Version,
		Description: "A cli application for parsing AWS access logs",
		Action:      func(c *cli.Context) error { return nil },
		Commands: []*cli.Command{
			{
				Name:        "s3",
				Description: "Parses S3 access logs and converts them to structured formats",
				Usage:       "Parses S3 access logs",
				UsageText:   fmt.Sprintf("%s s3", Name),
				Action:      parseS3Logs,
				Flags:       []cli.Flag{bufferFlag, fileFlag, gzipFlag, zipFlag, outputFlag, skipFlag, metadataFlag, globFlag},
			},
			{
				Name:        "cf",
				Description: "Parses CloudFront access logs and converts them to structured formats",
				Usage:       "Parses CloudFront access logs",
				UsageText:   fmt.Sprintf("%s cf", Name),
				Action:      parseCFLogs,
				Flags:       []cli.Flag{bufferFlag, fileFlag, gzipFlag, zipFlag, outputFlag, skipFlag, metadataFlag, globFlag},
			},
			{
				Name:        "alb",
				Description: "Parses ALB access logs and converts them to structured formats",
				Usage:       "Parses ALB access logs",
				UsageText:   fmt.Sprintf("%s alb", Name),
				Action:      parseALBLogs,
				Flags:       []cli.Flag{bufferFlag, fileFlag, gzipFlag, zipFlag, outputFlag, skipFlag, metadataFlag, globFlag},
			},
			{
				Name:        "nlb",
				Description: "Parses NLB access logs and converts them to structured formats",
				Usage:       "Parses NLB access logs",
				UsageText:   fmt.Sprintf("%s nlb", Name),
				Action:      parseNLBLogs,
				Flags:       []cli.Flag{bufferFlag, fileFlag, gzipFlag, zipFlag, outputFlag, skipFlag, metadataFlag, globFlag},
			},
			{
				Name:        "clb",
				Description: "Parses CLB access logs and converts them to structured formats",
				Usage:       "Parses CLB access logs",
				UsageText:   fmt.Sprintf("%s clb", Name),
				Action:      parseCLBLogs,
				Flags:       []cli.Flag{bufferFlag, fileFlag, gzipFlag, zipFlag, outputFlag, skipFlag, metadataFlag, globFlag},
			},
		},
	}
}

type parserMap map[string]*parser.Parser

func parseS3Logs(c *cli.Context) error {
	return parseLogs(c, parserMap{
		Text.String():       parser.New(s3Fields, s3Patterns, textLineHandler, textMetadataHandler),
		JSON.String():       parser.New(s3Fields, s3Patterns, nil, nil),
		PrettyJSON.String(): parser.New(s3Fields, s3Patterns, prettyJSONLineHandler, prettyJSONMetadataHandler),
	})
}

func parseCFLogs(c *cli.Context) error {
	return parseLogs(c, parserMap{
		Text.String():       parser.New(cfFields, cfPatterns, textLineHandler, textMetadataHandler),
		JSON.String():       parser.New(cfFields, cfPatterns, nil, nil),
		PrettyJSON.String(): parser.New(cfFields, cfPatterns, prettyJSONLineHandler, prettyJSONMetadataHandler),
	})
}

func parseALBLogs(c *cli.Context) error {
	return parseLogs(c, parserMap{
		Text.String():       parser.New(albFields, albPatterns, textLineHandler, textMetadataHandler),
		JSON.String():       parser.New(albFields, albPatterns, nil, nil),
		PrettyJSON.String(): parser.New(albFields, albPatterns, prettyJSONLineHandler, prettyJSONMetadataHandler),
	})
}

func parseNLBLogs(c *cli.Context) error {
	return parseLogs(c, parserMap{
		Text.String():       parser.New(nlbFields, nlbPatterns, textLineHandler, textMetadataHandler),
		JSON.String():       parser.New(nlbFields, nlbPatterns, nil, nil),
		PrettyJSON.String(): parser.New(nlbFields, nlbPatterns, prettyJSONLineHandler, prettyJSONMetadataHandler),
	})
}

func parseCLBLogs(c *cli.Context) error {
	return parseLogs(c, parserMap{
		Text.String():       parser.New(clbFields, clbPatterns, textLineHandler, textMetadataHandler),
		JSON.String():       parser.New(clbFields, clbPatterns, nil, nil),
		PrettyJSON.String(): parser.New(clbFields, clbPatterns, prettyJSONLineHandler, prettyJSONMetadataHandler),
	})
}

func parseLogs(c *cli.Context, m parserMap) error {
	p, ok := m[c.String(outputFlag.Name)]
	if !ok {
		return fmt.Errorf("cannot parse command line flags: valid output format %s", pipeJoin(outputFormat))
	}
	if err := isSingle(c, bufferFlag.Name, fileFlag.Name, gzipFlag.Name, zipFlag.Name); err != nil {
		return err
	}
	if err := isValidPair(c, zipFlag.Name, globFlag.Name); err != nil {
		return err
	}
	result, results, err := dispatch(c, p)
	if err != nil {
		return err
	}
	if result != nil {
		if err := printResult(c, result); err != nil {
			return err
		}
	} else {
		for _, r := range results {
			if err := printResult(c, r); err != nil {
				return err
			}
		}
	}
	return nil
}

func dispatch(c *cli.Context, p *parser.Parser) (result *parser.Result, results []*parser.Result, err error) {
	switch {
	case c.IsSet(bufferFlag.Name):
		result, err = p.ParseString(c.String(bufferFlag.Name), c.IntSlice(skipFlag.Name))
	case c.IsSet(fileFlag.Name):
		result, err = p.ParseFile(c.String(fileFlag.Name), c.IntSlice(skipFlag.Name))
	case c.IsSet(gzipFlag.Name):
		result, err = p.ParseGzip(c.String(gzipFlag.Name), c.IntSlice(skipFlag.Name))
	case c.IsSet(zipFlag.Name):
		results, err = p.ParseZipEntries(c.String(zipFlag.Name), c.IntSlice(skipFlag.Name), c.String(globFlag.Name))
	default:
		return nil, nil, fmt.Errorf("cannot parse command line flags: no valid input provided %s", pipeJoin([]string{bufferFlag.Name, fileFlag.Name, gzipFlag.Name, zipFlag.Name}))
	}
	return result, results, err
}

func printResult(c *cli.Context, r *parser.Result) error {
	if _, err := fmt.Println(strings.Join(r.Data, "\n")); err != nil {
		return err
	}
	if c.Bool(metadataFlag.Name) {
		if _, err := fmt.Println(r.Metadata); err != nil {
			return err
		}
	}
	return nil
}

func isSingle(c *cli.Context, flags ...string) error {
	i := 0
	for _, flag := range flags {
		if c.IsSet(flag) {
			i++
		}
	}
	if i > 1 {
		return fmt.Errorf("cannot parse command line flags: only one flag can be used %s", pipeJoin(flags))
	}
	return nil
}

func isValidPair(c *cli.Context, flagA, flagB string) error {
	if !c.IsSet(flagA) && c.IsSet(flagB) {
		return fmt.Errorf("cannot parse command line flags: %s is available for %s only", flagB, flagA)
	}
	return nil
}

func pipeJoin(s []string) string {
	return strings.Join(s, "|")
}
