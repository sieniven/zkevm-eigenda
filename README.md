# Polygon zkEVM Node with EigenDA Layer

- Proof of concept for using EigenDA layer for off-chain data availability with Polygon CDK zkEVM node.
- Contains a minimal zkevm node for mock batch sequencing for the PoC.
- Full integration of off-chain pipeline for submitting L2 batches data to EigenDA layer for off-chain data availability solution.
- Pipeline to retrieval client to retrieve EigenDA blobs and decode into zkevm L2 batches data.
- Full on-chain solution to integrate EigenDA blob verification pipeline with existing Polygon zkevm-validium interfaces for sequence verification.
- The PoC aims to be verified to be working on Ethereum Holesky network.

# Current testnet deployment

The current testnet deployment is on Ethereum Holesky, with the deployed contract addresses below.

| Name | Address |
| ----------- | ----------- |
| [`EigenDARollupUtils`](https://github.com/Layr-Labs/eigenda/blob/dbbe9d1df5741e7cc32d833df7b07a3ebc733ea7/contracts/src/libraries/EigenDARollupUtils.sol) | [`0xe65b98311240ea0d545fc5a7Fe10eE5B53e0E91f`](https://holesky.etherscan.io/address/0xe65b98311240ea0d545fc5a7fe10ee5b53e0e91f) |
| [`EigenDAVerifier`](https://github.com/sieniven/zkevm-eigenda/blob/9a094f2648b10e942126069f93aef4f33b8b0fa5/contracts/eigenda/src/EigenDAVerifier.sol) | [`0xe65b98311240ea0d545fc5a7Fe10eE5B53e0E91f`](https://holesky.etherscan.io/address/0xe65b98311240ea0d545fc5a7fe10ee5b53e0e91f) |
