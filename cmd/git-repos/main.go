package main

import (
	"context"
	"os"

	"konkero-project/mark1/config"
	"konkero-project/mark1/gitops"
	"konkero-project/mark1/massops"
	"konkero-project/mark1/org"
	"konkero-project/mark1/ui/repocard"

	"fyne.io/fyne/v2/app"
	"github.com/grafov/kiwi"
)

func main() {
	log := kiwi.New()
	kiwi.SinkTo(os.Stdout, kiwi.AsLogfmt()).Start()
	kiwi.Log("remote", config.Remote, "local", config.Local)
	a := app.New()
	w := a.NewWindow("Repositories info")
	ctx := context.Background()
	makeRepoList(ctx, log)
	w.SetContent(repocard.MakeGrid())
	massops.HandleJobs(ctx, log)
	w.ShowAndRun()
}

func makeRepoList(ctx context.Context, log *kiwi.Logger) {
	o := org.Init(log, config.Remote.Org, config.Access.GithubToken)
	remoteRepos, _ := o.RepoList(ctx)
next:
	for _, info := range remoteRepos {
		for _, n := range config.Remote.Excluded {
			if info.Name == n {
				continue next
			}
		}
		card := repocard.Add(info.Name, info.Desc)
		card.MakeCard()
		repo := gitops.Init(info)
		massops.Add(repo)
	}
}
