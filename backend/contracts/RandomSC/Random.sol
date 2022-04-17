pragma solidity ^0.5.0;

contract Random{
    string[] choices;

    constructor() public{
        choices = ["rock", "paper", "scissors"];
    }
    function randMod(uint256 random_number) public view returns(string memory){
        return choices[uint(keccak256(abi.encodePacked(random_number, block.timestamp, block.difficulty, msg.sender))) % 3];
    }
}