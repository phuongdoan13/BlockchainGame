import React, { Component } from 'react';

import Web3 from 'web3';
import "./AssetsBar.css";
import {CARD_ARRAY} from "../const/MemoryToken_const"
const axios = require('axios').default;
const PUBLIC_URL = process.env.PUBLIC_URL;

class AssetsBar extends Component {

    async componentWillMount() {
        await this.loadWeb3()
        await this.loadBlockchainData()
        this.setState({ cardArray: CARD_ARRAY.sort(() => 0.5 - Math.random()) })

    }

    async loadWeb3() {
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
    async loadSmartContractABI(){
        const res = await axios('http://localhost:8080/matchingABI');
        return await res.data;
    }

    async loadBlockchainData() {
        const web3 = window.web3
        const accounts = await web3.eth.getAccounts()
        this.setState({ account: accounts[0] })

        // Load smart contract
        const MemoryToken = await this.loadSmartContractABI()
        const networkId = await web3.eth.net.getId()
        const networkData = MemoryToken.networks[networkId]
        if(networkData) {
            const abi = MemoryToken.abi
            const address = networkData.address
            const token = new web3.eth.Contract(abi, address)
            this.setState({ token })
            const totalSupply = await token.methods.totalSupply().call()
            this.setState({ totalSupply })
            // Load Tokens
            let balanceOf = await token.methods.balanceOf(accounts[0]).call()
            for (let i = 0; i < balanceOf; i++) {
                let id = await token.methods.tokenOfOwnerByIndex(accounts[0], i).call()
                let tokenURI = await token.methods.tokenURI(id).call()

                this.setState({
                    tokenURIs: [...this.state.tokenURIs, tokenURI]
                })
            }
        } else {
            alert('Smart contract not deployed to detected network.')
        }
    }

    constructor(props) {
        super(props)
        this.state = {
            account: '0x0',
            token: null,
            totalSupply: 0,
            tokenURIs: [],
            cardArray: [],

        }
    }

    render() {
        return (
            <div>
                <h5>Tokens Collected:<span id="result">&nbsp;{this.state.tokenURIs.length}</span></h5>
                <div className="grid mb-4" >
                    { this.state.tokenURIs.map((tokenURI, key) => {
                        return(
                            <img
                                key={key}
                                src={tokenURI}
                            />
                        )
                    })}
                </div>
            </div>

        );
    }
}

export default AssetsBar;
