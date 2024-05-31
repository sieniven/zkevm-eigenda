# Polygon zkEVM Node with EigenDA Layer

- Proof of concept for using EigenDA layer for off-chain data availability with Polygon CDK zkEVM node.
- Contains a minimal zkevm node for mock batch sequencing for the PoC.
- Full integration of off-chain pipeline for submitting L2 batches data to EigenDA layer for off-chain data availability solution.
- Pipeline to retrieval client to retrieve EigenDA blobs and decode into zkevm L2 batches data.
- Full on-chain solution to integrate EigenDA blob verification pipeline with existing Polygon zkevm-validium interfaces for sequence verification.
- The PoC aims to be verified to be working on Ethereum Holesky network.
