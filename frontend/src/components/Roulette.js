import React, { Component, Suspense } from 'react';
// import LanguageSwitcher from './language-switcher';
import DiceGame from '../Roulette/containers/dice-game';
import '../styles/Roulette.css';
class Roulette extends Component {
    render(){
        return(
            <Suspense fallback="Loading translation...">
                <div data-testid="App" className="Roulette">
                    <DiceGame />
                </div>
            </Suspense>
        )
    }
}

export default Roulette;