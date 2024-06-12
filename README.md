# EigenDA Layer Integration with zkEVM Validium Node

## PoC Design

- Proof of concept for using EigenDA layer as the data availability provider for off-chain data availability with Polygon CDK zkEVM validium node.
- Full on-chain solution to integrate EigenDA blob verification pipeline with existing Polygon zkevm-validium interfaces for sequence verification.
  - Implements EigenDAVerifier contract that implements the [`IDataAvailabilityProtocol`](https://github.com/0xPolygonHermez/zkevm-contracts/blob/1ad7089d04910c319a257ff4f3674ffd6fc6e64e/contracts/v2/interfaces/IDataAvailabilityProtocol.sol) interface
  - To maintain backwards compatibility, the EigenDAVerifier contract is designed for an adaptor pattern verifies on-chain the data availability proofs with the EigenDAServiceManager.
- Full off-chain pipeline integration for submitting L2 batches data to EigenDA layer as our off-chain data availability solution.
  - Contains a minimal zkevm node sequencer and sequence sender services for mock batch sequencing to validate the functionality of the PoC.
  - Pipeline to retrieval client to retrieve EigenDA blobs and decode into zkevm L2 batches data.

## Testnet deployment

The PoC is verified on Ethereum Holesky with the deployed contract addresses below.

| Name | Address |
| ----------- | ----------- |
| [`EigenDARollupUtils`](https://github.com/Layr-Labs/eigenda/blob/dbbe9d1df5741e7cc32d833df7b07a3ebc733ea7/contracts/src/libraries/EigenDARollupUtils.sol) | [`0xe65b98311240ea0d545fc5a7Fe10eE5B53e0E91f`](https://holesky.etherscan.io/address/0xe65b98311240ea0d545fc5a7fe10ee5b53e0e91f) |
| [`EigenDAVerifier`](https://github.com/sieniven/zkevm-eigenda/blob/9a094f2648b10e942126069f93aef4f33b8b0fa5/contracts/eigenda/src/EigenDAVerifier.sol) | [`0x4AD03109f48B8a15B7496a9A764C6D00aef0aE36`](https://holesky.etherscan.io/address/0x4ad03109f48b8a15b7496a9a764c6d00aef0ae36) |
| [`EigenDAServiceManager`](https://github.com/Layr-Labs/eigenda/blob/a33b41561cc3fb4cd6d50a8738e4c5dca43ec0a5/contracts/src/core/EigenDAServiceManager.sol) | [`0xD4A7E1Bd8015057293f0D0A557088c286942e84b`](https://holesky.etherscan.io/address/0xa7227485e6C693AC4566fe168C5E3647c5c267f3) |
