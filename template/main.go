package main

import (
	"context"

	"github.com/sethvargo/go-githubactions"
	"github.com/thoughtgears/github-actions/template/generator"
)

func run() error {
	ctx := context.Background()
	action := githubactions.New()
	inputs := generator.NewFromInputs(action)
	return Run(ctx, inputs)
}

func main() {
	if err := run(); err != nil {
		githubactions.Fatalf("%v", err)
	}
}
