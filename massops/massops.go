// Package massops keeps repositories list and does operations on a
// selected set of repositories.
package massops

import (
	"context"

	"konkero-project/mark1/gitops"
)

var (
	allRepos []*gitops.Repo
	selected []*gitops.Repo
)

// Add adds repo to the all-repos list. Thread unsafe operation.
func Add(r *gitops.Repo) {
	allRepos = append(allRepos, r)
}

// Select adds marked repo to the selectable repos list. Thread unsafe
// operation.
func Select(r *gitops.Repo) {
	selected = append(selected, r)
}

func CloneAll(ctx context.Context) {
	for _, v := range allRepos {
		addJob(doClone, v)
	}
}

func PullAll(ctx context.Context) {
	for _, v := range allRepos {
		addJob(doPull, v)
	}
}

func PullSelected(ctx context.Context) {
	for _, v := range selected {
		addJob(doPull, v)
	}
}
