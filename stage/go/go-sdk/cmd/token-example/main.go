package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/octokit/go-sdk/pkg"
	"github.com/octokit/go-sdk/pkg/github/repos"
)

func main() {
	client, err := pkg.NewApiClient(
		pkg.WithUserAgent("my-user-agent"),
		pkg.WithRequestTimeout(500*time.Second),
		pkg.WithBaseUrl("https://api.github.com"),
		pkg.WithTokenAuthentication(os.Getenv("GITHUB_TOKEN")),
	)

	// equally valid:
	//client, err := pkg.NewApiClient()
	if err != nil {
		log.Fatalf("error creating client: %v", err)
	}

	issue, err := client.Repos().ByOwnerId("buckelij").ByRepoId("octokit-source-generator").Issues().ByIssue_number(1).Get(context.Background(), nil)
	if err != nil {
		log.Fatalf("error getting issue: %v\n", err)
	} else {
		log.Printf("issue: %v %v\n", *issue.GetBody(), labelsFromAdditionalData(issue.GetAdditionalData()))
	}
	newissue := repos.NewItemItemIssuesPostRequestBody()
	newissueBody := "hello"
	newissue.SetBody(&newissueBody)
	title := repos.NewItemItemIssuesPostRequestBody_IssuesPostRequestBody_title()
	title.SetString(&newissueBody)
	newissue.SetTitle(title)
	newissue.SetAdditionalData(map[string]any{"labels": []string{"bug"}})
	issuesCreator := client.Repos().ByOwnerId("buckelij").ByRepoId("octokit-source-generator").Issues()
	newIssueRes, err := issuesCreator.Post(context.Background(), newissue, nil)
	if err != nil || newIssueRes == nil {
		log.Fatalf("error creating issue: %v\n", err)
	} else {
		log.Printf("issue: %v %v\n", *newIssueRes.GetBody(), labelsFromAdditionalData(newIssueRes.GetAdditionalData()))
	}
}

func labelsFromAdditionalData(additionalData map[string]any) []string {
	if additionalData == nil {
		return []string{}
	}
	var labels []string
	if labelsData, ok := additionalData["labels"]; ok {
		labelsJson, err := json.Marshal(labelsData)
		if err != nil {
			fmt.Printf("error marshalling labels: %v\n", err)
		}
		var labelObjectSlice []struct{ Name string }
		var labelStringSlice []string
		json.Unmarshal(labelsJson, &labelObjectSlice)
		json.Unmarshal(labelsJson, &labelStringSlice)
		fmt.Printf("labelsFromAdditionalData %v %v\n", labelObjectSlice, labelStringSlice)
		switch {
		case labelObjectSlice != nil:
			for _, e := range labelObjectSlice {
				if e.Name != "" {
					labels = append(labels, e.Name)
				}
			}
		case labelStringSlice != nil:
			for _, e := range labelStringSlice {
				if e != "" {
					labels = append(labels, e)
				}
			}
		}
	}
	return labels
}
