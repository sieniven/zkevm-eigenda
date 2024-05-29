package dataavailability

import (
	"fmt"
	"math/big"
	"strings"

	disperser_rpc "github.com/Layr-Labs/eigenda/api/grpc/disperser"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type BlobData struct {
	BlobHeader            BlobHeader
	BlobVerificationProof BlobVerificationProof
}

type BlobHeader struct {
	Commitment       Commitment
	DataLength       uint32
	QuorumBlobParams []QuorumBlobParam
}

type BlobVerificationProof struct {
	BatchId        uint32
	BlobIndex      uint32
	BatchMetadata  BatchMetadata
	InclusionProof []byte
	QuorumIndices  []byte
}

type QuorumBlobParam struct {
	QuorumNumber                    uint8
	AdversaryThresholdPercentage    uint8
	ConfirmationThresholdPercentage uint8
	ChunkLength                     uint32
}

type Commitment struct {
	X *big.Int
	Y *big.Int
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

type ReducedBatchHeader struct {
	BlobHeadersRoot      common.Hash
	ReferenceBlockNumber uint32
}

func GetBlobData(info *disperser_rpc.BlobInfo) (BlobData, error) {
	header := GetBlobHeader(info.GetBlobHeader())
	proof, err := GetBlobVerificationProof(info.GetBlobVerificationProof())
	if err != nil {
		return BlobData{}, nil
	}
	return BlobData{BlobHeader: header, BlobVerificationProof: proof}, nil
}

func GetBlobHeader(header *disperser_rpc.BlobHeader) BlobHeader {
	quorums := []QuorumBlobParam{}
	for _, quorum := range header.GetBlobQuorumParams() {
		q := QuorumBlobParam{
			QuorumNumber:                    uint8(quorum.GetQuorumNumber()),
			AdversaryThresholdPercentage:    uint8(quorum.GetAdversaryThresholdPercentage()),
			ConfirmationThresholdPercentage: uint8(quorum.GetConfirmationThresholdPercentage()),
			ChunkLength:                     quorum.GetChunkLength(),
		}
		quorums = append(quorums, q)
	}

	return BlobHeader{
		Commitment: Commitment{
			X: new(big.Int).SetBytes(header.GetCommitment().GetX()),
			Y: new(big.Int).SetBytes(header.GetCommitment().GetY()),
		},
		DataLength:       header.GetDataLength(),
		QuorumBlobParams: quorums,
	}
}

func GetBlobVerificationProof(proof *disperser_rpc.BlobVerificationProof) (BlobVerificationProof, error) {
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
func (proof BlobVerificationProof) GetBatchHeaderHash() ([]byte, error) {
	headerABI, err := abi.JSON(strings.NewReader(reducedBlockHeaderAbiJSON))
	if err != nil {
		fmt.Println("failed to parse ABI: ", err)
		return nil, err
	}

	header := ReducedBatchHeader{
		BlobHeadersRoot:      proof.BatchMetadata.BatchHeader.BlobHeadersRoot,
		ReferenceBlockNumber: proof.BatchMetadata.BatchHeader.ReferenceBlockNumber,
	}

	data, err := headerABI.Pack(
		"",
		header,
	)
	if err != nil {
		fmt.Println("failed to pack reduced block header: ", err)
		return nil, err
	}
	return crypto.Keccak256(data), nil
}

// Fallible conversion method if blob info is empty.
func TryToDataAvailabilityMessage(blobData BlobData) ([]byte, error) {
	blobABI, err := abi.JSON(strings.NewReader(blobInfoAbiJSON))
	if err != nil {
		fmt.Println("failed to parse ABI: ", err)
		return nil, err
	}

	data, err := blobABI.Pack("", blobData)
	if err != nil {
		fmt.Println("failed to pack blob data: ", err)
		return nil, err
	}

	return data, nil
}

// Fallible conversion method if data availability message encoding is incorrect.
func TryFromDataAvailabilityMessage(msg []byte) (BlobData, error) {
	blobABI, err := abi.JSON(strings.NewReader(blobInfoAbiJSON))
	if err != nil {
		fmt.Println("failed to parse ABI: ", err)
		return BlobData{}, err
	}

	var blobData BlobData
	unpacked, err := blobABI.Constructor.Inputs.Unpack(msg)
	if err != nil {
		return BlobData{}, err
	}
	err = blobABI.Constructor.Inputs.Copy(&blobData, unpacked)
	if err != nil {
		return BlobData{}, err
	}

	return blobData, nil
}
