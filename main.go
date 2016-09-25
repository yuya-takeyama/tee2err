package main

import (
	"fmt"
	"io"
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/yuya-takeyama/argf"
)

// AppName is displayed in help command
const AppName = "tee2err"

type options struct {
	ShowVersion bool `short:"v" long:"version" description:"Show version"`
}

var opts options

func main() {
	parser := flags.NewParser(&opts, flags.Default^flags.PrintErrors)
	parser.Name = AppName
	parser.Usage = "[OPPTIONS] FILES..."

	files, err := parser.Parse()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}

	if opts.ShowVersion {
		fmt.Fprintf(os.Stdout, "%s v%s, build %s\n", AppName, Version, GitCommit)
		return
	}

	r, err := argf.From(files)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	w := io.MultiWriter(os.Stdout, os.Stderr)
	io.Copy(w, r)
}
