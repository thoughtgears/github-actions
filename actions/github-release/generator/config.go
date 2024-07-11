package generator

import (
	"strconv"
	"strings"

	"github.com/google/go-github/v63/github"
	"github.com/sethvargo/go-githubactions"
)

type Inputs struct {
	Version     string
	GithubToken string
	ReleaseName string
	Latest      string
	PreRelease  bool
	Body        string
	Owner       string
	Repo        string
	client      *github.Client
}

func NewFromInputs(action *githubactions.Action) (*Inputs, error) {
	var inputs Inputs

	inputs.Latest = action.GetInput("latest")
	if inputs.Latest == "" {
		inputs.Latest = "true"
	}

	preRelease := action.GetInput("prerelease")
	inputs.PreRelease, _ = strconv.ParseBool(preRelease)

	inputs.GithubToken = action.GetInput("github_token")
	inputs.ReleaseName = action.GetInput("release_name")
	inputs.Body = action.GetInput("description")
	inputs.client = github.NewClient(nil).WithAuthToken(inputs.GithubToken)

	ctx, err := action.Context()
	if err != nil {
		return nil, err
	}

	inputs.Version = strings.Split(ctx.Ref, "/")[2]
	inputs.Owner, inputs.Repo = ctx.Repo()

	return &inputs, nil
}
