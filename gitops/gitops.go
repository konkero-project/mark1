// Package gitops wraps operations for a single GIT repository.
package gitops

import (
	"context"
	"os"

	"konkero-project/mark1/config"
	"konkero-project/mark1/org"

	"github.com/go-git/go-git/v5"
)

type Repo struct {
	Name     string
	Desc     string
	CloneURL string
}

func Init(r org.RepoInfo) *Repo {
	return &Repo{
		Name:     r.Name,
		Desc:     r.Desc,
		CloneURL: r.CloneURL,
	}
}

func (r *Repo) Clone(ctx context.Context) error {
	_, err := git.PlainClone(config.Local.Basepath, false, &git.CloneOptions{
		URL:      r.CloneURL,
		Progress: os.Stdout,
	})
	return err
}
