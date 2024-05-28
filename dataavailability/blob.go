package dataavailability

import (
	"fmt"

	disperser_rpc "github.com/Layr-Labs/eigenda/api/grpc/disperser"
	"github.com/ethereum/go-ethereum/common"
)

type BlobVerificationProof struct {
	BatchId        uint32
	BlobIndex      uint32
	BatchMetadata  BatchMetadata
	InclusionProof []byte
	QuorumIndices  []byte
}

type BatchMetadata struct {
	// The header of the data store
	BatchHeader BatchHeader
	// The hash of the signatory record
	SignatoryRecordHash common.Hash
	// The block number at which the batch was confirmed
	ConfirmationBlockNumber uint32
}

type BatchHeader struct {
	BlobHeadersRoot common.Hash
	// Each byte is a different quorum number
	QuorumNumbers []byte
	// Every bytes is an amount less than 100 specifying the percentage of stake
	// that has signed in the corresponding quorum in `quorumNumbers`
	SignedStakeForQuorums []byte
	ReferenceBlockNumber  uint32
}

func GetVerificationProof(proof *disperser_rpc.BlobVerificationProof) (BlobVerificationProof, error) {
	if len(proof.BatchMetadata.BatchHeader.BatchRoot) != 32 {
		return BlobVerificationProof{}, fmt.Errorf("BlobHeadersRoot not type bytes32")
	}

	if len(proof.BatchMetadata.SignatoryRecordHash) != 32 {
		return BlobVerificationProof{}, fmt.Errorf("SignatoryRecordHash not type bytes32")
	}

	return BlobVerificationProof{
		BatchId:   proof.BatchId,
		BlobIndex: proof.BlobIndex,
		BatchMetadata: BatchMetadata{
			BatchHeader: BatchHeader{
				BlobHeadersRoot:       common.BytesToHash(proof.BatchMetadata.BatchHeader.BatchRoot),
				QuorumNumbers:         proof.BatchMetadata.BatchHeader.QuorumNumbers,
				SignedStakeForQuorums: proof.BatchMetadata.BatchHeader.QuorumSignedPercentages,
				ReferenceBlockNumber:  proof.BatchMetadata.BatchHeader.ReferenceBlockNumber,
			},
			SignatoryRecordHash:     common.BytesToHash(proof.BatchMetadata.SignatoryRecordHash),
			ConfirmationBlockNumber: proof.BatchMetadata.ConfirmationBlockNumber,
		},
		InclusionProof: proof.InclusionProof,
		QuorumIndices:  proof.QuorumIndexes,
	}, nil
}

// Get abi-encoded Keccak-256 hash of the reduced batch header
func (proof BlobVerificationProof) GetBatchHeaderHash() []byte {
	return nil
}

// Fallible conversion method if blob info is empty.
func TryToDataAvailabilityMessage(proof BlobVerificationProof) ([]byte, error) {
	return nil, nil
}

// Fallible conversion method if data availability message encoding is incorrect.
func TryFromDataAvailabilityMessage(msg []byte) (BlobVerificationProof, error) {
	return BlobVerificationProof{}, nil
}
