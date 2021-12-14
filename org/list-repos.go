package org

import (
	"context"

	"github.com/google/go-github/v41/github"
)

type Repo struct {
	Name     string
	Desc     string
	CloneURL string
}

func (o *org) RepoList(ctx context.Context) ([]Repo, error) {
	var err error
	opt := &github.RepositoryListByOrgOptions{
		Type:        "private",
		ListOptions: github.ListOptions{PerPage: pageSize},
	}
	var (
		repos, page []*github.Repository
		resp        *github.Response
	)
	for {
		page, resp, err = o.c.Repositories.ListByOrg(ctx, o.name, opt)
		if _, ok := err.(*github.RateLimitError); ok {
			o.l.Log("err", "hit rate limit")
			return nil, err
		}
		if err != nil {
			return nil, err
		}
		repos = append(repos, page...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	var (
		list []Repo
		desc string
	)
	for _, v := range repos {
		if v.Description != nil {
			desc = *v.Description
		}
		list = append(list, Repo{*v.Name, desc, *v.CloneURL})
	}
	return list, err
}
