import React, { useState } from "react";
import { withRouter } from "react-router";

import { loginRequest } from "../../../services/api-service";
import { setItemToSS } from "../../../services/storage-service";

import "./Login.css";

function Login(props) {
  const [loginErrorMessage, setErrorMessage] = useState(null);
  const [username, setUsername] = useState(null);
  const [password, setPassword] = useState(null);

  const handleUsernameChange = (event) => {
    setUsername(event.target.value);
  };

  const handlePasswordChange = (event) => {
    setPassword(event.target.value);
  };

  const loginUser = async () => {
    const resPayload = await loginRequest(username, password);
    if (resPayload.code === 200) {
      setItemToSS("userDetail", resPayload.response);
      console.log("go home");
      props.history.push(`/home`);
    } else {
      setErrorMessage(resPayload.message);
    }
  };

  return (
    <div className="app__login-container">
      <div className="app__form-row">
        <label>Username: </label>
        <input type="email" className="email" onChange={handleUsernameChange} />
      </div>

      <div className="app__form-row">
        <label>Password: </label>
        <input
          type="password"
          className="password"
          onChange={handlePasswordChange}
        />
      </div>

      <div className="app__form-row">
        <span className="error-message">
          {loginErrorMessage ? loginErrorMessage : ""}
        </span>
      </div>
      <div className="app__form-row" onClick={loginUser}>
        <button>Login</button>
      </div>
    </div>
  );
}

export default withRouter(Login);
