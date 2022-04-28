import { React, useState } from "react";
import "./FindDNA.scss";
import axiosConfig from "../axiosConfig";
import qs from "qs";

import { Template } from "../Templates/Template";
import { Button } from "../Component/Button";
import { Form } from "../Component/Form";

export const FindDNA = () => {
  const [query, setQuery] = useState("");

  const onChangeQuery = (event) => {
    //   console.log(event.target.value)
    setQuery(event.target.value);
  };

  const onSubmitHandler = (event) => {
    const body = {
      query: query,
    };
    try {
      axiosConfig.post(`/hasil/find`, qs.stringify(body)).then((res) => {
        console.log(res);
      });
    } catch (err) {
      alert(err.toString());
    }
  };
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
            onChange={onChangeQuery}
          />
        </div>
        <div className="finddna-button">
          <Button onClick={onSubmitHandler}>
            <p>Submit</p>
          </Button>
        </div>
      </div>
    </Template>
  );
};
