package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"

	"github.com/nekrassov01/access-log-parser"
	"github.com/urfave/cli/v2"
)

var (
	//go:embed completion/alpen.bash
	bashCompletion string

	//go:embed completion/alpen.zsh
	zshCompletion string

	//go:embed completion/alpen.ps1
	pwshCompletion string

	completion = []string{
		"bash",
		"zsh",
		"pwsh",
	}

	outputFormat = []string{
		"text",
		"json",
		"pretty-json",
	}
)

type Completion int

const (
	bash Completion = iota
	zsh
	pwsh
)

func (c Completion) String() string {
	if c >= 0 && int(c) < len(completion) {
		return completion[c]
	}
	return ""
}

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
		Value:   JSON.String(),
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
		Name:                 Name,
		Usage:                "AWS log parser/encoder",
		Version:              Version,
		Description:          "A cli application for parsing AWS access logs",
		HideHelpCommand:      true,
		EnableBashCompletion: true,
		Action:               doRootAction,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "completion",
				Aliases: []string{"c"},
				Usage:   fmt.Sprintf("select a shell to display completion scripts: %s", pipeJoin(completion)),
				Action:  doCompletion,
			},
		},
		Commands: []*cli.Command{
			{
				Name:            "s3",
				Description:     "Parses S3 access logs and converts them to structured formats",
				Usage:           "Parses S3 access logs",
				UsageText:       fmt.Sprintf("%s s3", Name),
				Action:          doS3Action,
				Flags:           []cli.Flag{bufferFlag, fileFlag, gzipFlag, zipFlag, outputFlag, skipFlag, metadataFlag, globFlag},
				Before:          validateFlags,
				HideHelpCommand: true,
			},
			{
				Name:            "cf",
				Description:     "Parses CloudFront access logs and converts them to structured formats",
				Usage:           "Parses CloudFront access logs",
				UsageText:       fmt.Sprintf("%s cf", Name),
				Action:          doCFAction,
				Flags:           []cli.Flag{bufferFlag, fileFlag, gzipFlag, zipFlag, outputFlag, skipFlag, metadataFlag, globFlag},
				Before:          validateFlags,
				HideHelpCommand: true,
			},
			{
				Name:            "alb",
				Description:     "Parses ALB access logs and converts them to structured formats",
				Usage:           "Parses ALB access logs",
				UsageText:       fmt.Sprintf("%s alb", Name),
				Action:          doALBAction,
				Flags:           []cli.Flag{bufferFlag, fileFlag, gzipFlag, zipFlag, outputFlag, skipFlag, metadataFlag, globFlag},
				Before:          validateFlags,
				HideHelpCommand: true,
			},
			{
				Name:            "nlb",
				Description:     "Parses NLB access logs and converts them to structured formats",
				Usage:           "Parses NLB access logs",
				UsageText:       fmt.Sprintf("%s nlb", Name),
				Action:          doNLBAction,
				Flags:           []cli.Flag{bufferFlag, fileFlag, gzipFlag, zipFlag, outputFlag, skipFlag, metadataFlag, globFlag},
				Before:          validateFlags,
				HideHelpCommand: true,
			},
			{
				Name:            "clb",
				Description:     "Parses CLB access logs and converts them to structured formats",
				Usage:           "Parses CLB access logs",
				UsageText:       fmt.Sprintf("%s clb", Name),
				Action:          doCLBAction,
				Flags:           []cli.Flag{bufferFlag, fileFlag, gzipFlag, zipFlag, outputFlag, skipFlag, metadataFlag, globFlag},
				Before:          validateFlags,
				HideHelpCommand: true,
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
	printResult(c, result, results)
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
		result, err = p.ParseString(c.String(bufferFlag.Name), c.IntSlice(skipFlag.Name))
		return result, nil, err
	case c.IsSet(fileFlag.Name):
		result, err = p.ParseFile(c.String(fileFlag.Name), c.IntSlice(skipFlag.Name))
		return result, nil, err
	case c.IsSet(gzipFlag.Name):
		result, err = p.ParseGzip(c.String(gzipFlag.Name), c.IntSlice(skipFlag.Name))
		return result, nil, err
	case c.IsSet(zipFlag.Name):
		results, err = p.ParseZipEntries(c.String(zipFlag.Name), c.IntSlice(skipFlag.Name), c.String(globFlag.Name))
		return nil, results, err
	default:
		return nil, nil, fmt.Errorf(
			"cannot parse command line flags: no valid input provided: %s",
			pipeJoin([]string{bufferFlag.Name, fileFlag.Name, gzipFlag.Name, zipFlag.Name}),
		)
	}
}

func printResult(c *cli.Context, result *parser.Result, results []*parser.Result) {
	var builder strings.Builder
	w := func(c *cli.Context, r *parser.Result) {
		for i, data := range r.Data {
			if i > 0 {
				builder.WriteRune('\n')
			}
			builder.WriteString(data)
		}
		builder.WriteRune('\n')
		if c.Bool(metadataFlag.Name) {
			builder.WriteString(r.Metadata)
			builder.WriteRune('\n')
		}
	}
	switch {
	case result != nil && results == nil:
		w(c, result)
	case result == nil && results != nil:
		for _, r := range results {
			w(c, r)
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
			"cannot parse command line flags: invalid completion shell: allowed values: %s",
			pipeJoin(completion),
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

func validateFlags(c *cli.Context) error {
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
