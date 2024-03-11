package main

import (
	_ "embed"
)

//go:embed completions/alpen.bash
var bashCompletion string

//go:embed completions/alpen.zsh
var zshCompletion string

//go:embed completions/alpen.ps1
var pwshCompletion string

type shell int

const (
	shbash shell = iota
	shzsh
	shpwsh
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

type output int

const (
	outputJSON output = iota
	outputPrettyJSON
	outputText
	outputLTSV
	outputTSV
)

var outputs = []string{
	"json",
	"pretty-json",
	"text",
	"ltsv",
	"tsv",
}

func (o output) String() string {
	if o >= 0 && int(o) < len(outputs) {
		return outputs[o]
	}
	return ""
}

type input int

const (
	inputStdin input = iota
	inputGzip
	inputZip
)

var inputs = []string{
	"default",
	"gz",
	"zip",
}

func (i input) String() string {
	if i >= 0 && int(i) < len(inputs) {
		return inputs[i]
	}
	return ""
}
