package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func createRequest(method, url, token string) (*http.Request, error) {

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	return req, nil
}

func GetIssuesByRepo(token string, repo string) (*IssuesListResult, error) {

	url := fmt.Sprintf(IssuesURLbyRepo, repo)

	logDebug(url)

	req, err := createRequest("GET", url, token)

	if err != nil {
		logDebug("%v", err)
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list query failed: %s", resp.Status)
	}

	logDebug("%v", resp.Status)

	var result IssuesListResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
