package config

import (
	"encoding/json"
	"fmt"

	"github.com/sieniven/polygoncdk-eigenda/etherman"
)

func (cfg *Config) loadNetworkConfig() {
	networkJSON := L1NetworkConfigJSON
	config, err := LoadGenesisFromJSONString(networkJSON)
	if err != nil {
		panic(fmt.Errorf("failed to load genesis configuration from file. Error: %v", err))
	}
	cfg.L1Config = config
}

func LoadGenesisFromJSONString(jsonStr string) (etherman.L1Config, error) {
	var cfg etherman.L1Config
	if err := json.Unmarshal([]byte(jsonStr), &cfg); err != nil {
		return etherman.L1Config{}, nil
	}
	return cfg, nil
}
