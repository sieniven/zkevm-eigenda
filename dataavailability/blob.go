package dataavailability

import (
	"bytes"
	"fmt"

	disperser_rpc "github.com/Layr-Labs/eigenda/api/grpc/disperser"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type BlobData struct {
	BlobHeader            BlobHeader
	BlobVerificationProof BlobVerificationProof
	BatchHeaderHash       []byte
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
	X common.Hash
	Y common.Hash
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

func GetBlobData(info *disperser_rpc.BlobInfo) (BlobData, error) {
	header := GetBlobHeader(info.GetBlobHeader())
	proof, err := GetBlobVerificationProof(info.GetBlobVerificationProof())
	if err != nil {
		return BlobData{}, err
	}
	return BlobData{
		BlobHeader:            header,
		BlobVerificationProof: proof,
		BatchHeaderHash:       info.GetBlobVerificationProof().GetBatchMetadata().GetBatchHeaderHash(),
	}, nil
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
			X: common.BytesToHash(header.GetCommitment().GetX()),
			Y: common.BytesToHash(header.GetCommitment().GetY()),
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

func TryEncodeToDataAvailabilityMessage(blobData BlobData) ([]byte, error) {
	parsedABI, err := abi.JSON(bytes.NewReader([]byte(blobDataABI)))
	if err != nil {
		return nil, err
	}

	// Encode the data
	encoded, err := parsedABI.Pack("decodeBlobData", blobData)
	if err != nil {
		return nil, err
	}

	return encoded, nil
}

func TryDecodeFromDataAvailabilityMessage(msg []byte) (BlobData, error) {
	var blobData BlobData

	// Parse the ABI
	parsedABI, err := abi.JSON(bytes.NewReader([]byte(blobDataABI)))
	if err != nil {
		return blobData, err
	}

	// Decode the data
	err = parsedABI.UnpackIntoInterface(&blobData, "decodeBlobData", msg)
	if err != nil {
		return blobData, err
	}

	return blobData, nil
}

func (blobData BlobData) DebugLogBlobData() {
	fmt.Println("Logging blob data...")

	fmt.Println("--- Blob header ---")
	fmt.Println("Blob header commitment x: ", blobData.BlobHeader.Commitment.X.Bytes())
	fmt.Println("Blob header commitment y: ", blobData.BlobHeader.Commitment.Y.Bytes())
	fmt.Println("Blob header data length: ", blobData.BlobHeader.DataLength)
	for idx, q := range blobData.BlobHeader.QuorumBlobParams {
		fmt.Printf("Blob header quorum %v quorum number: %v\n", idx, q.QuorumNumber)
		fmt.Printf("Blob header quorum %v quorum adversary threshold percentage: %v\n", idx, q.AdversaryThresholdPercentage)
		fmt.Printf("Blob header quorum %v quorum confirmation threshold percentage: %v\n", idx, q.ConfirmationThresholdPercentage)
		fmt.Printf("Blob header quorum %v quorum chunk length: %v\n", idx, q.ChunkLength)
	}

	fmt.Println("--- Blob verification proof ---")
	fmt.Println("Blob verification proof batch id: ", blobData.BlobVerificationProof.BatchId)
	fmt.Println("Blob verification proof blob idx: ", blobData.BlobVerificationProof.BlobIndex)

	fmt.Println("Blob verification proof batch metadata:")
	fmt.Println("Blob verification proof batch metadata batch header batch headers root: ", blobData.BlobVerificationProof.BatchMetadata.BatchHeader.BlobHeadersRoot.Bytes())
	fmt.Println("Blob verification proof batch metadata batch header quorum numbers: ", blobData.BlobVerificationProof.BatchMetadata.BatchHeader.QuorumNumbers)
	fmt.Println("Blob verification proof batch metadata batch header signed stake for quorums: ", blobData.BlobVerificationProof.BatchMetadata.BatchHeader.SignedStakeForQuorums)
	fmt.Println("Blob verification proof batch metadata batch header reference block number: ", blobData.BlobVerificationProof.BatchMetadata.BatchHeader.ReferenceBlockNumber)
	fmt.Println("Blob verification proof batch metadata signature record hash: ", blobData.BlobVerificationProof.BatchMetadata.SignatoryRecordHash.Bytes())
	fmt.Println("Blob verification proof batch metadata confirmation block number: ", blobData.BlobVerificationProof.BatchMetadata.ConfirmationBlockNumber)

	fmt.Println("Blob verification proof inclusion proof: ", blobData.BlobVerificationProof.InclusionProof)
	fmt.Println("Blob verification proof quorum indices: ", blobData.BlobVerificationProof.QuorumIndices)

	fmt.Println("--- Batch header hash ---")
	fmt.Println("Blob data batch header hash: ", blobData.BatchHeaderHash)
}

func DebugLogBlobInfoResponse(info *disperser_rpc.BlobInfo) {
	fmt.Println("Logging blob data...")

	fmt.Println("--- Blob header ---")
	fmt.Println("Blob header commitment x: ", info.BlobHeader.Commitment.X)
	fmt.Println("Blob header commitment y: ", info.BlobHeader.Commitment.Y)
	fmt.Println("Blob header data length: ", info.BlobHeader.DataLength)
	for idx, q := range info.BlobHeader.BlobQuorumParams {
		fmt.Printf("Blob header quorum %v quorum number: %v\n", idx, q.QuorumNumber)
		fmt.Printf("Blob header quorum %v quorum adversary threshold percentage: %v\n", idx, q.AdversaryThresholdPercentage)
		fmt.Printf("Blob header quorum %v quorum confirmation threshold percentage: %v\n", idx, q.ConfirmationThresholdPercentage)
		fmt.Printf("Blob header quorum %v quorum chunk length: %v\n", idx, q.ChunkLength)
	}

	fmt.Println("--- Blob verification proof ---")
	fmt.Println("Blob verification proof batch id: ", info.BlobVerificationProof.BatchId)
	fmt.Println("Blob verification proof blob idx: ", info.BlobVerificationProof.BlobIndex)

	fmt.Println("Blob verification proof batch metadata:")
	fmt.Println("Blob verification proof batch metadata batch header batch headers root: ", info.BlobVerificationProof.BatchMetadata.BatchHeader.BatchRoot)
	fmt.Println("Blob verification proof batch metadata batch header quorum numbers: ", info.BlobVerificationProof.BatchMetadata.BatchHeader.QuorumNumbers)
	fmt.Println("Blob verification proof batch metadata batch header signed stake for quorums: ", info.BlobVerificationProof.BatchMetadata.BatchHeader.QuorumSignedPercentages)
	fmt.Println("Blob verification proof batch metadata batch header reference block number: ", info.BlobVerificationProof.BatchMetadata.BatchHeader.ReferenceBlockNumber)
	fmt.Println("Blob verification proof batch metadata signature record hash: ", info.BlobVerificationProof.BatchMetadata.SignatoryRecordHash)
	fmt.Println("Blob verification proof batch metadata confirmation block number: ", info.BlobVerificationProof.BatchMetadata.ConfirmationBlockNumber)

	fmt.Println("Blob verification proof inclusion proof: ", info.BlobVerificationProof.InclusionProof)
	fmt.Println("Blob verification proof quorum indices: ", info.BlobVerificationProof.QuorumIndexes)

	fmt.Println("--- Batch header hash ---")
	fmt.Println("Blob verification proof batch metadata batch header hash: ", info.BlobVerificationProof.BatchMetadata.BatchHeaderHash)
}
