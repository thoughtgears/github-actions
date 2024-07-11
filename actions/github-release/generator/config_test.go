package generator_test

import (
	"os"
	"testing"

	"github.com/sethvargo/go-githubactions"
	"github.com/stretchr/testify/assert"
	"github.com/thoughtgears/github-actions/actions/github-release/generator"
)

func TestNewFromInputs(t *testing.T) {
	action := githubactions.New()

	timeInput := map[string]string{
		"ANSIC":       "Mon Jan _2 15:04:05 2006",
		"UnixDate":    "Mon Jan _2 15:04:05 MST 2006",
		"RubyDate":    "Mon Jan 02 15:04:05 -0700 2006",
		"RFC822":      "02 Jan 06 15:04 MST",
		"RFC822Z":     "02 Jan 06 15:04 -0700",
		"RFC850":      "Monday, 02-Jan-06 15:04:05 MST",
		"RFC1123":     "Mon, 02 Jan 2006 15:04:05 MST",
		"RFC1123Z":    "Mon, 02 Jan 2006 15:04:05 -0700",
		"RFC3339":     "2006-01-02T15:04:05Z07:00",
		"RFC3339Nano": "2006-01-02T15:04:05.999999999Z07:00",
		"Kitchen":     "3:04PM",
	}

	for name, format := range timeInput {
		os.Setenv("INPUT_TIME_FORMAT", name)
		inputs := generator.NewFromInputs(action)
		assert.Equal(t, format, inputs.TimeFormat)
		os.Unsetenv("INPUT_TIME_FORMAT")
	}
}
