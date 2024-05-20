package etherman

import (
	"context"

	ethmanTypes "github.com/sieniven/polygoncdk-eigenda/etherman/types"
)

type DataAvaibilityProvider interface {
	PostSequence(ctx context.Context, sequences []ethmanTypes.Sequence) ([]byte, error)
}
