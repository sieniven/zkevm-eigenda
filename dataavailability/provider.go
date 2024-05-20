package dataavailability

import (
	"context"

	"github.com/sieniven/polygoncdk-eigenda/etherman/types"
)

type DataAvailabilityProvider struct {
}

func (d DataAvailabilityProvider) PostSequence(ctx context.Context, sequences []types.Sequence) ([]byte, error) {
	return []byte{}, nil
}
