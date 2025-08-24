// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

contract CountNumber {
    uint256 number;

    function add() public {
        number++;
    }

    function get() public view returns (uint256) {
        return number;
    }
}
