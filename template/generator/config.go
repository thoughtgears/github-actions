package generator

import (
	"time"

	"github.com/sethvargo/go-githubactions"
)

var timeFormats = map[string]string{
	"ANSIC":       time.ANSIC,
	"UnixDate":    time.UnixDate,
	"RubyDate":    time.RubyDate,
	"RFC822":      time.RFC822,
	"RFC822Z":     time.RFC822Z,
	"RFC850":      time.RFC850,
	"RFC1123":     time.RFC1123,
	"RFC1123Z":    time.RFC1123Z,
	"RFC3339":     time.RFC3339,
	"RFC3339Nano": time.RFC3339Nano,
	"Kitchen":     time.Kitchen,
}

type Inputs struct {
	TimeFormat string
}

func NewFromInputs(action *githubactions.Action) *Inputs {
	timeFormatInput := action.GetInput("time_format")
	return &Inputs{
		TimeFormat: timeFormats[timeFormatInput],
	}
}
