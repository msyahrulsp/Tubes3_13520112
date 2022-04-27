import { React } from "react";
import "./CheckDNA.scss";

import { Template } from '../Templates/Template';
import { Button } from '../Component/Button';
import { Form } from "../Component/Form";


export const CheckDNA = () => {
    return (
        <Template>
            <div className="checkdna-container">
                <div className="addna-title">
                    <h3>Check DNA</h3>
                </div>
                <div className="checkdna-form">
                    <Form
                        title="Name"
                        type="text"
                        placeholder="Enter Name"
                    />
                    <Form
                        title="DNA Sequence"
                        type="file"
                    />
                    <Form 
                        title="Prediction"
                        type="text"
                        placeholder="Enter Disease Prediction"
                    />
                </div>
                <div className="addna-button">
                    <Button>
                        <p>Submit</p>
                    </Button>
                </div>
            </div>
        </Template>
    )
}