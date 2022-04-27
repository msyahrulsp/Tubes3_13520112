import { React } from "react";
import "./Home.scss";

import { Template } from '../Templates/Template';
import { Button } from '../Component/Button';

export const Home = () => {
    return (
        <Template>
            <div className="home-container">
                <div className="text-wrapper">
                    <div className="title-wrapper">
                        <p>Penerapan String Matching dan</p>
                        <p>Regular Expression dalam</p>
                        <p>DNA Pattern Matching</p>
                    </div>
                    <h1>Try it Now!</h1>
                    <a href="/add-dna">
                        <Button>
                            Add DNA
                        </Button>
                    </a>
                </div>
            </div>
        </Template>
    )
}