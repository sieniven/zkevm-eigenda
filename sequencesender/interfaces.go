package sequencesender

import (
	"context"

	ethmanTypes "github.com/sieniven/zkevm-eigenda/etherman/types"
)

type dataAbilitier interface {
	PostSequence(ctx context.Context, sequences []ethmanTypes.Sequence) ([]byte, error)
}
