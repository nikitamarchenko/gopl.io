/*
Exercise 4.11: Build a tool that lets users create, read, update, and delete
GitHub issues from the command line, invoking their preferred text editor when
substantial text input is required.
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"log"
)

var token *string
var repo *string
var debug *bool

func init() {
	token = flag.String("token", "", "github api token")
	repo = flag.String("repo", "", "repo in format {owner}/{repo}")
	debug = flag.Bool("debug", false, "debug output") 
}

func logDebug(format string, v ...interface{}) {
	if *debug {
		log.Printf(format, v...)
	}
}

func invalidCommand() {
	fmt.Println("Invalid command. Available commands: create, " +
			"read, update, delete.")
		fmt.Println("Flags:")
		flag.PrintDefaults()
		os.Exit(1)
}

func main() {
	flag.Parse()
	logDebug("Debug: true")

	if len(flag.Args()) == 0 {
		invalidCommand()
	}

	if len(*repo) == 0 {
		fmt.Println("Please specify repo arg for access github.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if len(*token) == 0 {
		*token = os.Getenv("GH_TOKEN")
	}

	if len(*token) == 0 {
		fmt.Println("Please specify token arg or GH_TOKEN env value for " +
		"access github.")
		flag.PrintDefaults()
		os.Exit(1)
	}
	
	logDebug("command: %s", flag.Args()[0])
	commandArgs := flag.Args()[1:]
	switch flag.Args()[0] {
	case "create":
		create(commandArgs)
	case "read":
		read(commandArgs)
	case "update":
		update(commandArgs)
	case "delete":
		delete(commandArgs)
	default:
		invalidCommand()
	}
}

func create(args []string) {
	logDebug("create arg: %s", args)
	flagSet := flag.NewFlagSet("create", flag.ExitOnError)

	var title string
	flagSet.StringVar(&title, "title", "", "title of new issue")

	err := flagSet.Parse(args)

	if err != nil {
		flagSet.PrintDefaults()
	}
	logDebug("title: %s", title)

	flagSet.PrintDefaults()
}

func read(args []string) {
	logDebug("read arg: %s", args)
}

func update(args []string) {
	logDebug("update arg: %s", args)
}

func delete(args []string) {
	logDebug("delete arg: %s", args)
}
