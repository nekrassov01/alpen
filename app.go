package main

import (
	_ "embed"
	"fmt"
	"strings"

	parser "github.com/nekrassov01/access-log-parser"
	"github.com/urfave/cli/v2"
)

//go:embed completions/alpen.bash
var bashCompletion string

//go:embed completions/alpen.zsh
var zshCompletion string

//go:embed completions/alpen.ps1
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
	JSON format = iota
	PrettyJSON
	Text
	LTSV
	TSV
)

var formats = []string{
	"json",
	"pretty-json",
	"text",
	"ltsv",
	"tsv",
}

func (f format) String() string {
	if f >= 0 && int(f) < len(formats) {
		return formats[f]
	}
	return ""
}

type app struct {
	cli  *cli.App
	dest dest
	flag flag
}

type dest struct {
	input    string
	file     string
	gzip     string
	zip      string
	output   string
	skip     cli.IntSlice
	metadata bool
	lineNum  bool
	header   bool
	glob     string
}

type flag struct {
	input    *cli.StringFlag
	file     *cli.PathFlag
	gzip     *cli.PathFlag
	zip      *cli.PathFlag
	output   *cli.StringFlag
	skip     *cli.IntSliceFlag
	metadata *cli.BoolFlag
	lineNum  *cli.BoolFlag
	header   *cli.BoolFlag
	glob     *cli.StringFlag
}

func newApp() *app {
	a := app{}
	a.flag.input = &cli.StringFlag{
		Name:        "input",
		Aliases:     []string{"i"},
		Usage:       "input from string",
		Destination: &a.dest.input,
	}
	a.flag.file = &cli.PathFlag{
		Name:        "file-path",
		Aliases:     []string{"f"},
		Usage:       "input from file path",
		Destination: &a.dest.file,
	}
	a.flag.gzip = &cli.PathFlag{
		Name:        "gzip-path",
		Aliases:     []string{"g"},
		Usage:       "input from gzip file path",
		Destination: &a.dest.gzip,
	}
	a.flag.zip = &cli.PathFlag{
		Name:        "zip-path",
		Aliases:     []string{"z"},
		Usage:       "input from zip file path",
		Destination: &a.dest.zip,
	}
	a.flag.output = &cli.StringFlag{
		Name:        "output",
		Aliases:     []string{"o"},
		Usage:       fmt.Sprintf("select output format: %s", pipeJoin(formats)),
		Destination: &a.dest.output,
		Value:       JSON.String(),
	}
	a.flag.skip = &cli.IntSliceFlag{
		Name:        "skip",
		Aliases:     []string{"s"},
		Usage:       "skip records by index",
		Destination: &a.dest.skip,
	}
	a.flag.metadata = &cli.BoolFlag{
		Name:        "metadata",
		Aliases:     []string{"m"},
		Usage:       "enable metadata output",
		Destination: &a.dest.metadata,
	}
	a.flag.lineNum = &cli.BoolFlag{
		Name:        "line-number",
		Aliases:     []string{"l"},
		Usage:       "set line number at the beginning of the line",
		Destination: &a.dest.lineNum,
	}
	a.flag.header = &cli.BoolFlag{
		Name:        "header",
		Aliases:     []string{"H"},
		Usage:       "set header: avairable for tsv output",
		Destination: &a.dest.header,
		Value:       false,
	}
	a.flag.glob = &cli.StringFlag{
		Name:        "glob-pattern",
		Aliases:     []string{"G"},
		Usage:       "filter glob pattern: available for parsing zip only",
		Destination: &a.dest.glob,
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
		a.flag.lineNum,
		a.flag.header,
		a.flag.glob,
	}
	a.cli = &cli.App{
		Name:                 Name,
		Usage:                "Access log parser/encoder CLI",
		Version:              Version,
		Description:          "A cli application for parsing various access logs",
		HideHelpCommand:      true,
		EnableBashCompletion: true,
		Action:               a.rootAction,
		Commands: []*cli.Command{
			{
				Name:            "completion",
				Description:     "Generate completion scripts for specified shell",
				Usage:           fmt.Sprintf("Generate completion scripts for specified shell: %s", pipeJoin(shells)),
				UsageText:       fmt.Sprintf("%s completion %s", Name, pipeJoin(shells)),
				HideHelpCommand: true,
				Action:          completionAction,
			},
			{
				Name:            "clf",
				Description:     "Parses apache common/combined log format and converts them to structured formats",
				Usage:           "Parses apache common/combined log format",
				UsageText:       fmt.Sprintf("%s clf", Name),
				HideHelpCommand: true,
				Flags:           flags,
				Before:          a.before,
				Action:          a.clfAction,
			},
			{
				Name:            "clfv",
				Description:     "Parses apache common/combined log format with vhost and converts them to structured formats",
				Usage:           "Parses apache common/combined log format with vhost",
				UsageText:       fmt.Sprintf("%s clfv", Name),
				HideHelpCommand: true,
				Flags:           flags,
				Before:          a.before,
				Action:          a.clfvAction,
			},
			{
				Name:            "s3",
				Description:     "Parses S3 access logs and converts them to structured formats",
				Usage:           "Parses S3 access logs",
				UsageText:       fmt.Sprintf("%s s3", Name),
				HideHelpCommand: true,
				Flags:           flags,
				Before:          a.before,
				Action:          a.s3Action,
			},
			{
				Name:            "cf",
				Description:     "Parses CloudFront access logs and converts them to structured formats",
				Usage:           "Parses CloudFront access logs",
				UsageText:       fmt.Sprintf("%s cf", Name),
				HideHelpCommand: true,
				Flags:           flags,
				Before:          a.before,
				Action:          a.cfAction,
			},
			{
				Name:            "alb",
				Description:     "Parses ALB access logs and converts them to structured formats",
				Usage:           "Parses ALB access logs",
				UsageText:       fmt.Sprintf("%s alb", Name),
				HideHelpCommand: true,
				Flags:           flags,
				Before:          a.before,
				Action:          a.albAction,
			},
			{
				Name:            "nlb",
				Description:     "Parses NLB access logs and converts them to structured formats",
				Usage:           "Parses NLB access logs",
				UsageText:       fmt.Sprintf("%s nlb", Name),
				HideHelpCommand: true,
				Flags:           flags,
				Before:          a.before,
				Action:          a.nlbAction,
			},
			{
				Name:            "clb",
				Description:     "Parses CLB access logs and converts them to structured formats",
				Usage:           "Parses CLB access logs",
				UsageText:       fmt.Sprintf("%s clb", Name),
				HideHelpCommand: true,
				Flags:           flags,
				Before:          a.before,
				Action:          a.clbAction,
			},
			{
				Name:            "ltsv",
				Description:     "Parses LTSV format logs and converts them to other structured formats",
				Usage:           "Parses LTSV format logs",
				UsageText:       fmt.Sprintf("%s ltsv", Name),
				HideHelpCommand: true,
				Flags:           flags,
				Before:          a.before,
				Action:          a.doLTSVAction,
			},
		},
	}
	return &a
}

func (a *app) clfAction(c *cli.Context) error {
	p, err := a.newParser(c, parser.NewApacheCLFRegexParser())
	if err != nil {
		return err
	}
	return a.out(c, p)
}

func (a *app) clfvAction(c *cli.Context) error {
	p, err := a.newParser(c, parser.NewApacheCLFWithVHostRegexParser())
	if err != nil {
		return err
	}
	return a.out(c, p)
}

func (a *app) s3Action(c *cli.Context) error {
	p, err := a.newParser(c, parser.NewS3RegexParser())
	if err != nil {
		return err
	}
	return a.out(c, p)
}

func (a *app) cfAction(c *cli.Context) error {
	p, err := a.newParser(c, parser.NewCFRegexParser())
	if err != nil {
		return err
	}
	return a.out(c, p)
}

func (a *app) albAction(c *cli.Context) error {
	p, err := a.newParser(c, parser.NewALBRegexParser())
	if err != nil {
		return err
	}
	return a.out(c, p)
}

func (a *app) nlbAction(c *cli.Context) error {
	p, err := a.newParser(c, parser.NewNLBRegexParser())
	if err != nil {
		return err
	}
	return a.out(c, p)
}

func (a *app) clbAction(c *cli.Context) error {
	p, err := a.newParser(c, parser.NewCLBRegexParser())
	if err != nil {
		return err
	}
	return a.out(c, p)
}

func (a *app) doLTSVAction(c *cli.Context) error {
	p, err := a.newParser(c, parser.NewLTSVParser())
	if err != nil {
		return err
	}
	return a.out(c, p)
}

func (a *app) newParser(c *cli.Context, p parser.Parser) (parser.Parser, error) {
	switch c.String(a.flag.output.Name) {
	case JSON.String():
	case PrettyJSON.String():
		p.SetLineHandler(parser.PrettyJSONLineHandler)
		p.SetMetadataHandler(parser.PrettyJSONMetadataHandler)
	case Text.String():
		p.SetLineHandler(parser.KeyValuePairLineHandler)
		p.SetMetadataHandler(parser.KeyValuePairMetadataHandler)
	case LTSV.String():
		p.SetLineHandler(parser.LTSVLineHandler)
		p.SetMetadataHandler(parser.LTSVMetadataHandler)
	case TSV.String():
		p.SetLineHandler(parser.TSVLineHandler)
		p.SetMetadataHandler(parser.TSVMetadataHandler)
	default:
		return nil, fmt.Errorf(
			"cannot parse command line flags: invalid output format: allowed values: %s",
			pipeJoin(formats),
		)
	}
	return p, nil
}

func (a *app) out(c *cli.Context, p parser.Parser) error {
	result, results, err := a.dispatch(c, p)
	if err != nil {
		return err
	}
	a.printResult(result, results)
	return nil
}

func (a *app) dispatch(c *cli.Context, p parser.Parser) (result *parser.Result, results []*parser.Result, err error) {
	switch {
	case c.IsSet(a.flag.input.Name):
		result, err = p.ParseString(
			a.dest.input,
			a.dest.skip.Value(),
			a.dest.lineNum,
		)
		return result, nil, err
	case c.IsSet(a.flag.file.Name):
		result, err = p.ParseFile(
			a.dest.file,
			a.dest.skip.Value(),
			a.dest.lineNum,
		)
		return result, nil, err
	case c.IsSet(a.flag.gzip.Name):
		result, err = p.ParseGzip(
			a.dest.gzip,
			a.dest.skip.Value(),
			a.dest.lineNum,
		)
		return result, nil, err
	case c.IsSet(a.flag.zip.Name):
		results, err = p.ParseZipEntries(
			a.dest.zip,
			a.dest.skip.Value(),
			a.dest.lineNum,
			a.dest.glob,
		)
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
		if a.dest.header && a.dest.output == TSV.String() {
			writeTSVHeader(r.Labels[0], a.dest.lineNum)
		}
		for i, data := range r.Data {
			if i > 0 {
				builder.WriteRune('\n')
			}
			builder.WriteString(data)
		}
		builder.WriteRune('\n')
		if a.dest.metadata {
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

func (a *app) before(c *cli.Context) error {
	if err := checkSingle(c, a.flag.input.Name, a.flag.file.Name, a.flag.gzip.Name, a.flag.zip.Name); err != nil {
		return err
	}
	if err := checkValidPair(c, a.flag.zip.Name, a.flag.glob.Name); err != nil {
		return err
	}
	if a.dest.header && a.dest.output != TSV.String() {
		return fmt.Errorf("cannot parse command line flags: `header` is available for output `tsv` only")
	}
	return nil
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

func writeTSVHeader(line []string, lineNumber bool) {
	if lineNumber {
		line = append([]string{"index"}, line...)
	}
	fmt.Println(strings.Join(line, "\t"))
}

func (a *app) rootAction(c *cli.Context) error {
	if c.Args().Len() == 0 && c.NumFlags() == 0 {
		return fmt.Errorf("cannot parse command line flags: no flag provided")
	}
	return nil
}

func completionAction(c *cli.Context) error {
	switch c.Args().First() {
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
	return nil
}
