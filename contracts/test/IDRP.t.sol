// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import {Test} from "forge-std/Test.sol";
import {IAccessControl} from "@openzeppelin/contracts/access/IAccessControl.sol";
import {IDRP} from "../src/IDRP.sol";

contract IDRPTest is Test {
    IDRP idrp;
    address admin = makeAddr("admin");
    address faucetSigner = makeAddr("faucetSigner");

    function setUp() public {
        vm.startPrank(admin);
        idrp = new IDRP(admin);
        idrp.grantRole(idrp.FAUCET_ROLE(), faucetSigner);
        vm.stopPrank();
    }

    function test_Metadata() public view {
        assertEq(idrp.name(), "Payung IDR (testnet)");
        assertEq(idrp.symbol(), "IDRP");
        assertEq(idrp.decimals(), 18);
    }

    function test_Faucet_MintsFixedAmount() public {
        address driver = makeAddr("driver");
        vm.prank(faucetSigner);
        idrp.faucet(driver);
        assertEq(idrp.balanceOf(driver), 50_000 ether);
    }

    function test_Faucet_NonFaucetRoleReverts() public {
        address stranger = makeAddr("stranger");
        bytes32 faucetRole = idrp.FAUCET_ROLE();
        vm.prank(stranger);
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, stranger, faucetRole)
        );
        idrp.faucet(stranger);
    }

    function test_Mint_OnlyAdmin() public {
        vm.prank(admin);
        idrp.mint(admin, 1_000_000 ether);
        assertEq(idrp.balanceOf(admin), 1_000_000 ether);

        address stranger = makeAddr("stranger");
        bytes32 adminRole = idrp.DEFAULT_ADMIN_ROLE();
        vm.prank(stranger);
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, stranger, adminRole)
        );
        idrp.mint(stranger, 1);
    }
}
