package generator

import (
	"context"

	"github.com/google/go-github/v63/github"
)

func (i *Inputs) Run() error {
	ctx := context.TODO()
	if _, _, err := i.client.Repositories.CreateRelease(ctx, i.Owner, i.Repo, &github.RepositoryRelease{
		TagName:    &i.Version,
		Name:       &i.ReleaseName,
		Body:       &i.Body,
		MakeLatest: &i.Latest,
		Prerelease: &i.PreRelease,
	}); err != nil {
		return err
	}
	return nil
}
