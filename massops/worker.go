package massops

import (
	"context"

	"konkero-project/mark1/gitops"

	"github.com/grafov/kiwi"
)

// It defines the number on jobs that executed in the same time.
const parallelOps = 5

var jobs = make(chan job, parallelOps)

// Type for registered operations.
type operation func(context.Context, *gitops.Repo) error

// // Enumeration of registered operations.
// const (
//	cloneOp operation = iota
//	pullOp
// )

type job struct {
	r  *gitops.Repo
	op operation
}

// HandleJobs handles jobs queue and runs appropriate workers up to
// parallelOps limit.
func HandleJobs(ctx context.Context, log *kiwi.Logger) {
	// XXX
	for i := 0; i < parallelOps; i++ {
		go func(ctx context.Context, log *kiwi.Logger, id int) {
			l := log.Fork().With("logger_id", id)
			l.Log("msg", "worker started")
			for {
				select {
				case j := <-jobs:
					j.op(ctx, j.r)
				}
			}

		}(ctx, log, i)
	}

}

func addJob(op operation, repo *gitops.Repo) {

}

func doClone(ctx context.Context, repo *gitops.Repo) error {
	repo.Clone(ctx)
	return nil
}

func doPull(ctx context.Context, repo *gitops.Repo) error {
	//	gitops.Pull(ctx, repo.CloneURL)
	return nil
}
