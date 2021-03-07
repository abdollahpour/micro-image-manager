package config

import (
	"io/ioutil"
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Configuration struct {
	Port    int32  `envconfig:"port" required:"true" default:"8080"`
	Host    string `envconfig:"host" required:"true" default:"0.0.0.0"`
	DistDir string `envconfig:"dist_dir" required:"true" default:"images"`
	TempDir string `envconfig:"temp_dir" required:"true"`
}

func NewEnvConfiguration() Configuration {
	var conf Configuration
	envconfig.Process("MIM", &conf)
	if len(conf.TempDir) == 0 {
		dir, err := ioutil.TempDir("", "*.html")
		if err != nil {
			log.Fatal(err)
		}
		conf.TempDir = dir
	}
	return conf
}
