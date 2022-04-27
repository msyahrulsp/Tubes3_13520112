import { React } from "react";
import "./Home.scss";

import { Template } from '../Templates/Template';

export const Home = () => {
    return (
        <Template>
            <div className="home-container">
                <div className="text-wrapper">
                    <p>Penerapan String Matching dan</p>
                    <p>Regular Expression dalam</p>
                    <p>DNA Pattern Matching</p>
                    <h1>Try it Now!</h1>
                    <a href="/add-dna">Add DNA →</a>
                </div>
            </div>
        </Template>
    )
}