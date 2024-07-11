package main

import (
	"context"
	"time"

	"github.com/thoughtgears/github-actions/template/generator"

	"github.com/sethvargo/go-githubactions"
)

func Run(_ context.Context, inputs *generator.Inputs) error {
	now := time.Now().Format(inputs.TimeFormat)
	githubactions.SetOutput("current_time", now)

	return nil
}
