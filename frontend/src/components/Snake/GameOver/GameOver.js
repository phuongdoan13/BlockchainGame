import React from 'react'
const axios = require('axios').default;
class GameOver extends React.Component {
    grantERC20(){
        // Transfer ERC20 to the player
        let url = 'http://localhost:8080/winERC20'
        let config = {
            headers: {
                "Content-Type": "application/json",
                'Access-Control-Allow-Origin': '*',
            }
        }
        let data = {
            "PublicAddress" : this.props.account,
            "Amount": this.props.score
        }
        axios.post(url, data, config)
        .then(
            (response) => {console.log(response)},
            (error) => {console.log(error);}
        );
    }

    componentDidMount () {
        this.grantERC20()
    }

    render() {
        return (
            <div
                id='GameBoard'
                style={{
                width: this.props.width,
                height: this.props.height,
                borderWidth: this.props.width / 50,
                marginTop: 100,}}>
                <div id='GameOver' style={{ fontSize: this.props.width / 15 }}>
                    <div id='GameOverText'>GAME OVER</div>
                    <div>Your score: {this.props.score}</div>
                    <div>
                          {this.props.newHighScore ? 'New local ' : 'Local '}high score:{' '}
                          {this.props.highScore}
                    </div>
                    <div id='PressSpaceText'>Press Space to restart</div>
                </div>
            </div>
        )
    }
}

export default GameOver