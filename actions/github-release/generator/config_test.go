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
		"INPUT_ACTION":       "my-action",
		"INPUT_LATEST":       "false",
		"INPUT_PRERELEASE":   "true",
		"GITHUB_REF":         "refs/tags/v1.0.0",
		"GITHUB_REPOSITORY":  "octocat/hello-world",
		"GITHUB_WORKSPACE":   "../../../",
	}

	setEnv(envMap)
	defer unsetEnv(envMap)

	action := githubactions.New()

	inputs, err := generator.NewFromInputs(action)
	assert.NoError(t, err)
	assert.Equal(t, envMap["INPUT_GITHUB_TOKEN"], inputs.GithubToken)
	assert.Equal(t, envMap["INPUT_ACTION"], inputs.Action)
	assert.Equal(t, envMap["INPUT_LATEST"], inputs.Latest)
	assert.True(t, inputs.PreRelease)
	assert.Equal(t, "v1.0.0", inputs.Version)
	assert.Equal(t, "octocat", inputs.Owner)
	assert.Equal(t, "hello-world", inputs.Repo)
	assert.Equal(t, "../../../", inputs.BasePath)
}

func TestNewFromInputsDefaults(t *testing.T) {
	envMap := map[string]string{
		"INPUT_github_token": "gh_token",
		"INPUT_release_name": "Action 1234 v1.0.0",
		"GITHUB_REF":         "refs/tags/v1.0.0",
		"GITHUB_REPOSITORY":  "octocat/hello-world",
	}

	setEnv(envMap)
	defer unsetEnv(envMap)

	action := githubactions.New()

	inputs, err := generator.NewFromInputs(action)
	assert.NoError(t, err)
	assert.Equal(t, envMap["INPUTS_github_token"], inputs.GithubToken)
	assert.Equal(t, envMap["INPUT_ACTION"], inputs.Action)
	assert.Equal(t, "true", inputs.Latest)
	assert.False(t, inputs.PreRelease)
	assert.Equal(t, "v1.0.0", inputs.Version)
	assert.Equal(t, "octocat", inputs.Owner)
	assert.Equal(t, "hello-world", inputs.Repo)
}

func TestInputs_EvaluateRelease(t *testing.T) {
	inputs := &generator.Inputs{
		Version:  "v1",
		BasePath: "../../../",
		Action:   "github-release",
		Release:  false,
	}

	err := inputs.ChangeLog()
	assert.NoError(t, err)
	assert.True(t, inputs.Release)
	assert.NotEmpty(t, inputs.Body)
}

func TestInputs_EvaluateReleaseFalse(t *testing.T) {
	inputs := &generator.Inputs{
		Version:  "v3",
		BasePath: "../../../",
		Action:   "github-release",
		Release:  false,
	}

	err := inputs.ChangeLog()
	assert.Error(t, err)
	assert.Equal(t, "version mismatch: version v3 not found in changelog", err.Error())
	assert.False(t, inputs.Release)
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
