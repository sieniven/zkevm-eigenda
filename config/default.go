package config

// DefaultValues is the default configuration
const DefaultValues = `
[Log]
Environment = "development" # "production" or "development"
Level = "info"
Outputs = ["stderr"]

[Etherman]
URL = "http://localhost:8545"

[EthTxManager]
FrequencyToMonitorTxs = "1s"
WaitTxToBeMined = "2m"
ForcedGas = 0
GasPriceMarginFactor = 1
MaxGasPriceLimit = 0

[SequenceSender]
WaitPeriodSendSequence = "5s"
MaxTxSizeForL1 = 131072
L2Coinbase = "0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266"
GasOffset = 80000
MaxBatchesForL1 = 10

[DataAvailability]
Hostname = "disperser-holesky.eigenda.xyz"
Port = 443
Timeout = "30s"
UseSecureGrpcFlag = true
RetrieveBlobStatusPeriod = "5s"
BlobStatusConfirmedTimeout = "15m"
`
