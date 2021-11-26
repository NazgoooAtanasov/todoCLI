package config

import(
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type DirectoryConfig struct {
	ForbidenDirectories []string
}

type PatternConfig struct {
	CustomRegex string
	Keyword string
	UrgencySuffix string
	CommentType string
}

type Config struct {
	Directory DirectoryConfig
	Pattern PatternConfig
}

func GetConfig() *Config {
	config, err := ioutil.ReadFile(".todocli.yaml")

	if err != nil {
		panic(err)
	}

	var data Config

	err = yaml.Unmarshal(config, &data)

	if err != nil {
		panic(err)
	}

	return &data
}
