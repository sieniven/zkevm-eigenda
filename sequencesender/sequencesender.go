package sequencesender

import (
	"context"

	"github.com/sieniven/polygoncdk-eigenda/etherman"
	"github.com/sieniven/polygoncdk-eigenda/ethtxmanager"
)

type SequenceSender struct {
	cfg          Config
	ethTxManager ethtxmanager.Client
	etherman     etherman.Client
}

// New inits sequence sender
func New(cfg Config, etherman etherman.Client, manager ethtxmanager.Client) (*SequenceSender, error) {
	return &SequenceSender{
		cfg:          cfg,
		etherman:     etherman,
		ethTxManager: manager,
	}, nil
}

func (s *SequenceSender) Start(ctx context.Context) {
	for {
		s.tryToSendSequence(ctx)
	}
}

func (s *SequenceSender) tryToSendSequence(ctx context.Context) {

}
