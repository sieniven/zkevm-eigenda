package config

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
	"xlayerda/config"

	"github.com/mitchellh/mapstructure"
	"github.com/sieniven/polygoncdk-eigenda/etherman"
	"github.com/sieniven/polygoncdk-eigenda/ethtxmanager"
	"github.com/sieniven/polygoncdk-eigenda/sequencesender"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

const FlagCfg = "cfg"

// Represents the configuration of the entire mock Polygon CDK Node
// The file is [TOML format]
// Example config:
// - `config/environments/mock/node.config.toml`
//
// [TOML format]: https://en.wikipedia.org/wiki/TOML
type Config struct {
	Etherman       etherman.Config
	EthTxManager   ethtxmanager.Config
	SequenceSender sequencesender.Config
	L1Config       etherman.L1Config         `mapstructure:"L1Config"`
	Key            config.KeystoreFileConfig `mapstructure:"Key"`
}

// Default parses the default configuration values
func Default() (*Config, error) {
	var cfg Config
	viper.SetConfigType("toml")

	err := viper.ReadConfig(bytes.NewBuffer([]byte(DefaultValues)))
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&cfg, viper.DecodeHook(mapstructure.TextUnmarshallerHookFunc()))
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

// Load loads the configuration
func Load(ctx *cli.Context) (*Config, error) {
	cfg, err := Default()
	if err != nil {
		return nil, err
	}
	configFilePath := ctx.String(FlagCfg)
	if configFilePath != "" {
		dirName, fileName := filepath.Split(configFilePath)

		fileExtension := strings.TrimPrefix(filepath.Ext(fileName), ".")
		fileNameWithoutExtension := strings.TrimSuffix(fileName, "."+fileExtension)

		viper.AddConfigPath(dirName)
		viper.SetConfigName(fileNameWithoutExtension)
		viper.SetConfigType(fileExtension)
	}
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("ZKEVM_NODE")
	err = viper.ReadInConfig()
	if err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)
		if ok {
			fmt.Println("config file not found")
		} else {
			fmt.Println("error reading config file: ", err)
			return nil, err
		}
	}

	decodeHooks := []viper.DecoderConfigOption{
		// this allows arrays to be decoded from env var separated by ",", example: MY_VAR="value1,value2,value3"
		viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(mapstructure.TextUnmarshallerHookFunc(), mapstructure.StringToSliceHookFunc(","))),
	}

	err = viper.Unmarshal(&cfg, decodeHooks...)
	if err != nil {
		return nil, err
	}

	// Get L1Config parameters
	cfg.loadNetworkConfig()

	return cfg, nil
}
