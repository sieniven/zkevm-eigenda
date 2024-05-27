package dataavailability

import (
	"encoding/binary"
	"fmt"
)

type BlobInfo struct {
	BlobIndex       uint32
	BatchHeaderHash []byte
}

// Fallible conversion method if blob info is empty.
func TryToDataAvailabilityMessage(blobInfo BlobInfo) ([]byte, error) {
	if len(blobInfo.BatchHeaderHash) == 0 {
		return nil, fmt.Errorf("empty blob header hash")
	}
	// Encoding scheme is the blob index (4 bytes), followed by the batch header hash
	bn := make([]byte, 4)
	binary.BigEndian.PutUint32(bn, blobInfo.BlobIndex)
	data := append(bn, blobInfo.BatchHeaderHash...)
	return data, nil
}

// Fallible conversion method if data availability message encoding is incorrect.
func TryFromDataAvailabilityMessage(msg []byte) (BlobInfo, error) {
	if len(msg) < 5 {
		return BlobInfo{}, fmt.Errorf("failed to decode data availability message")
	}
	info := BlobInfo{}
	bn := msg[:4]
	info.BlobIndex = binary.BigEndian.Uint32(bn)
	info.BatchHeaderHash = msg[4:]
	return info, nil
}
