package config

import (
	"github.com/BurntSushi/toml"
)

const confDir = "./config/env/"

var conf Config

type BasicConfig struct {
	Port string `toml:"port"`
}

type AwsConfig struct {
	Region      string `toml:"region"`
	BucketName  string `toml:"bucket_name"`
}

type BasicAuthConfig struct {
	UserName string `toml:"username"`
	Password string `toml:"password"`
}

type Config struct {
	Basic BasicConfig `toml:"basic"`
	Aws AwsConfig `toml:"aws"`
	BasicAuth BasicAuthConfig `toml:"auth"`
}

func NewConfig(env string) (Config, error) {

	confPath := confDir + env + ".toml"
	if _, err := toml.DecodeFile(confPath, &conf); err != nil {
		return conf, err
	}
	return conf, nil
}

func GetConfig() Config {
	return conf
}


