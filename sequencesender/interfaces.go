package sequencesender

import (
	"context"

	"github.com/sieniven/zkevm-eigenda/dataavailability"
	ethmanTypes "github.com/sieniven/zkevm-eigenda/etherman/types"
)

type dataAbilitier interface {
	PostSequence(ctx context.Context, sequences []ethmanTypes.Sequence) (dataavailability.BlobInfo, error)
}
