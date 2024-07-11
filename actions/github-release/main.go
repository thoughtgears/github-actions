package main

import (
	"github.com/sethvargo/go-githubactions"
	"github.com/thoughtgears/github-actions/actions/github-release/generator"
)

func run() error {
	action := githubactions.New()
	config, err := generator.NewFromInputs(action)
	if err != nil {
		return err
	}
	return config.Run()
}

func main() {
	if err := run(); err != nil {
		githubactions.Fatalf("%v", err)
	}
}
