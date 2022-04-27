import { React } from "react";
import "./AddDNA.scss";

import { Template } from '../Templates/Template';
import { Button } from '../Component/Button';

export const AddDNA = () => {
    return (
        <Template>
            <div className="adddna-container">
                <div className="addna-title">
                    <h3>Add DNA</h3>
                </div>
                <div className="addna-form">
                    <div className="addna-form-nama">
                    </div>
                    <div className="addna-form-file">
                    </div>
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