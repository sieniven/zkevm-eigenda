package dataavailability

import (
	"context"
	"encoding/binary"

	"github.com/Layr-Labs/eigenda/clients"
)

type DataAvailabilityProvider struct {
	client clients.DisperserClient
	cfg    clients.Config
}

// Factory method for a new DataAvailibilityProvider instance
func New(cfg Config) *DataAvailabilityProvider {
	c := clients.Config{
		Hostname:          cfg.Hostname,
		Port:              cfg.Port,
		Timeout:           cfg.Timeout.Duration,
		UseSecureGrpcFlag: cfg.UseSecureGrpcFlag,
	}
	s := MockBlobRequestSigner{}

	p := &DataAvailabilityProvider{
		client: clients.NewDisperserClient(&c, s),
		cfg:    c,
	}
	return p
}

func (d DataAvailabilityProvider) PostSequence(ctx context.Context, batchesData [][]byte) ([]byte, error) {
	return []byte{}, nil
}

func (d DataAvailabilityProvider) GetSequence(ctx context.Context, requestID []byte) ([][]byte, error) {
	return [][]byte{}, nil
}

func EncodeSequence(batchesData [][]byte) []byte {
	sequence := []byte{}
	metadata := []byte{}
	n := uint64(len(batchesData))
	bn := make([]byte, 8)
	binary.BigEndian.PutUint64(bn, n)
	metadata = append(metadata, bn...)

	for _, seq := range batchesData {
		// Add batch data to byte array
		sequence = append(sequence, seq...)

		// Add batch metadata to byte array
		n := uint64(len(seq))
		bn := make([]byte, 8)
		binary.BigEndian.PutUint64(bn, n)
		metadata = append(metadata, bn...)
	}
	sequence = append(metadata, sequence...)
	return sequence
}

func DecodeSequence(blobData []byte) [][]byte {
	bn := blobData[:8]
	n := binary.BigEndian.Uint64(bn)
	metadata := blobData[8 : 8*(n+1)]
	sequence := blobData[8*(n+1):]

	batchesData := [][]byte{}
	idx := uint64(0)
	for i := uint64(0); i < n; i++ {
		// Get batch data byte array length
		bn := metadata[8*i : 8*i+8]
		n := binary.BigEndian.Uint64(bn)
		batchesData = append(batchesData, sequence[idx:idx+n])
		idx += n
	}
	return batchesData
}
