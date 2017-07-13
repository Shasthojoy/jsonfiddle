////////////////////////////////////////////////////////////////////////////
// Program: jf
// Purpose: JSON Fiddling
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

//go:generate sh -v jsonfiddleCLIGen.sh

import (
	"fmt"
	"os"

	"github.com/mkideal/cli"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

type OptsT struct {
	Prefix  string
	Indent  string
	Compact bool
	Verbose int
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	progname  = "jf"
	VERSION   = "0.1.0"
	buildTime = "2017-07-13"
)

var (
	rootArgv *rootT
	Opts     OptsT
)

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
func main() {
	//NOTE: You can set any writer implements io.Writer
	// default writer is os.Stdout
	if err := cli.Root(root,
		cli.Tree(sortDef)).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println("")
}

//==========================================================================
// Main dispatcher

func jf(ctx *cli.Context) error {
	ctx.JSON(ctx.RootArgv())
	ctx.JSON(ctx.Argv())
	fmt.Println()

	return nil
}

//==========================================================================
// support functions

// abortOn will quit on anticipated errors gracefully without stack trace
func abortOn(errCase string, e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "[%s] %s error: %v\n", progname, errCase, e)
		os.Exit(1)
	}
}

// verbose will print info to stderr according to the verbose level setting
func verbose(levelSet, levelNow int, format string, args ...interface{}) {
	if levelNow >= levelSet {
		fmt.Fprintf(os.Stderr, "["+progname+"] "+format+"\n", args...)
	}
}
