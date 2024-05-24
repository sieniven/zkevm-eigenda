package dataavailability

import "encoding/binary"

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
