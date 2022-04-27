import { React } from "react";
import "./AddDNA.scss";

import { Template } from '../Templates/Template';
import { Button } from '../Component/Button';
import { Form } from "../Component/Form";

export const AddDNA = () => {
    return (
        <Template>
            <div className="adddna-container">
                <div className="addna-title">
                    <h3>Add DNA</h3>
                </div>
                <div className="addna-form">
                    <Form
                        title="Disease"
                        type="text"
                        placeholder="Enter Disease Name"
                    />
                    <Form
                        title="DNA Sequence"
                        type="file"
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