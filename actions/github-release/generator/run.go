package generator

import (
	"context"
	"time"

	"github.com/sethvargo/go-githubactions"
)

func Run(_ context.Context, inputs *Inputs) error {
	now := time.Now().Format(inputs.TimeFormat)
	githubactions.SetOutput("current_time", now)

	return nil
}
