package org

import (
	"context"

	"github.com/google/go-github/v41/github"
	"github.com/grafov/kiwi"
	"golang.org/x/oauth2"
)

const pageSize = 30

type org struct {
	name string
	c    *github.Client
	l    *kiwi.Logger
}

func Init(log *kiwi.Logger, orgname, key string) *org {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: key},
	)
	tc := oauth2.NewClient(ctx, ts)
	return &org{name: orgname, c: github.NewClient(tc), l: log}
}
