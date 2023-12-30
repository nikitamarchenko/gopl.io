/*
Exercise 4.11: Build a tool that lets users create, read, update, and delete
GitHub issues from the command line, invoking their preferred text editor when
substantial text input is required.
*/

package main

import (
	"flag"
	"fmt"
	"gopl.io/ch4/ex4.11/github"
	"log"
	"os"
	"os/exec"
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

	if *debug {
		github.EnableDebug()
	}

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

	var body string
	flagSet.StringVar(&body, "body", "", "body of new issue")

	err := flagSet.Parse(args)

	if err != nil {
		flagSet.PrintDefaults()
		os.Exit(1)
	}

	if len(title) == 0 {
		fmt.Println("Please specify title. It requires for creating issue.")
		flagSet.PrintDefaults()
		os.Exit(1)
	}

	logDebug("title: %s", title)

	if len(body) == 0 {
		body = runEditor()
	}

	logDebug("body: %s", body)

	result, err := github.CreateIssue(*token, *repo, title, body)
	if err != nil {
		log.Fatal(err)
	}

	if result != nil {
		fmt.Printf("Created: #%-5d %6.20s %.55s\n",
			result.Number, result.User.Login, result.Title)
	}
}

func read(args []string) {
	logDebug("read arg: %s", args)
	result, err := github.GetIssuesByRepo(*token, *repo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", len(*result))
	for _, item := range *result {
		fmt.Printf("#%-5d %6.20s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

func issueState(open, close bool) *bool {
	
	if !close && !open {
		return nil
	}

	var state bool

	if open {
		state = true
	}

	return &state
}

func update(args []string) {
	logDebug("update arg: %s", args)

	flagSet := flag.NewFlagSet("update", flag.ExitOnError)

	var id int
	flagSet.IntVar(&id, "id", 0, "id of issue")

	var title string
	flagSet.StringVar(&title, "title", "", "new title")

	var body string
	flagSet.StringVar(&body, "body", "", "new body")

	var editor bool
	flagSet.BoolVar(&editor, "editor", false, "use editor for body")

	var close bool
	flagSet.BoolVar(&close, "close", false, "close issue")

	var open bool
	flagSet.BoolVar(&open, "open", false, "open issue")

	err := flagSet.Parse(args)

	if err != nil {
		flagSet.PrintDefaults()
		os.Exit(1)
	}

	if id == 0 {
		fmt.Println("Please specify issue id. It requires for updating issue.")
		flagSet.PrintDefaults()
		os.Exit(1)
	}

	if id < 0 {
		fmt.Println("Please specify VALID issue id. " +
			"It requires for updating issue.")
		flagSet.PrintDefaults()
		os.Exit(1)
	}

	logDebug("title: %s", title)

	if len(body) == 0 && editor {
		body = runEditor()
	}

	logDebug("body: %s", body)

	if close && open {
		fmt.Println("Please choose between close and open operation.")
		flagSet.PrintDefaults()
		os.Exit(1)
	}

	state := issueState(open, close)

	result, err := github.UpdateIssue(*token, *repo, id, title, body, state)
	if err != nil {
		log.Fatal(err)
	}

	if result != nil {
		fmt.Printf("Updated: #%-5d %6.20s %.55s\n",
			result.Number, result.User.Login, result.Title)
	}
}

func delete(args []string) {
	logDebug("delete arg: %s", args)

	flagSet := flag.NewFlagSet("delete", flag.ExitOnError)

	var id int
	flagSet.IntVar(&id, "id", 0, "id of issue")

	err := flagSet.Parse(args)

	if err != nil {
		flagSet.PrintDefaults()
		os.Exit(1)
	}

	if id == 0 {
		fmt.Println("Please specify issue id. It requires for updating issue.")
		flagSet.PrintDefaults()
		os.Exit(1)
	}

	if id < 0 {
		fmt.Println("Please specify VALID issue id. " +
			"It requires for updating issue.")
		flagSet.PrintDefaults()
		os.Exit(1)
	}

	err = github.DeleteIssue(*token, *repo, id)
	if err != nil {
		log.Fatal(err)
	}
}

func runEditor() string {
	editor := os.Getenv("EDITOR")
	if len(editor) == 0 {
		editor = "nano"
	}

	file, err := os.CreateTemp("", "create-body-*")

	if err != nil {
		fmt.Println("Error: can't create tmp file.")
		os.Exit(1)
	}

	cmd := exec.Command(editor, file.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	logDebug("Running command and waiting for it to finish...")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error: can't run editor.")
		os.Exit(1)
	}

	text, err := os.ReadFile(file.Name())

	if err != nil {
		fmt.Printf("Error: can't read file %s.\n", file.Name())
		os.Exit(1)
	}

	logDebug("Command finished with error: %v", err)
	return string(text)
}
