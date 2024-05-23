package dataavailability

import (
	"context"
	"encoding/binary"
	"fmt"

	"github.com/Layr-Labs/eigenda/clients"
	"github.com/Layr-Labs/eigenda/encoding/utils/codec"
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
	blobData := EncodeSequence(batchesData)

	// Blob serialization
	blobData = codec.ConvertByPaddingEmptyByte(blobData)

	// Send blob to EigenDA disperser
	_, idBytes, err := d.client.DisperseBlob(ctx, blobData, []uint8{})
	if err != nil {
		fmt.Println("failed to send blob to EigenDA disperser")
		return []byte{}, nil
	}

	fmt.Println("sent blob to EigenDA disperser")
	return idBytes, nil
}

func (d DataAvailabilityProvider) GetSequence(ctx context.Context, requestID []byte) ([][]byte, error) {
	return [][]byte{}, nil
}

// EncodeSequence is the helper function to encode sequence data into 1D byte array. The
// encoding scheme is ensured to be lossless.
//
// The first n+1 8-bytes of the blob contains the metadata of the batches data.
// The first 8-bytes stores the size of the sequence, and the next 8-bytes will store the
// byte array length of every batch data.
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

// DecodeSequence is the helper function to decode 1D byte array into sequence data. The
// encoding scheme is ensured to be lossless.
//
// When decoding the blob data, the first n+1 8-bytes of the blob contains the metadata of
// the batches data.
// The first 8-bytes stores the size of the sequence, and the next 8-bytes will store the
// byte array length of every batch data.
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
