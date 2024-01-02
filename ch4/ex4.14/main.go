/*
Exercise 4.14: Create a web server that queries GitHub once and then allows
navigation of the list of bug reports, milestones, and users.
*/

package main

import (
	"flag"
	"fmt"
	"gopl.io/ch4/ex4.14/github"
	"html/template"
	"log"
	"net/http"
	"os"
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

func processFlags() {
	flag.Parse()

	if *debug {
		github.EnableDebug()
	}

	logDebug("Debug: true")

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
}

func main() {
	processFlags()

	// Issues
	issuesTemplateText, err := os.ReadFile("issues.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	issuesTemplate, err := template.New("issues").
		Parse(string(issuesTemplateText))
	if err != nil {
		log.Fatal(err)
	}

	issues, err := github.GetIssuesByRepo(*token, *repo)
	if err != nil {
		log.Fatal(err)
	}

	issuesHandler := func(w http.ResponseWriter, req *http.Request) {
		err := issuesTemplate.Execute(w, issues)
		if err != nil {
			log.Printf("error on issues %s", err)
		}
	}

	http.HandleFunc("/", issuesHandler)
	http.HandleFunc("/issues", issuesHandler)

	// Users
	usersTemplateText, err := os.ReadFile("users.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	usersTemplate, err := template.New("users").
		Parse(string(usersTemplateText))
	if err != nil {
		log.Fatal(err)
	}

	users, err := github.GetCollaborators(*token, *repo)
	if err != nil {
		log.Fatal(err)
	}

	userHandler := func(w http.ResponseWriter, req *http.Request) {
		err := usersTemplate.Execute(w, users)
		if err != nil {
			log.Printf("error on users %s", err)
		}
	}
	http.HandleFunc("/users", userHandler)

	// Milestones
	milestonesTemplateText, err := os.ReadFile("milestones.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	milestonesTemplate, err := template.New("milestones").
		Parse(string(milestonesTemplateText))
	if err != nil {
		log.Fatal(err)
	}

	milestones, err := github.GetMilestones(*token, *repo)
	if err != nil {
		log.Fatal(err)
	}

	milestonesHandler := func(w http.ResponseWriter, req *http.Request) {
		err := milestonesTemplate.Execute(w, milestones)
		if err != nil {
			log.Printf("error on milestones %s", err)
		}
	}
	http.HandleFunc("/milestones", milestonesHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
