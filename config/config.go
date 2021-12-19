package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config sections. Don't forget to add new sections to init()!
var (
	Access access
	Remote remote
	Local  local
)

// Access credentials for github.com.
type access struct {
	GithubToken string `split_words:"true"`
}

// Settings for remote service (github.com).
type remote struct {
	Org      string
	Excluded []string
}

// Settings for local filesystem.
type local struct {
	Basepath string
}

// Implicitly init all the sections of the config. Don't forget add a
// new sections here!
func init() {
	if err := envconfig.Process("", &Access); err != nil {
		log.Fatal("fatal config error for 'access' section", err.Error())
	}
	if err := envconfig.Process("", &Remote); err != nil {
		log.Fatal("fatal config error for 'remote' section", err.Error())
	}
	if err := envconfig.Process("", &Local); err != nil {
		log.Fatal("fatal config error for 'local' section", err.Error())
	}
}
