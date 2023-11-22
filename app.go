package main

import (
	_ "embed"
	"fmt"
	"os"
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

type shell int

const (
	bash shell = iota
	zsh
	pwsh
)

var shells = []string{
	"bash",
	"zsh",
	"pwsh",
}

func (s shell) String() string {
	if s >= 0 && int(s) < len(shells) {
		return shells[s]
	}
	return ""
}

type format int

const (
	Text format = iota
	JSON
	PrettyJSON
)

var formats = []string{
	"text",
	"json",
	"pretty-json",
}

func (f format) String() string {
	if f >= 0 && int(f) < len(formats) {
		return formats[f]
	}
	return ""
}

type app struct {
	cli *cli.App
	destination
	flag
}

type destination struct {
	completion string
	input      string
	file       string
	gzip       string
	zip        string
	output     string
	skip       cli.IntSlice
	metadata   bool
	glob       string
}

type flag struct {
	completion *cli.StringFlag
	input      *cli.StringFlag
	file       *cli.PathFlag
	gzip       *cli.PathFlag
	zip        *cli.PathFlag
	output     *cli.StringFlag
	skip       *cli.IntSliceFlag
	metadata   *cli.BoolFlag
	glob       *cli.StringFlag
}

func newApp() *app {
	a := app{}
	a.flag.completion = &cli.StringFlag{
		Name:        "completion",
		Aliases:     []string{"c"},
		Usage:       fmt.Sprintf("select a shell to display completion scripts: %s", pipeJoin(shells)),
		Destination: &a.destination.completion,
	}
	a.flag.input = &cli.StringFlag{
		Name:        "input",
		Aliases:     []string{"i"},
		Usage:       "input from string",
		Destination: &a.destination.input,
	}
	a.flag.file = &cli.PathFlag{
		Name:        "file-path",
		Aliases:     []string{"f"},
		Usage:       "input from file path",
		Destination: &a.destination.file,
	}
	a.flag.gzip = &cli.PathFlag{
		Name:        "gzip-path",
		Aliases:     []string{"g"},
		Usage:       "input from gzip file path",
		Destination: &a.destination.gzip,
	}
	a.flag.zip = &cli.PathFlag{
		Name:        "zip-path",
		Aliases:     []string{"z"},
		Usage:       "input from zip file path",
		Destination: &a.destination.zip,
	}
	a.flag.output = &cli.StringFlag{
		Name:        "output",
		Aliases:     []string{"o"},
		Usage:       fmt.Sprintf("select output format: %s", pipeJoin(formats)),
		Destination: &a.destination.output,
		Value:       JSON.String(),
	}
	a.flag.skip = &cli.IntSliceFlag{
		Name:        "skip",
		Aliases:     []string{"s"},
		Usage:       "skip records by index",
		Destination: &a.destination.skip,
	}
	a.flag.metadata = &cli.BoolFlag{
		Name:        "metadata",
		Aliases:     []string{"m"},
		Usage:       "enable metadata output",
		Destination: &a.destination.metadata,
	}
	a.flag.glob = &cli.StringFlag{
		Name:        "glob-pattern",
		Aliases:     []string{"G"},
		Usage:       "filter glob pattern: available for parsing zip only",
		Destination: &a.destination.glob,
		Value:       "*",
	}
	flags := []cli.Flag{
		a.flag.input,
		a.flag.file,
		a.flag.gzip,
		a.flag.zip,
		a.flag.output,
		a.flag.skip,
		a.flag.metadata,
		a.flag.glob,
	}
	a.cli = &cli.App{
		Name:                 Name,
		Usage:                "Access log parser/encoder CLI",
		Version:              Version,
		Description:          "A cli application for parsing various access logs",
		HideHelpCommand:      true,
		EnableBashCompletion: true,
		Action:               a.doRootAction,
		Flags:                []cli.Flag{a.flag.completion},
		Commands: []*cli.Command{
			{
				Name:            "clf",
				Description:     "Parses apache common/combined log format and converts them to structured formats",
				Usage:           "Parses apache common/combined log format",
				UsageText:       fmt.Sprintf("%s clf", Name),
				HideHelpCommand: true,
				Flags:           flags,
				Before:          a.doValidate,
				Action:          a.doApacheCLFAction,
			},
			{
				Name:            "clfv",
				Description:     "Parses apache common/combined log format with vhost and converts them to structured formats",
				Usage:           "Parses apache common/combined log format with vhost",
				UsageText:       fmt.Sprintf("%s clfv", Name),
				HideHelpCommand: true,
				Flags:           flags,
				Before:          a.doValidate,
				Action:          a.doApacheCLFWithVHostAction,
			},
			{
				Name:            "s3",
				Description:     "Parses S3 access logs and converts them to structured formats",
				Usage:           "Parses S3 access logs",
				UsageText:       fmt.Sprintf("%s s3", Name),
				HideHelpCommand: true,
				Flags:           flags,
				Before:          a.doValidate,
				Action:          a.doS3Action,
			},
			{
				Name:            "cf",
				Description:     "Parses CloudFront access logs and converts them to structured formats",
				Usage:           "Parses CloudFront access logs",
				UsageText:       fmt.Sprintf("%s cf", Name),
				HideHelpCommand: true,
				Flags:           flags,
				Before:          a.doValidate,
				Action:          a.doCFAction,
			},
			{
				Name:            "alb",
				Description:     "Parses ALB access logs and converts them to structured formats",
				Usage:           "Parses ALB access logs",
				UsageText:       fmt.Sprintf("%s alb", Name),
				HideHelpCommand: true,
				Flags:           flags,
				Before:          a.doValidate,
				Action:          a.doALBAction,
			},
			{
				Name:            "nlb",
				Description:     "Parses NLB access logs and converts them to structured formats",
				Usage:           "Parses NLB access logs",
				UsageText:       fmt.Sprintf("%s nlb", Name),
				HideHelpCommand: true,
				Flags:           flags,
				Before:          a.doValidate,
				Action:          a.doNLBAction,
			},
			{
				Name:            "clb",
				Description:     "Parses CLB access logs and converts them to structured formats",
				Usage:           "Parses CLB access logs",
				UsageText:       fmt.Sprintf("%s clb", Name),
				HideHelpCommand: true,
				Flags:           flags,
				Before:          a.doValidate,
				Action:          a.doCLBAction,
			},
		},
	}
	return &a
}

func (a *app) run() error {
	return a.cli.Run(os.Args)
}

func (a *app) doApacheCLFAction(c *cli.Context) error {
	return a.doAction(c, generateApacheCLFPatterns())
}

func (a *app) doApacheCLFWithVHostAction(c *cli.Context) error {
	return a.doAction(c, generateApacheCLFWithVHostPatterns())
}

func (a *app) doS3Action(c *cli.Context) error {
	return a.doAction(c, generateS3Patterns())
}

func (a *app) doCFAction(c *cli.Context) error {
	return a.doAction(c, generateCFPatterns())
}

func (a *app) doALBAction(c *cli.Context) error {
	return a.doAction(c, generateALBPatterns())
}

func (a *app) doNLBAction(c *cli.Context) error {
	return a.doAction(c, generateNLBPatterns())
}

func (a *app) doCLBAction(c *cli.Context) error {
	return a.doAction(c, generateCLBPatterns())
}

func (a *app) doAction(c *cli.Context, patterns []*regexp.Regexp) error {
	p, err := a.newParser(c)
	if err != nil {
		return err
	}
	if err := p.AddPatterns(patterns); err != nil {
		return err
	}
	result, results, err := a.dispatch(c, p)
	if err != nil {
		return err
	}
	a.printResult(result, results)
	return nil
}

func (a *app) newParser(c *cli.Context) (*parser.Parser, error) {
	switch c.String(a.flag.output.Name) {
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
			pipeJoin(formats),
		)
	}
}

func (a *app) dispatch(c *cli.Context, p *parser.Parser) (result *parser.Result, results []*parser.Result, err error) {
	switch {
	case c.IsSet(a.flag.input.Name):
		result, err = p.ParseString(a.destination.input, a.destination.skip.Value())
		return result, nil, err
	case c.IsSet(a.flag.file.Name):
		result, err = p.ParseFile(a.destination.file, a.destination.skip.Value())
		return result, nil, err
	case c.IsSet(a.flag.gzip.Name):
		result, err = p.ParseGzip(a.destination.gzip, a.destination.skip.Value())
		return result, nil, err
	case c.IsSet(a.flag.zip.Name):
		results, err = p.ParseZipEntries(a.destination.zip, a.destination.skip.Value(), a.destination.glob)
		return nil, results, err
	default:
		return nil, nil, fmt.Errorf(
			"cannot parse command line flags: no valid input provided: %s",
			pipeJoin([]string{a.flag.input.Name, a.flag.file.Name, a.flag.gzip.Name, a.flag.zip.Name}),
		)
	}
}

func (a *app) printResult(result *parser.Result, results []*parser.Result) {
	var builder strings.Builder
	w := func(r *parser.Result) {
		for i, data := range r.Data {
			if i > 0 {
				builder.WriteRune('\n')
			}
			builder.WriteString(data)
		}
		builder.WriteRune('\n')
		if a.destination.metadata {
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

func (a *app) doRootAction(c *cli.Context) error {
	if c.Args().Len() == 0 && c.NumFlags() == 0 {
		return fmt.Errorf("cannot parse command line flags: no flag provided")
	}
	if c.IsSet(a.flag.completion.Name) {
		switch a.destination.completion {
		case bash.String():
			fmt.Println(bashCompletion)
		case zsh.String():
			fmt.Println(zshCompletion)
		case pwsh.String():
			fmt.Println(pwshCompletion)
		default:
			return fmt.Errorf(
				"cannot parse command line flags: invalid completion shell: allowed values: %s",
				pipeJoin(shells),
			)
		}
	}
	return nil
}

func (a *app) doValidate(c *cli.Context) error {
	if err := checkSingle(c, a.flag.input.Name, a.flag.file.Name, a.flag.gzip.Name, a.flag.zip.Name); err != nil {
		return err
	}
	return checkValidPair(c, a.flag.zip.Name, a.flag.glob.Name)
}

func checkSingle(c *cli.Context, flags ...string) error {
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

func checkValidPair(c *cli.Context, a, b string) error {
	if !c.IsSet(a) && c.IsSet(b) {
		return fmt.Errorf("cannot parse command line flags: `%s` is available for `%s` only", b, a)
	}
	return nil
}

func pipeJoin(s []string) string {
	return strings.Join(s, "|")
}
