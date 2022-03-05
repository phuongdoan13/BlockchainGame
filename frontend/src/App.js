import React from 'react';
import { BrowserRouter as Router, Route, Redirect } from 'react-router-dom';
// // import image
// import logo from './images/logo.svg';
// import ethereum_logo from './images/ethereum_logo.gif'

// import component
import NavigationBar from './components/NavigationBar';
import Homepage from './components/Homepage';
import Roulette from './components/Roulette';
import Matching from './components/Matching';
import Board from './components/Snake/Board/Board.js'
function App() {
  return (
    <div className="App">
        <NavigationBar></NavigationBar>
        <Router>
          <main>
            <Route path="/" exact component={Homepage} />
            <Route path="/matching"  component={Matching} />
            <Route path="/roulette"  component={Roulette} />
            <Route path="/snake" component={Board}></Route>
          </main>
        </Router>
    </div>
  );
}

export default App;
