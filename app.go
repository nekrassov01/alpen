package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"

	"github.com/nekrassov01/access-log-parser"
	"github.com/urfave/cli/v2"
)

//go:embed completion/alpen.bash
var bashCompletion string

//go:embed completion/alpen.zsh
var zshCompletion string

//go:embed completion/alpen.ps1
var pwshCompletion string

type ShellCompletion int

const (
	bash ShellCompletion = iota
	zsh
	pwsh
)

var shellCompletion = []string{
	"bash",
	"zsh",
	"pwsh",
}

func (s ShellCompletion) String() string {
	if s >= 0 && int(s) < len(shellCompletion) {
		return shellCompletion[s]
	}
	return ""
}

type OutputFormat int

const (
	Text OutputFormat = iota
	JSON
	PrettyJSON
)

var outputFormat = []string{
	"text",
	"json",
	"pretty-json",
}

func (o OutputFormat) String() string {
	if o >= 0 && int(o) < len(outputFormat) {
		return outputFormat[o]
	}
	return ""
}

var (
	completion string
	buffer     string
	file       string
	gzip       string
	zip        string
	output     string
	skip       cli.IntSlice
	metadata   bool
	glob       string
)

var (
	bufferFlag = &cli.StringFlag{
		Name:        "buffer",
		Aliases:     []string{"b"},
		Usage:       "input from buffer",
		Destination: &buffer,
	}

	fileFlag = &cli.PathFlag{
		Name:        "file-path",
		Aliases:     []string{"f"},
		Usage:       "input from file path",
		Destination: &file,
	}

	gzipFlag = &cli.PathFlag{
		Name:        "gzip-path",
		Aliases:     []string{"g"},
		Usage:       "input from gzip file path",
		Destination: &gzip,
	}

	zipFlag = &cli.PathFlag{
		Name:        "zip-path",
		Aliases:     []string{"z"},
		Usage:       "input from zip file path",
		Destination: &zip,
	}

	outputFlag = &cli.StringFlag{
		Name:        "output",
		Aliases:     []string{"o"},
		Usage:       fmt.Sprintf("select output format: %s", pipeJoin(outputFormat)),
		Destination: &output,
		Value:       JSON.String(),
	}

	skipFlag = &cli.IntSliceFlag{
		Name:        "skip",
		Aliases:     []string{"s"},
		Usage:       "skip records by index",
		Destination: &skip,
	}

	metadataFlag = &cli.BoolFlag{
		Name:        "metadata",
		Aliases:     []string{"m"},
		Usage:       "enable metadata output",
		Destination: &metadata,
	}

	globFlag = &cli.StringFlag{
		Name:        "glob-pattern",
		Aliases:     []string{"G"},
		Usage:       "filter glob pattern: available for parsing zip only",
		Destination: &glob,
		Value:       "*",
	}
)

func NewApp() *cli.App {
	return &cli.App{
		Name:                 Name,
		Usage:                "AWS log parser/encoder",
		Version:              Version,
		Description:          "A cli application for parsing AWS access logs",
		HideHelpCommand:      true,
		EnableBashCompletion: true,
		Action:               doRootAction,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "completion",
				Aliases:     []string{"c"},
				Usage:       fmt.Sprintf("select a shellCompletion to display completion scripts: %s", pipeJoin(shellCompletion)),
				Destination: &completion,
				Action:      doCompletion,
			},
		},
		Commands: []*cli.Command{
			{
				Name:            "s3",
				Description:     "Parses S3 access logs and converts them to structured formats",
				Usage:           "Parses S3 access logs",
				UsageText:       fmt.Sprintf("%s s3", Name),
				HideHelpCommand: true,
				Flags:           []cli.Flag{bufferFlag, fileFlag, gzipFlag, zipFlag, outputFlag, skipFlag, metadataFlag, globFlag},
				Before:          doValidate,
				Action:          doS3Action,
			},
			{
				Name:            "cf",
				Description:     "Parses CloudFront access logs and converts them to structured formats",
				Usage:           "Parses CloudFront access logs",
				UsageText:       fmt.Sprintf("%s cf", Name),
				HideHelpCommand: true,
				Flags:           []cli.Flag{bufferFlag, fileFlag, gzipFlag, zipFlag, outputFlag, skipFlag, metadataFlag, globFlag},
				Before:          doValidate,
				Action:          doCFAction,
			},
			{
				Name:            "alb",
				Description:     "Parses ALB access logs and converts them to structured formats",
				Usage:           "Parses ALB access logs",
				UsageText:       fmt.Sprintf("%s alb", Name),
				HideHelpCommand: true,
				Flags:           []cli.Flag{bufferFlag, fileFlag, gzipFlag, zipFlag, outputFlag, skipFlag, metadataFlag, globFlag},
				Before:          doValidate,
				Action:          doALBAction,
			},
			{
				Name:            "nlb",
				Description:     "Parses NLB access logs and converts them to structured formats",
				Usage:           "Parses NLB access logs",
				UsageText:       fmt.Sprintf("%s nlb", Name),
				HideHelpCommand: true,
				Flags:           []cli.Flag{bufferFlag, fileFlag, gzipFlag, zipFlag, outputFlag, skipFlag, metadataFlag, globFlag},
				Before:          doValidate,
				Action:          doNLBAction,
			},
			{
				Name:            "clb",
				Description:     "Parses CLB access logs and converts them to structured formats",
				Usage:           "Parses CLB access logs",
				UsageText:       fmt.Sprintf("%s clb", Name),
				HideHelpCommand: true,
				Flags:           []cli.Flag{bufferFlag, fileFlag, gzipFlag, zipFlag, outputFlag, skipFlag, metadataFlag, globFlag},
				Before:          doValidate,
				Action:          doCLBAction,
			},
		},
	}
}

func doS3Action(c *cli.Context) error {
	return doAction(c, generateS3Patterns())
}

func doCFAction(c *cli.Context) error {
	return doAction(c, generateCFPatterns())
}

func doALBAction(c *cli.Context) error {
	return doAction(c, generateALBPatterns())
}

func doNLBAction(c *cli.Context) error {
	return doAction(c, generateNLBPatterns())
}

func doCLBAction(c *cli.Context) error {
	return doAction(c, generateCLBPatterns())
}

func doAction(c *cli.Context, patterns []*regexp.Regexp) error {
	p, err := newParser(c)
	if err != nil {
		return err
	}
	if err := p.AddPatterns(patterns); err != nil {
		return err
	}
	result, results, err := dispatch(c, p)
	if err != nil {
		return err
	}
	printResult(result, results)
	return nil
}

func newParser(c *cli.Context) (*parser.Parser, error) {
	switch c.String(outputFlag.Name) {
	case Text.String():
		return parser.NewParser(
			parser.WithLineHandler(textLineHandler),
			parser.WithMetadataHandler(textMetadataHandler),
		), nil
	case JSON.String():
		return parser.NewParser(), nil
	case PrettyJSON.String():
		return parser.NewParser(
			parser.WithLineHandler(prettyJSONLineHandler),
			parser.WithMetadataHandler(prettyJSONMetadataHandler),
		), nil
	default:
		return nil, fmt.Errorf(
			"cannot parse command line flags: invalid output format: allowed values: %s",
			pipeJoin(outputFormat),
		)
	}
}

func dispatch(c *cli.Context, p *parser.Parser) (result *parser.Result, results []*parser.Result, err error) {
	switch {
	case c.IsSet(bufferFlag.Name):
		result, err = p.ParseString(buffer, skip.Value())
		return result, nil, err
	case c.IsSet(fileFlag.Name):
		result, err = p.ParseFile(file, skip.Value())
		return result, nil, err
	case c.IsSet(gzipFlag.Name):
		result, err = p.ParseGzip(gzip, skip.Value())
		return result, nil, err
	case c.IsSet(zipFlag.Name):
		results, err = p.ParseZipEntries(zip, skip.Value(), glob)
		return nil, results, err
	default:
		return nil, nil, fmt.Errorf(
			"cannot parse command line flags: no valid input provided: %s",
			pipeJoin([]string{bufferFlag.Name, fileFlag.Name, gzipFlag.Name, zipFlag.Name}),
		)
	}
}

func printResult(result *parser.Result, results []*parser.Result) {
	var builder strings.Builder
	w := func(r *parser.Result) {
		for i, data := range r.Data {
			if i > 0 {
				builder.WriteRune('\n')
			}
			builder.WriteString(data)
		}
		builder.WriteRune('\n')
		if metadata {
			builder.WriteString(r.Metadata)
			builder.WriteRune('\n')
		}
	}
	switch {
	case result != nil && results == nil:
		w(result)
	case result == nil && results != nil:
		for _, r := range results {
			w(r)
		}
	default:
	}
	fmt.Println(builder.String())
}

func doCompletion(_ *cli.Context, s string) error {
	switch s {
	case bash.String():
		fmt.Println(bashCompletion)
	case zsh.String():
		fmt.Println(zshCompletion)
	case pwsh.String():
		fmt.Println(pwshCompletion)
	default:
		return fmt.Errorf(
			"cannot parse command line flags: invalid completion shellCompletion: allowed values: %s",
			pipeJoin(shellCompletion),
		)
	}
	return nil
}

func doRootAction(c *cli.Context) error {
	if c.Args().Len() == 0 && c.NumFlags() == 0 {
		return fmt.Errorf("cannot parse command line flags: no flag provided")
	}
	return nil
}

func doValidate(c *cli.Context) error {
	if err := isSingle(c, bufferFlag.Name, fileFlag.Name, gzipFlag.Name, zipFlag.Name); err != nil {
		return err
	}
	return isValidPair(c, zipFlag.Name, globFlag.Name)
}

func isSingle(c *cli.Context, flags ...string) error {
	i := 0
	for _, flag := range flags {
		if c.IsSet(flag) {
			i++
		}
	}
	if i > 1 {
		return fmt.Errorf("cannot parse command line flags: only one flag can be used: %s", pipeJoin(flags))
	}
	return nil
}

func isValidPair(c *cli.Context, flagA, flagB string) error {
	if !c.IsSet(flagA) && c.IsSet(flagB) {
		return fmt.Errorf("cannot parse command line flags: `%s` is available for `%s` only", flagB, flagA)
	}
	return nil
}

func pipeJoin(s []string) string {
	return strings.Join(s, "|")
}
