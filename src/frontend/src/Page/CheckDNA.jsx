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
  const [message, setMessage] = useState("");
  const [data, setData] = useState(null);

  const onChangeDisease = (event) => {
    setNamaPenyakit(event.target.value);
  };

  const onChangeNamaPengguna = (event) => {
    setNamaPengguna(event.target.value);
  };

  const onInputDNAHandler = (event) => {
    const reader = new FileReader();
    reader.onloadend = (e) => {
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
        if (res.data.message === "OK") {
          setData(res.data.Data);
          if (res.data.Data === null) {
            setMessage("No Data Found");
          } else {
            setMessage("");
          }
        } else {
          setData(null);
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
        <div className="checkna-button" onClick={onSubmitHandler}>
          <Button>
            <p>Submit</p>
          </Button>
        </div>
        <div className="checkdna-result">
          {data !== null ? data.map((item) => {
            let bool = item.hasil ? "True" : "False";
            return (
              <div className="checkdna-result-item">
                <h3>Result</h3>
                <p>{item.tanggal} - {item.namaPengguna} - {item.namaPenyakit} - {item.persentase}% - {bool}</p>
              </div>
            );
          }) : null}
        </div>
        <p className="message">{message}</p>
      </div>
    </Template>
  );
};
