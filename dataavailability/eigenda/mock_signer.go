package eigenda

import "github.com/Layr-Labs/eigenda/core"

// Mock EigenLayer signer for PoC purposes. On production, there will be
// a need to register a BlobRequestSigner with EigenLayer.
type MockBlobRequestSigner struct{}

func (s MockBlobRequestSigner) SignBlobRequest(header core.BlobAuthHeader) ([]byte, error) {
	return []byte{}, nil
}

func (s MockBlobRequestSigner) GetAccountID() string {
	return ""
}
