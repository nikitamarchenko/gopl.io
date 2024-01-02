package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func createRequest(method, url, token, body string) (*http.Request, error) {

	req, err := http.NewRequest(method, url, strings.NewReader(body))

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

	req, err := createRequest("GET", url, token, "")

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

func CreateIssue(token, repo, title, body string) (*Issue, error) {

	url := fmt.Sprintf(IssuesURLbyRepo, repo)

	logDebug(url)

	issue := IssueCreateRequestPayload{Title: title, Body: body}

	payload, err := json.Marshal(issue)

	if err != nil {
		logDebug("Can't marshal payload %v with error %v", issue, err)
		return nil, err
	}

	logDebug("Payload %s", string(payload))

	req, err := createRequest("POST", url, token, string(payload))

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

	logDebug("%v", resp.Status)

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("create query failed: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil

}

func UpdateIssue(token string, repo string,
	id int, title string, body string, state *bool) (*Issue, error) {

	url := fmt.Sprintf(IssuesURLbyRepoAndId, repo, id)

	logDebug(url)

	issue := IssueUpdateRequestPayload{Title: title, Body: body}

	if state != nil {
		if *state {
			issue.State = "open"
		} else {
			issue.State = "closed"
		}
	}

	payload, err := json.Marshal(issue)

	if err != nil {
		logDebug("Can't marshal payload %v with error %v", issue, err)
		return nil, err
	}

	logDebug("Payload %s", string(payload))

	req, err := createRequest("PATCH", url, token, string(payload))

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

	logDebug("%v", resp.Status)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("create query failed: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func DeleteIssue(token string, repo string, id int) error {

	return fmt.Errorf("deleting issues is not supported by github API")

	// url := fmt.Sprintf(IssuesURLbyRepoAndId, repo, id)

	// logDebug(url)

	// req, err := createRequest("DELETE", url, token, "")

	// if err != nil {
	// 	logDebug("%v", err)
	// 	return err
	// }

	// client := &http.Client{}
	// resp, err := client.Do(req)

	// if err != nil {
	// 	return err
	// }

	// if resp != nil {
	// 	defer resp.Body.Close()
	// }

	// logDebug("%v", resp.Status)

	// if resp.StatusCode != http.StatusOK {
	// 	return fmt.Errorf("delete query failed: %s", resp.Status)
	// }

	// return nil
}

func GetCollaborators(token string, repo string) (
	*CollaboratorsListResult, error) {

	url := fmt.Sprintf(CollaboratorsURL, repo)

	logDebug(url)

	req, err := createRequest("GET", url, token, "")

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

	var result CollaboratorsListResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}


func GetMilestones(token string, repo string) (*MilestoneListResult, error) {

	url := fmt.Sprintf(MilestonesURL, repo)

	logDebug(url)

	req, err := createRequest("GET", url, token, "")

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

	var result MilestoneListResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
