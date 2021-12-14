package main

import (
	"context"
	"os"

	"konkero-project/mark1/config"
	"konkero-project/mark1/org"
	"konkero-project/mark1/ui/repocard"

	"fyne.io/fyne/v2/app"
	"github.com/grafov/kiwi"
)

func main() {
	log := kiwi.New()
	kiwi.SinkTo(os.Stdout, kiwi.AsLogfmt()).Start()

	kiwi.Log("repos", config.Repos, "access", config.Access) // XXX
	o := org.Init(log, config.Repos.Org, config.Access.GithubToken)
	ctx := context.Background()
	repos, _ := o.RepoList(ctx)

	a := app.New()
	w := a.NewWindow("Repositories info")

	for _, v := range repos {
		card := repocard.Add(v.Name, v.Desc)
		card.MakeCard()
	}
	w.SetContent(repocard.MakeGrid())
	w.ShowAndRun()
}
