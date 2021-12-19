package massops

import (
	"context"

	"konkero-project/mark1/gitops"
)

// It defines the number on jobs that executed in the same time.
const parallelOps = 5

var jobs = make(chan job, parallelOps)

// Type for registered operations.
type operation int

// Enumeration of registered operations.
const (
	cloneOp operation = iota
	pullOp
)

type job struct {
	r  *gitops.Repo
	op operation
}

// HandleJobs handles jobs queue and runs appropriate workers up to
// parallelOps limit.
func HandleJobs(ctx context.Context) {
}

func addJob(op operation, repo *gitops.Repo) {

}

func doClone(ctx context.Context, repo *gitops.Repo) {
	repo.Clone(ctx)
}

func doPull(ctx context.Context, repo *gitops.Repo) {
	//	gitops.Pull(ctx, repo.CloneURL)
}
