package dataavailability

import (
	"encoding/binary"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestEncodeBlobData(t *testing.T) {
	data := BlobData{
		BlobHeader: BlobHeader{
			Commitment: Commitment{
				X: common.BytesToHash(big.NewInt(12345).Bytes()),
				Y: common.BytesToHash(big.NewInt(67890).Bytes()),
			},
			DataLength: 100,
			QuorumBlobParams: []QuorumBlobParam{
				{
					QuorumNumber:                    1,
					AdversaryThresholdPercentage:    50,
					ConfirmationThresholdPercentage: 75,
					ChunkLength:                     1024,
				},
			},
		},
		BlobVerificationProof: BlobVerificationProof{
			BatchId:   1,
			BlobIndex: 2,
			BatchMetadata: BatchMetadata{
				BatchHeader: BatchHeader{
					BlobHeadersRoot:       [32]byte{},
					QuorumNumbers:         []byte{1, 2, 3},
					SignedStakeForQuorums: []byte{50, 60, 70},
					ReferenceBlockNumber:  12345,
				},
				SignatoryRecordHash:     [32]byte{},
				ConfirmationBlockNumber: 54321,
			},
			InclusionProof: []byte{0x01, 0x02, 0x03},
			QuorumIndices:  []byte{0x04, 0x05, 0x06},
		},
	}
	msg := EncodeToDataAvailabilityMessage(data)
	assert.NotNil(t, msg)
	assert.NotEmpty(t, msg)
}

func TestEncodeDecodeBlobData(t *testing.T) {
	data := BlobData{
		BlobHeader: BlobHeader{
			Commitment: Commitment{
				X: common.BytesToHash(big.NewInt(12345).Bytes()),
				Y: common.BytesToHash(big.NewInt(67890).Bytes()),
			},
			DataLength: 100,
			QuorumBlobParams: []QuorumBlobParam{
				{
					QuorumNumber:                    1,
					AdversaryThresholdPercentage:    50,
					ConfirmationThresholdPercentage: 75,
					ChunkLength:                     1024,
				},
			},
		},
		BlobVerificationProof: BlobVerificationProof{
			BatchId:   1,
			BlobIndex: 2,
			BatchMetadata: BatchMetadata{
				BatchHeader: BatchHeader{
					BlobHeadersRoot:       [32]byte{},
					QuorumNumbers:         []byte{1, 2, 3},
					SignedStakeForQuorums: []byte{50, 60, 70},
					ReferenceBlockNumber:  12345,
				},
				SignatoryRecordHash:     [32]byte{},
				ConfirmationBlockNumber: 54321,
			},
			InclusionProof: []byte{0x01, 0x02, 0x03},
			QuorumIndices:  []byte{0x04, 0x05, 0x06},
		},
	}
	msg := EncodeToDataAvailabilityMessage(data)
	assert.NotNil(t, msg)
	assert.NotEmpty(t, msg)

	// Check blob header
	decoded_data := DecodeFromDataAvailabilityMessage(msg)
	assert.Equal(t, data.BlobHeader.Commitment.X, decoded_data.BlobHeader.Commitment.X)
	assert.Equal(t, data.BlobHeader.Commitment.Y, decoded_data.BlobHeader.Commitment.Y)
	assert.Equal(t, data.BlobHeader.DataLength, decoded_data.BlobHeader.DataLength)
	for idx, q := range data.BlobHeader.QuorumBlobParams {
		assert.Equal(t, q.QuorumNumber, decoded_data.BlobHeader.QuorumBlobParams[idx].QuorumNumber)
		assert.Equal(t, q.AdversaryThresholdPercentage, decoded_data.BlobHeader.QuorumBlobParams[idx].AdversaryThresholdPercentage)
		assert.Equal(t, q.ConfirmationThresholdPercentage, decoded_data.BlobHeader.QuorumBlobParams[idx].ConfirmationThresholdPercentage)
		assert.Equal(t, q.ChunkLength, decoded_data.BlobHeader.QuorumBlobParams[idx].ChunkLength)
	}

	// Check blob verification proof
	assert.Equal(t, data.BlobVerificationProof.BatchId, decoded_data.BlobVerificationProof.BatchId)
	assert.Equal(t, data.BlobVerificationProof.BlobIndex, decoded_data.BlobVerificationProof.BlobIndex)
	assert.Equal(t, data.BlobVerificationProof.BatchMetadata, decoded_data.BlobVerificationProof.BatchMetadata)
	assert.Equal(t, data.BlobVerificationProof.BatchMetadata.BatchHeader.BlobHeadersRoot, decoded_data.BlobVerificationProof.BatchMetadata.BatchHeader.BlobHeadersRoot)
	assert.Equal(t, data.BlobVerificationProof.BatchMetadata.BatchHeader.QuorumNumbers, decoded_data.BlobVerificationProof.BatchMetadata.BatchHeader.QuorumNumbers)
	assert.Equal(t, data.BlobVerificationProof.BatchMetadata.BatchHeader.SignedStakeForQuorums, decoded_data.BlobVerificationProof.BatchMetadata.BatchHeader.SignedStakeForQuorums)
	assert.Equal(t, data.BlobVerificationProof.BatchMetadata.BatchHeader.ReferenceBlockNumber, decoded_data.BlobVerificationProof.BatchMetadata.BatchHeader.ReferenceBlockNumber)
	assert.Equal(t, data.BlobVerificationProof.BatchMetadata.SignatoryRecordHash, decoded_data.BlobVerificationProof.BatchMetadata.SignatoryRecordHash)
	assert.Equal(t, data.BlobVerificationProof.BatchMetadata.ConfirmationBlockNumber, decoded_data.BlobVerificationProof.BatchMetadata.ConfirmationBlockNumber)
	assert.Equal(t, data.BlobVerificationProof.InclusionProof, decoded_data.BlobVerificationProof.InclusionProof)
	assert.Equal(t, data.BlobVerificationProof.QuorumIndices, decoded_data.BlobVerificationProof.QuorumIndices)
}

func TestBlockHeaderHash(t *testing.T) {
	data := BlobData{
		BlobHeader: BlobHeader{
			Commitment: Commitment{
				X: common.BytesToHash(big.NewInt(12345).Bytes()),
				Y: common.BytesToHash(big.NewInt(67890).Bytes()),
			},
			DataLength: 100,
			QuorumBlobParams: []QuorumBlobParam{
				{
					QuorumNumber:                    1,
					AdversaryThresholdPercentage:    50,
					ConfirmationThresholdPercentage: 75,
					ChunkLength:                     1024,
				},
			},
		},
		BlobVerificationProof: BlobVerificationProof{
			BatchId:   1,
			BlobIndex: 2,
			BatchMetadata: BatchMetadata{
				BatchHeader: BatchHeader{
					BlobHeadersRoot:       [32]byte{},
					QuorumNumbers:         []byte{1, 2, 3},
					SignedStakeForQuorums: []byte{50, 60, 70},
					ReferenceBlockNumber:  12345,
				},
				SignatoryRecordHash:     [32]byte{},
				ConfirmationBlockNumber: 54321,
			},
			InclusionProof: []byte{0x01, 0x02, 0x03},
			QuorumIndices:  []byte{0x04, 0x05, 0x06},
		},
	}
	hash := data.BlobVerificationProof.GetBatchHeaderHash()
	assert.NotNil(t, hash)
	assert.NotEmpty(t, hash)

	// Manually add the bytes and see if they match
	b := data.BlobVerificationProof.BatchMetadata.BatchHeader.BlobHeadersRoot.Bytes()
	bn := make([]byte, 4)
	binary.BigEndian.PutUint32(bn, data.BlobVerificationProof.BatchMetadata.BatchHeader.ReferenceBlockNumber)
	b = append(b, bn...)
	calc_hash := crypto.Keccak256(b)
	assert.Equal(t, calc_hash, hash)
}
