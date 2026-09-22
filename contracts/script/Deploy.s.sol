// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import {Script, console} from "forge-std/Script.sol";
import {IDRP} from "../src/IDRP.sol";
import {PayungPool} from "../src/PayungPool.sol";

/// @notice Deploys IDRP + PayungPool, seeds the 5 MVP zones, grants operational roles,
///         mints deployer working capital and funds the pool.
/// @dev Reads PRIVATE_KEY (deployer/admin), ORACLE_ADDRESS and FAUCET_ADDRESS from env.
contract DeployScript is Script {
    uint256 constant INITIAL_PREMIUM_PER_WEEK = 5_000 ether;
    uint16 constant THRESHOLD_MM = 20;
    uint256 constant PAYOUT_PER_DAY = 25_000 ether;
    uint8 constant MAX_DAYS_PER_WEEK = 3;

    uint256 constant DEPLOYER_MINT = 5_000_000 ether;
    uint256 constant INITIAL_POOL_FUNDING = 2_000_000 ether;

    function run() external {
        uint256 deployerKey = vm.envUint("PRIVATE_KEY");
        address oracle = vm.envAddress("ORACLE_ADDRESS");
        address faucet = vm.envAddress("FAUCET_ADDRESS");
        address deployer = vm.addr(deployerKey);

        vm.startBroadcast(deployerKey);

        IDRP idrp = new IDRP(deployer);
        PayungPool pool = new PayungPool(address(idrp), deployer);

        pool.grantRole(pool.ORACLE_ROLE(), oracle);
        idrp.grantRole(idrp.FAUCET_ROLE(), faucet);

        pool.createZone(1, "Yogyakarta", INITIAL_PREMIUM_PER_WEEK, THRESHOLD_MM, PAYOUT_PER_DAY, MAX_DAYS_PER_WEEK);
        pool.createZone(2, "Sleman", INITIAL_PREMIUM_PER_WEEK, THRESHOLD_MM, PAYOUT_PER_DAY, MAX_DAYS_PER_WEEK);
        pool.createZone(3, "Bantul", INITIAL_PREMIUM_PER_WEEK, THRESHOLD_MM, PAYOUT_PER_DAY, MAX_DAYS_PER_WEEK);
        pool.createZone(4, "Surakarta", INITIAL_PREMIUM_PER_WEEK, THRESHOLD_MM, PAYOUT_PER_DAY, MAX_DAYS_PER_WEEK);
        pool.createZone(5, "Semarang", INITIAL_PREMIUM_PER_WEEK, THRESHOLD_MM, PAYOUT_PER_DAY, MAX_DAYS_PER_WEEK);

        idrp.mint(deployer, DEPLOYER_MINT);
        idrp.approve(address(pool), INITIAL_POOL_FUNDING);
        pool.fundPool(INITIAL_POOL_FUNDING);

        vm.stopBroadcast();

        console.log("IDRP deployed at:", address(idrp));
        console.log("PayungPool deployed at:", address(pool));
        console.log("Oracle role granted to:", oracle);
        console.log("Faucet role granted to:", faucet);
    }
}
