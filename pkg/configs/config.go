package configs

import "github.com/tkanos/gonfig"

type Configuration struct {
	Port             string `json:"port"`
	ConnectionString string `json:"connectionString"`
	DSN              string `json:"dsn"`
}

func Init() (Configuration, error) {
	conf := Configuration{}
	err := gonfig.GetConf("./configs.json", &conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}
