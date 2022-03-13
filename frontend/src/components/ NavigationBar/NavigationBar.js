import React, { Component } from 'react';
// import Web3 from 'web3';

// import css
import '../../App.css';

class NavigationBar extends Component {
    constructor(props) {
        super(props);
        this.state = {
          account: 'Welcome back!',
        }
    }
    // async componentWillMount() {
    //     await this.loadWeb3();
    //     await this.loadAccountData();
    // }

    // async loadWeb3() {
    //     /*
    //     * Method to connect to the blockchain
    //     */
    //     if (window.ethereum) {
    //         window.web3 = new Web3(window.ethereum)
    //         await window.ethereum.enable()
    //     }
    //     else if (window.web3) {
    //         window.web3 = new Web3(window.web3.currentProvider)
    //     }
    //     else {
    //         window.alert('Non-Ethereum browser detected. You should consider trying MetaMask!')
    //     }
    // }

    // async loadAccountData() {
    //     /*
    //     * Method for getting the current user's data;
    //     */
    //     const web3 = window.web3;
    //     const accounts = await web3.eth.getAccounts();
    //     this.setState({ account: accounts[0]});
    // }

    render() {
        return (
            <>
                <nav className="navbar navbar-dark fixed-top bg-dark flex-md-nowrap p-0 shadow">
                    <a className="navbar-brand col-sm-3 col-md-2 mr-0" target="_self" rel="noopener noreferrer" href="/">
                      &nbsp; Ethereum world
                    </a>
                    <a className="navbar-brand col-sm-3 col-md-2 mr-0" target="_self" rel="noopener noreferrer" href="/matching">
                    <img src={process.env.PUBLIC_URL + '/NavigationBar/brain.png'} width="30" height="30" className="d-inline-block align-top" alt="" />
                    &nbsp; Matching
                    </a>
                    <a className="navbar-brand col-sm-3 col-md-2 mr-0" target="_self" rel="noopener noreferrer" href="/snake">
                    <img src={process.env.PUBLIC_URL + '/NavigationBar/Snake.png'} width="30" height="30" className="d-inline-block align-top" alt="" />
                    &nbsp; Snake
                    </a>
                    {/* <a className="navbar-brand col-sm-3 col-md-2 mr-0" target="_blank" rel="noopener noreferrer">
                        <img src={annonymous} width="30" height="30" className="d-inline-block align-top" alt="" />
                        &nbsp; {this.state.account}
                    </a> */}
                </nav>
            </>
        );
    }
}

export default NavigationBar;
