package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v63/github"
	"github.com/sethvargo/go-githubactions"
	"github.com/thoughtgears/github-actions/actions/github-release/generator"
)

func release(inputs *generator.Inputs) error {
	ctx := context.TODO()
	client := github.NewClient(nil).WithAuthToken(inputs.GithubToken)

	if _, _, err := client.Repositories.CreateRelease(ctx, inputs.Owner, inputs.Repo, &github.RepositoryRelease{
		TagName:    github.String(inputs.Version),
		Name:       github.String(fmt.Sprintf("%s, %s", inputs.Action, inputs.Version)),
		Body:       github.String(inputs.Body),
		MakeLatest: github.String(inputs.Latest),
		Prerelease: github.Bool(inputs.PreRelease),
	}); err != nil {
		return err
	}
	return nil
}

func run() error {
	action := githubactions.New()
	config, err := generator.NewFromInputs(action)
	if err != nil {
		return err
	}

	if config.Release {
		return release(config)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		githubactions.Fatalf("%v", err)
	}
}
