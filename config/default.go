package config

// DefaultValues is the default configuration
const DefaultValues = `
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
`
