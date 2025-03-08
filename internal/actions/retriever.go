package actions

import (
	"fmt"
)

func RetrieveAction(action string) (*GithubAction, error) {
	actionFile, err := FetchGitHubAction(action)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch action.yml: %v", err)
	}

	yamlContent, err := actionFile.GetContent()
	if err != nil {
		return nil, fmt.Errorf("failed to decode actin.yml content: %v", err)
	}

	actionData, err := ParseActionYML(yamlContent)
	if err != nil {
		return nil, fmt.Errorf("failed to parse action.yml: %v", err)
	}

	return actionData, nil
}
func RetrieveAllActions(repo string) ([]*GithubAction, error) {
	actionFiles, err := FetchAllActions(repo)
	if err != nil {
		return nil, fmt.Errorf("error fetching all actions: %v", err)
	}

	var actions []*GithubAction
	for _, actionFile := range actionFiles {

		// kludge
		yamlContent, err := actionFile.GetContent()
		if err != nil {
			fmt.Printf("Skipping %s due to decoding error: %v\n", *actionFile.Path, err)
			continue
		}

		actionData, err := ParseActionYML(yamlContent)
		if err != nil {
			fmt.Printf("Skipping %s due to YAML parsing error: %v\n", *actionFile.Path, err)
			continue
		}

		actions = append(actions, actionData)
	}
	// kludge
	if len(actions) == 0 {
		return nil, fmt.Errorf("no valid GitHub Actions found in repository")
	}

	return actions, nil
}
