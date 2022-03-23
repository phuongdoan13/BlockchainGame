pragma solidity >=0.5.0;

import "./ERC721Full.sol";
contract MemoryToken is ERC721Full  {
    string public name;
    string public symbol;
    constructor(string memory _name, string memory _symbol) ERC721Full(_name, _symbol) public {
        // a constructor inherit ERC721Full
        // param:
        // * name: the fullname of the smart contract
        // * symbol: the symbol name of the smart contract
        name =  _name;
        symbol = _symbol;
    }

    function mint(address _to, string memory _tokenURI) public returns(bool) {
        uint _tokenId = totalSupply().add(1); // increment the id of the token by 1;
       _mint(_to, _tokenId);
       _setTokenURI(_tokenId, _tokenURI);
       return true;
    }
}
