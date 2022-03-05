import React, { Component } from 'react';
import { Link } from 'react-router-dom';


// import css
import '../styles/App.css';

class Homepage extends Component {
    render(){
        return(
            <div className="App">
                <header className="App-header">
                    <div>
                        <img src={process.env.PUBLIC_URL + '/ethereum_logo.gif'} className="App-logo" alt="logo"/>
                    </div>

                    <div className={"buttons"}>
                        <Link to="/matching">
                            <button className="button">Matching</button>
                        </Link>
                        <Link to="/snake">
                            <button className="button">Snake</button>
                        </Link>
                    </div>

                </header>
            </div>
        );
    }
}

export default Homepage;