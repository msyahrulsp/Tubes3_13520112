import { React, useState } from "react";
import "./CheckDNA.scss";
import axiosConfig from "../axiosConfig";
import qs from "qs";

import { Template } from "../Templates/Template";
import { Button } from "../Component/Button";
import { Form } from "../Component/Form";

export const CheckDNA = () => {
  const [namaPengguna, setNamaPengguna] = useState("");
  const [namaPenyakit, setNamaPenyakit] = useState("");
  const [DNA, setDNA] = useState("");

  const onChangeDisease = (event) => {
    //   console.log(event.target.value)
    setNamaPenyakit(event.target.value);
  };

  const onChangeNamaPengguna = (event) => {
    //   console.log(event.target.value)
    setNamaPengguna(event.target.value);
  };

  const onInputDNAHandler = (event) => {
    const reader = new FileReader();
    reader.onloadend = (e) => {
      //   const content = e.target.result;
      //   console.log('file content',  content)
      setDNA(e.target.result);
    };
    reader.readAsText(event.target.files[0]);
  };

  const onSubmitHandler = (event) => {
    const body = {
      namaPenyakit: namaPenyakit,
      namaPengguna: namaPengguna,
      DNA: DNA,
    };
    try {
      axiosConfig.post(`/hasil`, qs.stringify(body)).then((res) => {
        console.log(res);
      });
    } catch (err) {
      alert(err.toString());
    }
  };

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
            onChange={onChangeNamaPengguna}
          />
          <Form title="DNA Sequence" type="file" onChange={onInputDNAHandler} />
          <Form
            title="Prediction"
            type="text"
            placeholder="Enter Disease Prediction"
            onChange={onChangeDisease}
          />
        </div>
        <div className="addna-button" onClick={onSubmitHandler}>
          <Button>
            <p>Submit</p>
          </Button>
        </div>
      </div>
    </Template>
  );
};
