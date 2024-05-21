package dataavailability

import "github.com/sieniven/zkevm-eigenda/config/types"

type Config struct {
	Hostname          string
	Port              string
	Timeout           types.Duration
	UseSecureGrpcFlag bool
}
