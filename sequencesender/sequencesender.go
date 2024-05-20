package sequencesender

import (
	"context"
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"

	"github.com/sieniven/polygoncdk-eigenda/etherman"
	"github.com/sieniven/polygoncdk-eigenda/etherman/types"
	"github.com/sieniven/polygoncdk-eigenda/ethtxmanager"
)

const (
	ethTxManagerOwner = "sequencer"
	monitoredIDFormat = "sequence-from-%v-to-%v"
)

type SequenceSender struct {
	cfg              Config
	ethTxManager     *ethtxmanager.Client
	etherman         *etherman.Client
	sendSequenceFlag atomic.Bool
	lastBatchNum     uint64

	// data availability layer
	da DataAvaibilityProvider
}

// New inits sequence sender
func New(cfg Config, etherman *etherman.Client, manager *ethtxmanager.Client) (*SequenceSender, error) {
	s := SequenceSender{
		cfg:          cfg,
		etherman:     etherman,
		ethTxManager: manager,
		lastBatchNum: 0,
	}
	s.sendSequenceFlag.Store(false)

	return &s, nil
}

func (s *SequenceSender) SendSequenceHandle(ctx context.Context) {
	for {
		// TODO: run a client directly that listens for flag switches
	}
}

func (s *SequenceSender) Start(ctx context.Context) {
	for {
		s.tryToSendSequence(ctx)
	}
}

func (s *SequenceSender) tryToSendSequence(ctx context.Context) {
	retry := false
	// process monitored sequences before starting a next cycle
	s.ethTxManager.ProcessPendingMonitoredTxs(ctx, ethTxManagerOwner, func(result ethtxmanager.MonitoredTxResult) {
		if result.Status == ethtxmanager.MonitoredTxStatusFailed {
			retry = true
			fmt.Println("failed to send sequence")
		}
	})

	if retry {
		return
	}

	// Check if should send mock sequence to L1
	if s.sendSequenceFlag.Load() {
		fmt.Println("getting sequences to send")
		s.sendSequenceFlag.Store(false)

		numSequences := rand.Intn(10)
		sequences, err := s.getMockSequencesToSend(numSequences)
		if err != nil || len(sequences) == 0 {
			if err != nil {
				fmt.Printf("error getting sequences: %v\n", err)
			} else {
				fmt.Println("waiting for sequences to be worth sending to L1")
			}
			time.Sleep(s.cfg.WaitPeriodSendSequence)
			return
		}

		sequenceCount := len(sequences)
		fmt.Printf("sending sequences to L1. From batch %d to batch %d\n", sequences[0].BatchNumber, sequences[0].BatchNumber+uint64(sequenceCount))

		// Add sequence to be monitored
		firstSequence := sequences[0]
		lastSequence := sequences[sequenceCount-1]
		dataAvailabilityMessage, err := s.da.PostSequence(ctx, sequences)
		if err != nil {
			fmt.Printf("error posting sequences to the data availability protocol: %v\n", err)
			return
		}
		to, data, err := s.etherman.BuildMockSequenceBatchesTxData(
			s.cfg.SenderAddress, sequences, uint64(lastSequence.LastL2BLockTimestamp), firstSequence.BatchNumber-1, s.cfg.L2Coinbase, dataAvailabilityMessage)
		if err != nil {
			fmt.Printf("error estimating new sequenceBatches to add to eth tx manager: %v\n", err)
			return
		}

		monitoredTxID := fmt.Sprintf(monitoredIDFormat, firstSequence.BatchNumber, lastSequence.BatchNumber)
		err = s.ethTxManager.Add(ctx, ethTxManagerOwner, monitoredTxID, s.cfg.SenderAddress, to, nil, data, s.cfg.GasOffset)
		if err != nil {
			fmt.Printf("error to add sequences tx to eth tx manager: %v\n", err)
			return
		}
	} else {
		// No sequnce to send
		time.Sleep(time.Second)
	}
}

// getMockSequencesToSend is a mock function to replicate Polygon CDK's getSequencesToSend.
// We generate an array of mock sequences to be sent to the L1.
func (s *SequenceSender) getMockSequencesToSend(numSequences int) ([]types.Sequence, error) {
	sequences := []types.Sequence{}
	data := []byte("hihihihihihihihihihihihihihihihihihi")

	// Add sequences until too big for a single L1 tx or last batch is reached
	for i := 0; i < numSequences; i++ {
		// Create a mock sequence
		seq := types.Sequence{
			BatchL2Data:          data,
			BatchNumber:          s.lastBatchNum,
			LastL2BLockTimestamp: time.Now().Unix(),
		}
		s.lastBatchNum += 1
		sequences = append(sequences, seq)
		if len(sequences) == int(s.cfg.MaxBatchesForL1) {
			fmt.Printf(
				"sequence should be sent to L1, because MaxBatchesForL1 (%d) has been reached\n",
				s.cfg.MaxBatchesForL1,
			)
			return sequences, nil
		}
	}

	// Reach the latest batch. Decide if its worth it to send the sequence, or wait for new batches
	if len(sequences) == 0 {
		fmt.Println("no batches to be sequenced")
		return nil, nil
	}

	fmt.Println("sequences should be sent to L1, too long since didnt send anything to L1")
	return sequences, nil
}

// SetDataProvider sets the data provider
func (s *SequenceSender) SetDataProvider(da DataAvaibilityProvider) {
	s.da = da
}
