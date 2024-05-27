package etherman

import "github.com/sieniven/zkevm-eigenda/log"

func init() {
	log.Init(log.Config{
		Level:   "debug",
		Outputs: []string{"stderr"},
	})
}
