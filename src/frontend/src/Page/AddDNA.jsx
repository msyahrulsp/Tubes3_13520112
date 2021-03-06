import { React, useState } from "react";
import "./AddDNA.scss";
import axiosConfig from "../axiosConfig";
import qs from "qs";

import { Template } from "../Templates/Template";
import { Button } from "../Component/Button";
import { Form } from "../Component/Form";

export const AddDNA = () => {
  const [namaPenyakit, setNamaPenyakit] = useState("");
  const [sequenceDNA, setSequenceDNA] = useState("");
  const [message, setMessage] = useState("");

  const onChangeDisease = (event) => {
    setNamaPenyakit(event.target.value);
  };

  const onInputDNAHandler = (event) => {
    const reader = new FileReader();
    reader.onloadend = (e) => {
      setSequenceDNA(e.target.result);
    };
    reader.readAsText(event.target.files[0]);
  };

  const onSubmitHandler = (event) => {
    const body = {
      namaPenyakit: namaPenyakit,
      sequenceDNA: sequenceDNA,
    };
    try {
      axiosConfig.post(`/penyakit`, qs.stringify(body)).then((res) => {
        if (res.data.message === "OK") {
          setMessage("Success Adding New Disease Sequence");
        } else {
          setMessage("Please Check Your Input")
        }
        console.log(res);
      });
    } catch (err) {
      alert(err.toString());
    }
  };

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
            onChange={onChangeDisease}
          />
          <Form title="DNA Sequence" type="file" onChange={onInputDNAHandler} />
        </div>
        <div className="addna-button">
          <Button onClick={onSubmitHandler}>
            <p>Submit</p>
          </Button>
        </div>
        <p className="message">{message}</p>
      </div>
    </Template>
  );
};
