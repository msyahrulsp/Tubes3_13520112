import { React, useState } from "react";
import "./FindDNA.scss";
import axiosConfig from "../axiosConfig";
import qs from "qs";

import { Template } from "../Templates/Template";
import { Button } from "../Component/Button";
import { Form } from "../Component/Form";

export const FindDNA = () => {
  const [query, setQuery] = useState("");
  const [data, setData] = useState(null);
  const [message, setMessage] = useState();

  const onChangeQuery = (event) => {
    setQuery(event.target.value);
  };

  const onSubmitHandler = (event) => {
    const body = {
      query: query,
    };
    try {
      axiosConfig.post(`/hasil/find`, qs.stringify(body)).then((res) => {
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
        <div className="finddna-result">
          {data !== null ? data.map((item) => {
            let bool = item.hasil === 1 ? "True" : "False";
            return (
              <div className="finddna-result-item">
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