package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetIssuesByRepo(token string, repo string) (*IssuesListResult, error) {

	url := fmt.Sprintf(IssuesURLbyRepo, repo)

	logDebug(url)

	resp, err := http.Get(url)

	if err != nil {
		logDebug("%v", err)
		return nil, err
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	logDebug("%v", resp.Status)

	var result IssuesListResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

