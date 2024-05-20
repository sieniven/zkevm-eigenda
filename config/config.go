package config

import (
	"bytes"
	"path/filepath"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/sieniven/polygoncdk-eigenda/etherman"
	"github.com/sieniven/polygoncdk-eigenda/ethtxmanager"
	"github.com/sieniven/polygoncdk-eigenda/sequencesender"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
	"google.golang.org/appengine/log"
)

// Represents the configuration of the entire mock Polygon CDK Node
// The file is [TOML format]
// You could find some examples:
// - `config/environments/local/local.node.config.toml`: running a permisionless node
// - `config/environments/mainnet/node.config.toml`
// - `config/environments/public/node.config.toml`
// - `test/config/test.node.config.toml`: configuration for a trusted node used in CI
//
// [TOML format]: https://en.wikipedia.org/wiki/TOML
type Config struct {
	Etherman       etherman.Config
	EthTxManager   ethtxmanager.Config
	SequenceSender sequencesender.Config
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
	configFilePath := ctx.String("cfg")
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
			log.Infof("config file not found")
		} else {
			log.Infof("error reading config file: ", err)
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

	return cfg, nil
}
