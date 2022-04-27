import { React } from "react";
import "./FindDNA.scss";

import { Template } from '../Templates/Template';
import { Button } from '../Component/Button';

export const FindDNA = () => {
    return (
        <Template>
            <div className="finddna-container">
                <div className="finddna-title">
                    <h3>Find DNA</h3>
                </div>
                <div className="finddna-form">
                    <div className="finddna-form-nama">
                    </div>
                    <div className="finddna-form-file">
                    </div>
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