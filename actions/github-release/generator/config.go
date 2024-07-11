package generator

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/sethvargo/go-githubactions"
)

type Inputs struct {
	Version     string
	GithubToken string
	Action      string
	Latest      string
	PreRelease  bool
	Owner       string
	Repo        string
	Release     bool
	Body        string
	BasePath    string
}

func NewFromInputs(action *githubactions.Action) (*Inputs, error) {
	var inputs Inputs

	ctx, err := action.Context()
	if err != nil {
		return nil, err
	}

	inputs.Latest = action.GetInput("latest")
	if inputs.Latest == "" {
		inputs.Latest = "true"
	}

	preRelease := action.GetInput("prerelease")
	inputs.PreRelease, _ = strconv.ParseBool(preRelease)
	inputs.GithubToken = action.GetInput("github_token")
	inputs.Action = action.GetInput("action")
	inputs.Version = strings.Split(ctx.Ref, "/")[2]
	inputs.Owner, inputs.Repo = ctx.Repo()
	inputs.BasePath = ctx.Workspace

	return &inputs, nil
}

func (i *Inputs) EvaluateRelease() error {
	changelogPath := fmt.Sprintf("%s/actions/%s/CHANGELOG.md", i.BasePath, i.Action)
	changelogContent, err := os.ReadFile(changelogPath)
	if err != nil {
		return err
	}

	lines := strings.Split(string(changelogContent), "\n")
	re := regexp.MustCompile(`^## \[(v[0-9.]+)]`)
	for _, line := range lines {
		// Find the top ## header
		matches := re.FindStringSubmatch(line)
		if len(matches) > 0 {
			headerVersion := matches[1]
			// Compare with the version in the struct
			if strings.Contains(headerVersion, i.Version) {
				i.Release = true
				break
			} else {
				return fmt.Errorf("version mismatch: file version %s, struct version %s", headerVersion, i.Version)
			}
		}
	}

	return nil
}
