package config

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"gopkg.in/yaml.v3"
)

const configFileName string = ".todocli.yaml"

type directoryConfig struct {
	ForbidenDirectories []string
}

type patternConfig struct {
	CustomRegex string
	Keyword string
	UrgencySuffix string
	CommentType string
}

type Config struct {
	Directory *directoryConfig
	Pattern *patternConfig
	ParsingRegex *regexp.Regexp
}

// default parsing regex "^\\s*// (TODO)(O*): ([a-zA-Z ]*)$"
func (config *Config) constructParsingRegex() {
	config.ParsingRegex = regexp.MustCompile(fmt.Sprintf(
		"^\\s*%s (%s)(%s*): ([a-zA-Z ]*)$",
		config.Pattern.CommentType,
		config.Pattern.Keyword,
		config.Pattern.UrgencySuffix,
	))
}

func (config *Config) setDefaults() {
	directoryConfig := &directoryConfig {
		ForbidenDirectories: []string{"node_modules", ".git"},
	}

	patternConfig := &patternConfig{
		CustomRegex: "",
		Keyword: "TODO",
		UrgencySuffix: "O",
		CommentType: "//",
	}

	config.Directory = directoryConfig
	config.Pattern = patternConfig

	config.constructParsingRegex()
}

func GetConfig() *Config {
	config, err := ioutil.ReadFile(configFileName)
	var data Config
	data.setDefaults()

	if err != nil {
		fmt.Printf("[INFO] Custom config file '%s' does not exist in the current working direcotry, proceeding with the default one.\n", configFileName)
		return &data
	}

	err = yaml.Unmarshal(config, &data)

	if err != nil {
		fmt.Printf("[WARN] Error parsing custom config file '%s', proceeding with the default one.\n", configFileName)
		return &data
	}

	data.constructParsingRegex()

	return &data
}
