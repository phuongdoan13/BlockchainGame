pragma solidity ^0.5.0;

interface IERC20 {

    function totalSupply() external view returns (uint256);
    function balanceOf(address account) external view returns (uint256);
    function allowance(address owner, address spender) external view returns (uint256);

    function transfer(address recipient, uint256 amount) external returns (bool);
    function approve(address spender, uint256 amount) external returns (bool);
    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);


    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);
}

library SafeMath {
    function sub(uint256 a, uint256 b) internal pure returns (uint256) {
        assert(b <= a);
        return a - b;
    }

    function add(uint256 a, uint256 b) internal pure returns (uint256) {
        uint256 c = a + b;
        assert(c >= a);
        return c;
    }
}

contract ERC20Token is IERC20 {
    // Contract ERC20 token that inherit the IERC20 interface (standard)
    string public constant name = "ERC20Token";
    string public constant symbol = "E20";
    uint256 public constant decimals = 18;
    address payable deployer;

    // Map address with balance for viewing
    event Approval(address indexed tokenOwner, address indexed spender, uint tokens);
    event Transfer(address indexed from, address indexed to, uint tokens);


    mapping(address => uint256) balances;

    mapping(address => mapping (address => uint256)) allowed;

    uint256 totalSupply_;

    // using SafeMath for uint256;

    using SafeMath for uint256;

    constructor(uint256 total) public {
        // Constructor to specify the number of tokens available
        totalSupply_ = total;
        balances[msg.sender] = totalSupply_; // the person who created the smart contract will initally own all
        deployer = msg.sender;
    }

    function totalSupply() public view returns (uint256) {
        // Get the total supply of token
        return totalSupply_;
    }

    function balanceOf(address tokenOwner) public view returns (uint256) {
        // Get the balance of an owner
        return balances[tokenOwner];
    }

    function transferToPlayer(uint256 numTokens) public returns (bool) {
        // Transfer from the deployer to the specified receiver
        require(numTokens <= balances[deployer]);
        balances[deployer] = balances[deployer].sub(numTokens);
        balances[msg.sender] = balances[msg.sender].add(numTokens);
        emit Transfer(deployer, msg.sender, numTokens);
        return true;
    }

    function transfer(address receiver, uint256 numTokens) public returns (bool) {
        // Transfer from the deployer to the specified receiver
        require(numTokens <= balances[msg.sender]);
        balances[msg.sender] = balances[msg.sender].sub(numTokens);
        balances[receiver] = balances[receiver].add(numTokens);
        emit Transfer(msg.sender, receiver, numTokens);
        return true;
    }

    function approve(address delegate, uint256 numTokens) public returns (bool){
        // Approve of the allowance
        allowed[msg.sender][delegate] = numTokens;
        emit Approval(msg.sender, delegate, numTokens);
        return true;
    }
    
    function allowance(address owner, address delegate) public view returns (uint) {
        return allowed[owner][delegate];
    }

    function transferFrom(address owner, address buyer, uint256 numTokens) public returns(bool){
        require(numTokens <= balances[owner]);
        require(numTokens <= allowed[owner][msg.sender]);

        balances[owner] = balances[owner].sub(numTokens);
        allowed[owner][msg.sender] = allowed[owner][msg.sender].sub(numTokens);
        balances[buyer] = balances[buyer].add(numTokens);
        emit Transfer(owner, buyer, numTokens);
        return true;
    }

}
