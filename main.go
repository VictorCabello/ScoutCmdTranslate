/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"traductor/cmd"

	"github.com/docopt/docopt-go"
)

func main() {
	usage := `traductor.

Usage:
  traductor [--reverse] <msg>
  traductor -r  | --reverse

Options:
  -h --help     Show this screen.
  -r --reverse   Parse in the other way`

	arguments, _ := docopt.ParseDoc(usage)
	msg := arguments["<msg>"].(string)
	reverse := arguments["--reverse"].(bool)
	cmd.Execute(msg, reverse)
}
