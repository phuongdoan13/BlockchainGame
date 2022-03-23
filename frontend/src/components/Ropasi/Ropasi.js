import React, { useState } from "react";
import Header from "./Header";
import Play from "./Play";
import Game from "./Game";
import Footer from "./Footer";
import { Switch, Route } from "react-router-dom";

function Ropasi(props) {
    const [myChoice, setMyChoice] = useState("");
    const [score, setScore] = useState(0);
    console.log(props.location.pathname + "/game1")
    return (
        <>
            <div className="container">
                <Header score={score} />
                <Switch>
                    <Route path={props.location.pathname + '/game1'}>
                        <Play setMyChoice={setMyChoice} />
                    </Route>
                    <Route path="./game">
                        <Game myChoice={myChoice} score={score} setScore={setScore} />
                    </Route>
                </Switch>
            </div>
            {/* <Footer /> */}
        </>
    );
    }

export default Ropasi;
