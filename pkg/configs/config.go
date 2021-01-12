package configs

import "github.com/tkanos/gonfig"

type Configuration struct {
	Port             int
	ConnectionString string
}

func Init() (Configuration, error) {
	conf := Configuration{}
	err := gonfig.GetConf("./configs.json", &conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}
