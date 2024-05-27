package types

const (
	// FORKID_BLUEBERRY is the fork id 4
	FORKID_BLUEBERRY = 4
	// FORKID_DRAGONFRUIT is the fork id 5
	FORKID_DRAGONFRUIT = 5
	// FORKID_INCABERRY is the fork id 6
	FORKID_INCABERRY = 6
	// FORKID_ETROG is the fork id 7
	FORKID_ETROG = 7
	// FORKID_ELDERBERRY is the fork id 8
	FORKID_ELDERBERRY = 8
	// FORKID_9 is the fork id 9
	FORKID_9 = 9
)

// ForkID is a sturct to track the ForkID event.
type ForkID struct {
	BatchNumber uint64
	ForkID      uint64
	Version     string
}
