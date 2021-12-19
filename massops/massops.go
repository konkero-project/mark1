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

func Add(r *gitops.Repo) {
	allRepos = append(allRepos, r)
}

func CloneAll(ctx context.Context) {
	for _, v := range allRepos {
		addJob(cloneOp, v)
	}
}

func PullAll(ctx context.Context) {
	for _, v := range allRepos {
		addJob(pullOp, v)
	}
}

func PullSelected(ctx context.Context) {
	for _, v := range selected {
		addJob(pullOp, v)
	}
}
