// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { ERC721 } from "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import { ERC721Bridge_Initializer } from "./CommonTest.t.sol";
import { LibRLP } from "./RLP.t.sol";
import { SliceMintableERC721 } from "../universal/SliceMintableERC721.sol";
import { SliceMintableERC721Factory } from "../universal/SliceMintableERC721Factory.sol";

contract SliceMintableERC721Factory_Test is ERC721Bridge_Initializer {
    SliceMintableERC721Factory internal factory;

    event SliceMintableERC721Created(
        address indexed localToken,
        address indexed remoteToken,
        address deployer
    );

    function setUp() public override {
        super.setUp();

        // Set up the token pair.
        factory = new SliceMintableERC721Factory(address(L2Bridge), 1);

        // Label the addresses for nice traces.
        vm.label(address(factory), "SliceMintableERC721Factory");
    }

    function test_constructor_succeeds() external {
        assertEq(factory.BRIDGE(), address(L2Bridge));
        assertEq(factory.REMOTE_CHAIN_ID(), 1);
    }

    function test_createSliceMintableERC721_succeeds() external {
        // Predict the address based on the factory address and nonce.
        address predicted = LibRLP.computeAddress(address(factory), 1);

        // Expect a token creation event.
        vm.expectEmit(true, true, true, true);
        emit SliceMintableERC721Created(predicted, address(1234), alice);

        // Create the token.
        vm.prank(alice);
        SliceMintableERC721 created = SliceMintableERC721(
            factory.createSliceMintableERC721(address(1234), "L2Token", "L2T")
        );

        // Token address should be correct.
        assertEq(address(created), predicted);

        // Should be marked as created by the factory.
        assertEq(factory.isSliceMintableERC721(address(created)), true);

        // Token should've been constructed correctly.
        assertEq(created.name(), "L2Token");
        assertEq(created.symbol(), "L2T");
        assertEq(created.REMOTE_TOKEN(), address(1234));
        assertEq(created.BRIDGE(), address(L2Bridge));
        assertEq(created.REMOTE_CHAIN_ID(), 1);
    }

    function test_createSliceMintableERC721_zeroRemoteToken_reverts() external {
        // Try to create a token with a zero remote token address.
        vm.expectRevert("SliceMintableERC721Factory: L1 token address cannot be address(0)");
        factory.createSliceMintableERC721(address(0), "L2Token", "L2T");
    }
}
