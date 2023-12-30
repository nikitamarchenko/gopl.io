package github

import (
	"log"
	"time"
)

const IssuesURLbyRepo = "https://api.github.com/repos/%s/issues"

const IssuesURLbyRepoAndId = "https://api.github.com/repos/%s/issues/%d"

type IssuesListResult []*Issue

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type IssueCreateRequestPayload struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type IssueUpdateRequestPayload struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
	State string `json:"state,omitempty"`
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

var debug bool

func EnableDebug() {
	debug = true
}

func logDebug(format string, v ...interface{}) {
	if debug {
		log.Printf(format, v...)
	}
}
