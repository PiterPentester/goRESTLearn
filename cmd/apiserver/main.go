package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/PiterPentester/goRESTLearn/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to server config file")
}

func main() {
	flag.Parse()

	config := apiserver.CreateConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}