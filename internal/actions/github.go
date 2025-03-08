package actions

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/google/go-github/v55/github"
	"golang.org/x/oauth2"
)

func getGitHubClient(authenticated bool) *github.Client {
	ctx := context.Background()
	token := os.Getenv("GITHUB_TOKEN")

	if authenticated && token != "" {
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		return github.NewClient(oauth2.NewClient(ctx, ts))
	}
	return github.NewClient(&http.Client{})
}

// TODO: SEPARATE VALIDATION

func parseActionInput(action string) (owner, repo, ref string, err error) {
	parts := strings.Split(action, "@")
	if len(parts) != 2 {
		return "", "", "", fmt.Errorf("invalid action format: expected owner/repo@version")
	}

	ownerRepo := strings.Split(parts[0], "/")
	if len(ownerRepo) != 2 {
		return "", "", "", fmt.Errorf("invalid action format: expected owner/repo")
	}

	return ownerRepo[0], ownerRepo[1], parts[1], nil
}

// TODO: SEPARATE VALIDATION

func fetchActionFile(client *github.Client, ctx context.Context, owner, repo, ref string) (*github.RepositoryContent, error) {
	fileContent, _, _, err := client.Repositories.GetContents(ctx, owner, repo, "action.yml", &github.RepositoryContentGetOptions{Ref: ref})
	return fileContent, err
}

func FetchGitHubAction(action string) (*github.RepositoryContent, error) {
	ctx := context.Background()
	owner, repo, ref, err := parseActionInput(action)
	if err != nil {
		return nil, err
	}

	client := getGitHubClient(false)
	content, err := fetchActionFile(client, ctx, owner, repo, ref)

	if err != nil {
		token := os.Getenv("GITHUB_TOKEN")
		if token == "" {
			return nil, fmt.Errorf("rate limit exceeded or private repo; set GITHUB_TOKEN")
		}

		client = getGitHubClient(true)
		contentWithAuth, errWithAuth := fetchActionFile(client, ctx, owner, repo, ref)
		if errWithAuth != nil {
			return nil, fmt.Errorf("authenticated fetch failed: %v", errWithAuth)
		}
		return contentWithAuth, nil
	}

	return content, nil
}

// TODO: client.Repositories is wrong here
func FetchAllActions(repo string) ([]*github.RepositoryContent, error) {
	ctx := context.Background()

	parts := strings.Split(repo, "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid repository format: expects owner/repo")
	}
	owner, repository := parts[0], parts[1]

	client := getGitHubClient(false)

	_, contents, _, err := client.Repositories.GetContents(ctx, owner, repository, "/", nil)
	if err != nil {
		token := os.Getenv("GITHUB_TOKEN")
		if token == "" {
			return nil, fmt.Errorf("rate limit exceeded or private repo; set GITHUB_TOKEN")
		}

		client = getGitHubClient(true)
		var authErr error
		_, contents, _, authErr = client.Repositories.GetContents(ctx, owner, repository, "/", nil)
		if authErr != nil {
			return nil, fmt.Errorf("authenticated fetch failed: %v", authErr)
		}
	}

	return filterActionFiles(contents), nil
}

// May remove
func filterActionFiles(contents []*github.RepositoryContent) []*github.RepositoryContent {
	var actionFiles []*github.RepositoryContent
	for _, content := range contents {
		if *content.Type == "file" && (*content.Name == "action.yml" || *content.Name == "action.yaml") {
			actionFiles = append(actionFiles, content)
		}
	}
	return actionFiles
}
