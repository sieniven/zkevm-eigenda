package sequencesender

import (
	"bufio"
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/sieniven/zkevm-eigenda/etherman"
	"github.com/sieniven/zkevm-eigenda/etherman/types"
	"github.com/sieniven/zkevm-eigenda/ethtxmanager"
	"github.com/sieniven/zkevm-eigenda/log"
	batchTypes "github.com/sieniven/zkevm-eigenda/sequencesender/types"
)

const (
	ethTxManagerOwner                = "sequencer"
	monitoredIDFormat                = "sequence-from-%v-to-%v"
	sendSequnceFlagTriggerBufferSize = 5
)

type SequenceSender struct {
	cfg              Config
	ethTxManager     *ethtxmanager.Client
	etherman         *etherman.Client
	sendSequenceFlag atomic.Bool
	lastBatchNum     uint64

	// data availability layer
	da dataAbilitier
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

// SetDataProvider sets the data provider
func (s *SequenceSender) SetDataProvider(da dataAbilitier) {
	s.da = da
}

func (s *SequenceSender) SendSequenceHandle(ctx context.Context, reader *bufio.Reader) {
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println(err)
		} else if char == 's' {
			s.sendSequenceFlag.Store(true)
		} else {
			fmt.Println("unknown command received, skippping")
		}
		time.Sleep(time.Second)
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

		sequences, err := s.getMockSequencesToSend()
		if err != nil || len(sequences) == 0 {
			if err != nil {
				fmt.Printf("error getting sequences: %v\n", err)
			} else {
				fmt.Println("waiting for sequences to be worth sending to L1")
			}
			time.Sleep(s.cfg.WaitPeriodSendSequence.Duration)
			return
		}

		sequenceCount := len(sequences)
		firstSequence := sequences[0]
		lastSequence := sequences[sequenceCount-1]
		fmt.Printf("sending sequences to L1. From batch %d to batch %d\n", firstSequence.BatchNumber, lastSequence.BatchNumber)

		// Add sequence to be monitored
		daMessage, err := s.da.PostSequence(ctx, sequences)
		if err != nil {
			fmt.Printf("error posting sequences to the data availability protocol: %v\n", err)
			return
		}
		to, data, err := s.etherman.BuildMockSequenceBatchesTxData(
			s.cfg.SenderAddress, sequences, uint64(lastSequence.LastL2BLockTimestamp), firstSequence.BatchNumber-1, s.cfg.L2Coinbase, daMessage)
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
// We generate mock batches with max batch data bytes size and convert them into sequences.
func (s *SequenceSender) getMockSequencesToSend() ([]types.Sequence, error) {
	// Generate mock batch data for max configured size
	data := make([]byte, s.cfg.MaxBatchBytesSize)
	for i := uint64(0); i < s.cfg.MaxBatchBytesSize; i++ {
		data[i] = byte(10)
	}

	sequences := []types.Sequence{}

	// Add sequences until too big for a single L1 tx or last batch is reached
	for {
		// Create mock batch
		batch := batchTypes.Batch{
			BatchNumber: uint64(s.lastBatchNum),
			Coinbase:    s.cfg.L2Coinbase,
			BatchL2Data: data,
			Timestamp:   time.Now(),
		}
		s.lastBatchNum += 1

		seq := types.Sequence{
			BatchL2Data:          batch.BatchL2Data,
			BatchNumber:          batch.BatchNumber,
			LastL2BLockTimestamp: batch.Timestamp.Unix(),
		}
		sequences = append(sequences, seq)
		if len(sequences) == int(s.cfg.MaxBatchesForL1) {
			log.Infof(
				"sequence should be sent to L1, because MaxBatchesForL1 (%d) has been reached",
				s.cfg.MaxBatchesForL1,
			)
			break
		}
	}
	return sequences, nil
}
