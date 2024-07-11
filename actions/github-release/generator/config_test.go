package generator_test

import (
	"os"
	"testing"

	"github.com/sethvargo/go-githubactions"
	"github.com/stretchr/testify/assert"
	"github.com/thoughtgears/github-actions/actions/github-release/generator"
)

func TestNewFromInputs(t *testing.T) {
	envMap := map[string]string{
		"INPUT_GITHUB_TOKEN": "gh_token",
		"INPUT_RELEASE_NAME": "Action 1234 v1.0.0",
		"INPUT_DESCRIPTION":  "Release description, body of the release",
		"INPUT_LATEST":       "false",
		"INPUT_PRERELEASE":   "true",
		"GITHUB_REF":         "refs/tags/v1.0.0",
		"GITHUB_REPOSITORY":  "octocat/hello-world",
	}

	setEnv(envMap)
	defer unsetEnv(envMap)

	action := githubactions.New()

	inputs, err := generator.NewFromInputs(action)
	assert.NoError(t, err)
	assert.Equal(t, envMap["INPUT_GITHUB_TOKEN"], inputs.GithubToken)
	assert.Equal(t, envMap["INPUT_RELEASE_NAME"], inputs.ReleaseName)
	assert.Equal(t, envMap["INPUT_DESCRIPTION"], inputs.Body)
	assert.Equal(t, envMap["INPUT_LATEST"], inputs.Latest)
	assert.True(t, inputs.PreRelease)
	assert.Equal(t, "v1.0.0", inputs.Version)
	assert.Equal(t, "octocat", inputs.Owner)
	assert.Equal(t, "hello-world", inputs.Repo)

}

func TestNewFromInputsDefaults(t *testing.T) {
	envMap := map[string]string{
		"INPUT_github_token": "gh_token",
		"INPUT_release_name": "Action 1234 v1.0.0",
		"INPUT_description":  "Release description, body of the release",
		"GITHUB_REF":         "refs/tags/v1.0.0",
		"GITHUB_REPOSITORY":  "octocat/hello-world",
	}

	setEnv(envMap)
	defer unsetEnv(envMap)

	action := githubactions.New()

	inputs, err := generator.NewFromInputs(action)
	assert.NoError(t, err)
	assert.Equal(t, envMap["INPUTS_github_token"], inputs.GithubToken)
	assert.Equal(t, envMap["INPUTS_release_name"], inputs.ReleaseName)
	assert.Equal(t, envMap["INPUTS_description"], inputs.ReleaseName)
	assert.Equal(t, "true", inputs.Latest)
	assert.False(t, inputs.PreRelease)
	assert.Equal(t, "v1.0.0", inputs.Version)
	assert.Equal(t, "octocat", inputs.Owner)
	assert.Equal(t, "hello-world", inputs.Repo)

}

func setEnv(input map[string]string) {
	for key, value := range input {
		os.Setenv(key, value)
	}
}

func unsetEnv(input map[string]string) {
	for key := range input {
		os.Unsetenv(key)
	}
}
