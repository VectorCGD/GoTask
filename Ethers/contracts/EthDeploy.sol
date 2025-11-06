// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;


contract EthDeploy {
    string data;
    function SetData(string calldata _data) public {
        data = _data;
    }
    function GetData() public view returns(string memory){
        return data;
    }
}