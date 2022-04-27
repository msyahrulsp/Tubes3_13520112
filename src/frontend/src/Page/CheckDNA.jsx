import { React } from "react";
import "./CheckDNA.scss";

import { Template } from '../Templates/Template';
import { Button } from '../Component/Button';

export const CheckDNA = () => {
    return (
        <Template>
            <div className="checkdna-container">
                <div className="addna-title">
                    <h3>Check DNA</h3>
                </div>
                <div className="checkdna-form">
                    <div className="addna-form-nama">
                    </div>
                    <div className="addna-form-file">
                    </div>
                    <div className="addna-form-prediction">
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