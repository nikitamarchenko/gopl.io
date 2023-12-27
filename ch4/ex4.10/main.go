/*
Exercise 4.10: Modify issues to report the results in age categories, say less
than a month old, less than a year old, and more than a year old.
*/

package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"gopl.io/ch4/github"
)

type IssueByDate []*github.Issue

func (a IssueByDate) Len() int      { return len(a) }
func (a IssueByDate) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a IssueByDate) Less(i, j int) bool {
	return a[i].CreatedAt.After(a[j].CreatedAt)
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	var lessMonth IssueByDate
	var lessYear IssueByDate
	var moreYear IssueByDate

	now := time.Now()

	lessMonthDate := time.Date(now.Year(), now.Month()-1, now.Day(), now.Hour(),
		now.Minute(), 0, 0, now.Location())

	lessYearDate := time.Date(now.Year()-1, now.Month(), now.Day(), now.Hour(),
		now.Minute(), 0, 0, now.Location())

	for _, item := range result.Items {
		if item.CreatedAt.After(lessMonthDate) {
			lessMonth = append(lessMonth, item)
		} else if item.CreatedAt.After(lessYearDate) {
			lessYear = append(lessYear, item)
		} else {
			moreYear = append(moreYear, item)
		}
	}

	sort.Sort(lessMonth)
	sort.Sort(lessYear)
	sort.Sort(moreYear)

	printGithubItems("Less than month:", lessMonth)
	printGithubItems("Less than year:", lessYear)
	printGithubItems("More than year:", moreYear)
}

func printGithubItems(title string, list []*github.Issue) {
	fmt.Printf("%s\n", title)
	for _, item := range list {
		fmt.Printf("#%-5d %s %9.9s %.55s\n",
			item.Number, item.CreatedAt.Format(time.RFC822),
			item.User.Login, item.Title)
	}
}
