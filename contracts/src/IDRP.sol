// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";

/// @title IDRP
/// @notice Mock IDR-pegged testnet token. 1 IDRP = Rp 1, 18 decimals.
contract IDRP is ERC20, AccessControl {
    bytes32 public constant FAUCET_ROLE = keccak256("FAUCET_ROLE");

    uint256 public constant FAUCET_AMOUNT = 50_000 ether;

    constructor(address admin) ERC20("Payung IDR (testnet)", "IDRP") {
        _grantRole(DEFAULT_ADMIN_ROLE, admin);
    }

    /// @notice Mints FAUCET_AMOUNT to `to`. Only callable by the backend faucet signer.
    function faucet(address to) external onlyRole(FAUCET_ROLE) {
        _mint(to, FAUCET_AMOUNT);
    }

    /// @notice Mints an arbitrary amount to `to`. Only callable by admin, used to fund the pool.
    function mint(address to, uint256 amount) external onlyRole(DEFAULT_ADMIN_ROLE) {
        _mint(to, amount);
    }
}
