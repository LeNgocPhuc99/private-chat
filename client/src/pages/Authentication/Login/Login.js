import React, { useState } from "react";
import { withRouter } from "react-router";

import { loginRequest } from "../../../services/api-service";
import { setItemToLS } from "../../../services/storage-service";

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
    // console.log(username);
    // console.log(password);
    const resPayload = await loginRequest(username, password);
    if (resPayload.code === 200) {
      setItemToLS("userDetail", resPayload.response);
      console.log("go home");
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
