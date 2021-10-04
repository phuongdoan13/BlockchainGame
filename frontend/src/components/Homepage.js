import React, { Component } from 'react';
// import image
import ethereum_logo from '../images/ethereum_logo.gif'

// import css
import '../styles/App.css';
class Homepage extends Component {
    render(){
        return(
            <div className="App">
                <header className="App-header">
                    <div>
                        <img src={ethereum_logo} className="App-logo" alt="logo"/>
                    </div>

                    <div className={"buttons"}>
                        <button className="button">Matching</button>
                        <button className="button">Roulette</button>
                    </div>

                </header>
            </div>
        );
    }
}

export default Homepage;