import { React } from "react";
import "./FindDNA.scss";

import { Template } from '../Templates/Template';
import { Button } from '../Component/Button';
import { Form } from "../Component/Form";

export const FindDNA = () => {
    return (
        <Template>
            <div className="finddna-container">
                <div className="finddna-title">
                    <h3>Find DNA</h3>
                </div>
                <div className="finddna-form">
                    <Form
                        type="text"
                        placeholder="Enter Value"
                    />
                </div>
                <div className="finddna-button">
                    <Button>
                        <p>Submit</p>
                    </Button>
                </div>
            </div>
        </Template>
    )
}