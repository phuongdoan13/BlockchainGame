import React, { useState, Component } from 'react';
import './Ropasi.css';
import Choices from '../choice/choice';
import Scoreboard from '../scoreboard/scoreboard';
import { Result } from '../result/result';
import Web3 from 'web3';
import RandomOrg from 'random-org';
const axios = require('axios').default;

const choiceOptions = ['rock', 'paper', 'scissors'];
const random_abi_json = require('../contracts_abi/random_abi.json');
const erc20_abi_json = require('../contracts_abi/erc20_abi.json');
class Ropasi extends Component{
  constructor(props){
    super(props);
    this.state = {
      wins: 0,
      losses: 0,
      status: 'win',
      player: 'Rock',
      computer: 'Paper'
    };
  };

  async componentWillMount() {
    await this.loadWeb3()
  }

  async loadWeb3() {
    // Load web3
    if (window.ethereum) {
      window.web3 = new Web3(window.ethereum)
      await window.ethereum.enable()
    }
    else if (window.web3) {
      window.web3 = new Web3(window.web3.currentProvider)
    }
    else {
      window.alert('Non-Ethereum browser detected. You should consider trying MetaMask!')
    }
  }
  async getContract(contract_json){
    // Get the contract instance
    const web3 = window.web3;
    const abi = contract_json.abi;
    const address = contract_json.address;

    try{
      const contract = new web3.eth.Contract(abi, address);
      console.log(contract)
      return contract;
    }catch(err){
      alert(err.stack);
      window.location.reload(); 
    }
  }

  async getRandomSeed(){
    // Get a random number from random.org website
    var random = new RandomOrg({ apiKey: 'f7a7b104-48cd-4ee2-bebd-749a5b1d58d9' });
    return await random.generateIntegers({min: 0, max: 1000, n: 1});
  }

  async getComputerChoice() {
    // Get the computer choice 'rock', 'paper', 'scissors'
    
    // Get random seed
    var randomSeed;
    try{
      var response = await this.getRandomSeed();
      randomSeed = response.random.data[0];
    }catch(err){
      randomSeed = Math.floor(Math.random() * 1000);
      console.log(err.stack);
    }
    
    // Get contract instance and get Computer's choice
    try{
      const contract = await this.getContract(random_abi_json);
      const computerChoice = await contract.methods.randMod(randomSeed).call();
      return computerChoice;
    }catch(err){
      console.log(err.stack);
      // window.location.reload();
    }
  }

  playerWinMoney = async () => {
    // Transfer money from the dealer to player
    const AMOUNT = 1;
     // Get contract instance and get Computer's choice
     try{
        const dealerPublicAddress = erc20_abi_json.dealerPublicAddress; // # Look here
        const accounts = await window.web3.eth.getAccounts();
        const currentPlayerPublicAddress = accounts[0]; // # Look here
        const contract = await this.getContract(erc20_abi_json);

        await contract.methods.transferToPlayer(AMOUNT)
            .send({from : currentPlayerPublicAddress})
            .catch(console.error);

        alert("You win 1 coin!")
     }catch(err){
        alert(err);
     }

  }

  async playerLoseMoney(){
    // Transfer money from player to the dealer
    const AMOUNT = 1;
    try{
      const dealerPublicAddress = erc20_abi_json.dealerPublicAddress; // # Look here
      const accounts = await window.web3.eth.getAccounts();
      const currentPlayerPublicAddress = accounts[0]; // # Look here
      const contract = await this.getContract(erc20_abi_json);
      await contract.methods.transfer(dealerPublicAddress, AMOUNT)
          .send({from : currentPlayerPublicAddress})
          .catch(console.error);

      alert("You lose 1 coin!")
    }catch(err){
      alert(err.stack);
      window.location.reload(); 
    }
  }
  async game (userChoice){
    // Determine the game result;
    await this.getComputerChoice().then(computerChoice => {
      const rock = 'rock';
      const scissors = 'scissors';
      const paper = 'paper';
      switch (userChoice.toLowerCase() + computerChoice.toLowerCase()) {
        case rock + scissors:
        case paper + rock:
        case scissors + paper:
          this.setState({ status: 'win', player: userChoice, computer: computerChoice });
          break;
        case scissors + rock:
        case rock + paper:
        case paper + scissors:
          this.setState({ status: 'lose', player: userChoice, computer: computerChoice });
          break;
        default:
          this.setState({ status: 'draw', player: userChoice, computer: computerChoice });
          break;
      }
      if(this.state.status == 'win'){
        this.playerWinMoney();
      }else if(this.state.status == 'lose'){
        this.playerLoseMoney();
      }
    })
  };

  onChoiceClick = async (choice) => {
    // Handle click
    await this.game(choice);
    switch (this.state.status) {
      case 'win':
        this.state.wins += 1;
        break;
      case 'lose':
        this.state.losses += 1;
        break;
      default:
        break;
    }
    return this.state;
  };

  render(){
    return (
      <div className="Ropasi">
        <header></header>
        {/* <Scoreboard score={this.state} /> */}
        <Result gameStatus={this.state} />
        <Choices choices={choiceOptions} onClick={this.onChoiceClick} />
        <p>Make your move!</p>
      </div>
    );
  }
}

export default Ropasi;
