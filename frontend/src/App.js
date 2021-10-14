import React from 'react';

// import image
import logo from './images/logo.svg';
import ethereum_logo from './images/ethereum_logo.gif'

// import component
import NavigationBar from './components/NavigationBar';
import Homepage from './components/Homepage';
import Roulette from './components/Roulette';

function App() {
  return (
    <div className="App">
        <NavigationBar></NavigationBar>
        {/* <Homepage></Homepage> */}
        <Roulette></Roulette>
    </div>
  );
}

export default App;
