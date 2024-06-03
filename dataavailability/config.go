package dataavailability

import (
	"github.com/sieniven/zkevm-eigenda/config/types"
)

type Config struct {
	Hostname                 string         `mapstructure:"Hostname"`
	Port                     string         `mapstructure:"Port"`
	Timeout                  types.Duration `mapstructure:"Timeout"`
	UseSecureGrpcFlag        bool           `mapstructure:"UseSecureGrpcFlag"`
	RetrieveBlobStatusPeriod types.Duration `mapstructure:"RetrieveBlobStatusPeriod"`
}
