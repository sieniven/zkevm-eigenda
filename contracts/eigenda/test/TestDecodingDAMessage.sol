// SPDX-License-Identifier: MIT

pragma solidity 0.8.20;

import "forge-std/Test.sol";
import "../src/EigenDAVerifier.sol";
import {IEigenDAServiceManager} from "eigenda/interfaces/IEigenDAServiceManager.sol";
import {BN254} from "eigenlayer-middleware/libraries/BN254.sol";

contract TestDecodeDAMessage is Test {
    address constant adminAddress = 0x0000000000000000000000000000000000000001;
    address constant serviceManager = 0x0000000000000000000000000000000000000002;
    // bytes testMessage = hex"00000000000000000000000000000000000000000000000000000000000030390000000000000000000000000000000000000000000000000000000000010932000000640000000101324b00000400000000010000000200000000000000000000000000000000000000000000000000000000000000000000000301020300000003323c460000303900000000000000000000000000000000000000000000000000000000000000000000d431000000030102030000000304050600000003010203";
    bytes testMessage = hex"f8b2f84df842a00000000000000000000000000000000000000000000000000000000000003039a0000000000000000000000000000000000000000000000000000000000001093264c7c601324b820400f85d0102f851eca000000000000000000000000000000000000000000000000000000000000000008301020383323c46823039a0000000000000000000000000000000000000000000000000000000000000000082d431830102038304050683010203";
    EigenDAVerifier verifierContract;

    function setUp() public {
        verifierContract = new EigenDAVerifier(adminAddress, IEigenDAServiceManager(serviceManager));        
    }

    function testDecodeTestMessage() public view {
        // Get correct blob data
        IEigenDAServiceManager.QuorumBlobParam[] memory params = new IEigenDAServiceManager.QuorumBlobParam[](1);
        params[0] = IEigenDAServiceManager.QuorumBlobParam(uint8(1), uint8(50), uint8(75), uint32(1024));
        EigenDAVerifier.BlobData memory data = EigenDAVerifier.BlobData(
            IEigenDAServiceManager.BlobHeader(
                BN254.G1Point(uint256(12345), uint256(67890)),
                uint32(100),
                params
            ),
            EigenDARollupUtils.BlobVerificationProof(
                uint32(1),
                uint32(2),
                IEigenDAServiceManager.BatchMetadata(
                    IEigenDAServiceManager.BatchHeader(
                        bytes32(0),
                        hex"010203",
                        hex"323c46",
                        uint32(12345)
                    ),
                    bytes32(0),
                    uint32(54321)
                ),
                hex"010203",
                hex"040506"
            ),
            hex"010203"
        );
        EigenDAVerifier.BlobData memory decodedData = verifierContract.decodeBlobData(testMessage);
        assertEq(data.blobHeader.commitment.X, decodedData.blobHeader.commitment.X);
    }
}
