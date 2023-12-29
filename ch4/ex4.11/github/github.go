// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 110.
//!+

// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github

import (
	"log"
	"time"
)

const IssuesURLbyRepo = "https://api.github.com/repos/%s/issues"

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
