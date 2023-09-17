package config

import (
	"encoding/json"
	"fmt"
	"os"

	cm "github.com/GideonLewis/crypto-analyzer/pkg/model"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type JsonConfig struct {
}

type EnvConfig struct {
	MySQL cm.MySQL
}

var (
	EnvCfg  *EnvConfig
	JsonCfg *JsonConfig
)

func GetEnvConfig() (*EnvConfig, error) {
	if EnvCfg == nil {
		envCfg := &EnvConfig{}
		err := godotenv.Load()
		if err != nil {
			log.Errorf(fmt.Sprintf("%s", err))
			return nil, err
		}

		err = envconfig.Process("", envCfg)
		if err != nil {
			log.Errorf(fmt.Sprintf("%s", err))
			return nil, err
		}

		EnvCfg = envCfg
	}

	return EnvCfg, nil
}

func GetJsonConfig() (*JsonConfig, error) {
	if JsonCfg == nil {
		jsonCfg := &JsonConfig{}
		configFile, err := os.Open("")
		if err != nil {
			log.Errorf(fmt.Sprintf("%s", err))
			return nil, err
		}
		defer configFile.Close()
		jsonParser := json.NewDecoder(configFile)
		jsonParser.Decode(&jsonCfg)

		JsonCfg = jsonCfg
	}

	return JsonCfg, nil
}
