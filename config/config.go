package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

var (
	Access access
	Repos  repos
)

type access struct {
	GithubToken string `split_words:"true"`
}

type repos struct {
	Org           string
	ExcludedRepos []string `split_words:"true"`
}

func init() {
	if err := envconfig.Process("", &Access); err != nil {
		log.Fatal("fatal config error for 'access' section", err.Error())
	}
	if err := envconfig.Process("", &Repos); err != nil {
		log.Fatal("fatal config error for 'repos' section", err.Error())
	}

}
