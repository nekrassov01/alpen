package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mattn/go-isatty"
	parser "github.com/nekrassov01/access-log-parser"
	"github.com/urfave/cli/v2"
)

type app struct {
	cli  *cli.App
	dest dest
	flag flag
}

type dest struct {
	input   string
	output  string
	result  bool
	glob    string
	labels  cli.StringSlice
	filters cli.StringSlice
	skip    cli.IntSlice
	prefix  bool
	unmatch bool
	num     bool
}

type flag struct {
	input   *cli.StringFlag
	output  *cli.StringFlag
	result  *cli.BoolFlag
	glob    *cli.StringFlag
	labels  *cli.StringSliceFlag
	filters *cli.StringSliceFlag
	skip    *cli.IntSliceFlag
	prefix  *cli.BoolFlag
	unmatch *cli.BoolFlag
	num     *cli.BoolFlag
}

func newApp() *app {
	a := app{}
	a.flag.input = &cli.StringFlag{
		Name:        "input",
		Aliases:     []string{"i"},
		Usage:       fmt.Sprintf("select input type: %s", pipeJoin(inputs)),
		Destination: &a.dest.input,
		Value:       inputStdin.String(),
	}
	a.flag.output = &cli.StringFlag{
		Name:        "output",
		Aliases:     []string{"o"},
		Usage:       fmt.Sprintf("select output type: %s", pipeJoin(outputs)),
		Destination: &a.dest.output,
		Value:       outputJSON.String(),
	}
	a.flag.result = &cli.BoolFlag{
		Name:        "result",
		Aliases:     []string{"r"},
		Usage:       "enable result output",
		Destination: &a.dest.result,
	}
	a.flag.glob = &cli.StringFlag{
		Name:        "glob",
		Aliases:     []string{"g"},
		Usage:       "filter glob pattern: available for parsing zip only",
		Destination: &a.dest.glob,
		Value:       "*",
	}
	a.flag.labels = &cli.StringSliceFlag{
		Name:        "labels",
		Aliases:     []string{"l"},
		Usage:       "select labels to output with labels",
		Destination: &a.dest.labels,
	}
	a.flag.filters = &cli.StringSliceFlag{
		Name:        "filters",
		Aliases:     []string{"f"},
		Usage:       "set filter expressions: allowed operator: >|>=|<|<=|==|!=|==*|!=*|=~|!~|=~*|!~*",
		Destination: &a.dest.filters,
	}
	a.flag.skip = &cli.IntSliceFlag{
		Name:        "skip",
		Aliases:     []string{"s"},
		Usage:       "skip lines by line number",
		Destination: &a.dest.skip,
	}
	a.flag.prefix = &cli.BoolFlag{
		Name:        "prefix",
		Aliases:     []string{"p"},
		Usage:       "enable line prefix: PROCESSED|UNMATCHED",
		Destination: &a.dest.prefix,
	}
	a.flag.unmatch = &cli.BoolFlag{
		Name:        "unmatch",
		Aliases:     []string{"u"},
		Usage:       "enable output of unmatched lines",
		Destination: &a.dest.unmatch,
	}
	a.flag.num = &cli.BoolFlag{
		Name:        "num",
		Aliases:     []string{"n"},
		Usage:       "set line number at the beginning of the line",
		Destination: &a.dest.num,
	}
	flags := []cli.Flag{
		a.flag.input,
		a.flag.output,
		a.flag.result,
		a.flag.glob,
		a.flag.labels,
		a.flag.filters,
		a.flag.skip,
		a.flag.prefix,
		a.flag.unmatch,
		a.flag.num,
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
	opt, err := a.opt()
	if err != nil {
		return err
	}
	return a.out(c, parser.NewApacheCLFRegexParser(c.Context, os.Stdout, opt))
}

func (a *app) clfvAction(c *cli.Context) error {
	opt, err := a.opt()
	if err != nil {
		return err
	}
	return a.out(c, parser.NewApacheCLFWithVHostRegexParser(c.Context, os.Stdout, opt))
}

func (a *app) s3Action(c *cli.Context) error {
	opt, err := a.opt()
	if err != nil {
		return err
	}
	return a.out(c, parser.NewS3RegexParser(c.Context, os.Stdout, opt))
}

func (a *app) cfAction(c *cli.Context) error {
	opt, err := a.opt()
	if err != nil {
		return err
	}
	return a.out(c, parser.NewCFRegexParser(c.Context, os.Stdout, opt))
}

func (a *app) albAction(c *cli.Context) error {
	opt, err := a.opt()
	if err != nil {
		return err
	}
	return a.out(c, parser.NewALBRegexParser(c.Context, os.Stdout, opt))
}

func (a *app) nlbAction(c *cli.Context) error {
	opt, err := a.opt()
	if err != nil {
		return err
	}
	return a.out(c, parser.NewNLBRegexParser(c.Context, os.Stdout, opt))
}

func (a *app) clbAction(c *cli.Context) error {
	opt, err := a.opt()
	if err != nil {
		return err
	}
	return a.out(c, parser.NewCLBRegexParser(c.Context, os.Stdout, opt))
}

func (a *app) doLTSVAction(c *cli.Context) error {
	opt, err := a.opt()
	if err != nil {
		return err
	}
	return a.out(c, parser.NewLTSVParser(c.Context, os.Stdout, opt))
}

func (a *app) before(c *cli.Context) error {
	if a.dest.input != inputZip.String() && c.IsSet(a.flag.glob.Name) {
		return fmt.Errorf("\"glob\" is only valid with zip")
	}
	return nil
}

func (a *app) rootAction(c *cli.Context) error {
	if c.Args().Len() == 0 && c.NumFlags() == 0 {
		return fmt.Errorf("no flag provided")
	}
	return nil
}

func completionAction(c *cli.Context) error {
	switch c.Args().First() {
	case shbash.String():
		fmt.Println(bashCompletion)
	case shzsh.String():
		fmt.Println(zshCompletion)
	case shpwsh.String():
		fmt.Println(pwshCompletion)
	default:
		return fmt.Errorf("invalid completion shell: allowed values: %s", pipeJoin(shells))
	}
	return nil
}

func (a *app) opt() (parser.Option, error) {
	opt := parser.Option{
		Labels:       a.flag.labels.GetDestination(),
		Filters:      a.flag.filters.GetDestination(),
		SkipLines:    a.flag.skip.GetDestination(),
		Prefix:       a.dest.prefix,
		UnmatchLines: a.dest.unmatch,
		LineNumber:   a.dest.num,
	}
	switch a.dest.output {
	case outputJSON.String():
		opt.LineHandler = parser.JSONLineHandler
	case outputPrettyJSON.String():
		opt.LineHandler = parser.PrettyJSONLineHandler
	case outputText.String():
		opt.LineHandler = parser.KeyValuePairLineHandler
	case outputLTSV.String():
		opt.LineHandler = parser.LTSVLineHandler
	case outputTSV.String():
		opt.LineHandler = parser.TSVLineHandler
	default:
		return parser.Option{}, fmt.Errorf("invalid output type: allowed values: %s", pipeJoin(outputs))
	}
	return opt, nil
}

func (a *app) out(c *cli.Context, p parser.Parser) error {
	var r *parser.Result
	var err error
	switch a.dest.input {
	case inputStdin.String():
		if isatty.IsTerminal(os.Stdin.Fd()) {
			r, err = p.ParseFile(c.Args().First())
		} else {
			r, err = p.Parse(os.Stdin)
		}
	case inputGzip.String():
		r, err = p.ParseGzip(c.Args().First())
	case inputZip.String():
		r, err = p.ParseZipEntries(c.Args().First(), a.dest.glob)
	default:
		return fmt.Errorf("invalid input type: allowed values: %s", pipeJoin(inputs))
	}
	if err != nil {
		return err
	}
	if a.dest.result {
		fmt.Println(r)
	}
	return nil
}

func pipeJoin(s []string) string {
	return strings.Join(s, "|")
}
