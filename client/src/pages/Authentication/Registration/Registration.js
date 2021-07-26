import React, { useState } from "react";
import { withRouter } from "react-router";

import { registerRequest } from "../../../services/api-service";
import { setItemToLS } from "../../../services/storage-service";


import "./Registration.css";

function Registration(props) {
  const [registrationErrorMessage, setErrorMessage] = useState(null);
  const [username, setUsername] = useState(null);
  const [password, setPassword] = useState(null);

  const handlePasswordChange = (event) => {
    setPassword(event.target.value);
  };

  const handleUsernameChange = (event) => {
    setUsername(event.target.value);
  };

  const registerUser = async () => {
    // verify 
    var usernameRegex = /^[a-zA-Z0-9]+$/;
    if (!usernameRegex.test(username)) {
      setErrorMessage(
        "Your username is not valid. Only characters A-Z, a-z and 0-9 are  acceptable."
      );
      return;
    }
    
    // request to server
    const resPayload = await registerRequest(username, password);
    if(resPayload.code === 200) {
      setItemToLS("userDetails", resPayload.response);
      console.log("go home")
    } else {
      setErrorMessage(resPayload.message);
    }
  };

  return (
    <div className="app__register-container">
      <div className="app__form-row">
        <label>Username:</label>
        <input type="email" className="email" onChange={handleUsernameChange} />
      </div>

      <div className="app__form-row">
        <label>Password:</label>
        <input
          type="password"
          className="password"
          onChange={handlePasswordChange}
        />
      </div>

      <div className="app__form-row">
        <span className="error-message">
          {registrationErrorMessage ? registrationErrorMessage : ""}
        </span>
      </div>
      <div className="app__form-row">
        <button onClick={registerUser}>Registration</button>
      </div>
    </div>
  );
}

export default withRouter(Registration);
