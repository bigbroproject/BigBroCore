package models

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Protocol struct {
	Type string `yaml:"type"`
	Port int `yaml:"port"`
	Server string `yaml:"server"`
}

type Service struct{
	Name string `yaml:"name"`
	Protocols []Protocol `yaml:"protocols"`
}


type Config struct{
	Services []Service `yaml:"services"`
}

/**
	Get config from yaml file
 */
func ConfigFromFile(filePath string ) (*Config,error){
	dat, err := ioutil.ReadFile(filePath)
	if err!=nil {
		return nil, err
	}

	conf := Config{}
	err = yaml.Unmarshal([]byte(dat), &conf)
	if err!=nil {
		return nil, err
	}
	return &conf,nil
}
