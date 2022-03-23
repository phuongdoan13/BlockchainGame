import React from 'react';
import { BrowserRouter as Router, Route, Redirect } from 'react-router-dom';
// // import image
// import logo from './images/logo.svg';
// import ethereum_logo from './images/ethereum_logo.gif'

// import component
import NavigationBar from './components/ NavigationBar/NavigationBar';
import Homepage from './components/Homepage/Homepage';
import Matching from './components/Matching/Matching';
import Board from './components/Snake/Board/Board.js';
import Ropasi from './components/Ropasi/Ropasi';
import Play from './components/Ropasi/Play.js';
function App() {
  return (
    <div className="App">
        <NavigationBar></NavigationBar>
        <Router>
          <main>
            <Route path="/" exact component={Homepage} />
            <Route path="/matching"  component={Matching} />
            <Route path="/snake" component={Board}></Route>
            <Route path="/ropasi" component={Ropasi}></Route>
            <Route path="/ropasi/game" component={Play}></Route>
          </main>
        </Router>
    </div>
  );
}

export default App;
