import React, { useState } from "react";
import { withRouter } from "react-router";
import { DebounceInput } from "react-debounce-input";
import { Form, Button } from "react-bootstrap";

import { registerRequest, usernameRegistrationCheckRequest } from "../../../services/api-service";
import { setItemToSS } from "../../../services/storage-service";

function Registration(props) {
  const [registrationErrorMessage, setErrorMessage] = useState(null);
  const [userEmptyError, setUsernameError] = useState(null);
  const [passwordEmptyError, setPasswordError] = useState(null);
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
    var check = true;
    if (username === "" || username === null || username === undefined) {
      check = false;
      setUsernameError("Username can't empty");
    }

    if (password === "" || password === null || password === undefined) {
      check = false;
      setPasswordError("Password can't empty");
    }

    if (check) {
      var usernameRegex = /^[a-zA-Z0-9]+$/;
      if (!usernameRegex.test(username)) {
        setErrorMessage(
          "Your username is not valid. Only characters A-Z, a-z and 0-9 are  acceptable."
        );
        return;
      }

      // check whether username exist or not
      const isValidUsername =  await usernameRegistrationCheckRequest(username);
      if(!isValidUsername.response) {
        setErrorMessage("Your username is registered");
        return;
      }
      // request to server
      const resPayload = await registerRequest(username, password);
      if (resPayload.code === 200) {
        setItemToSS("userDetail", resPayload.response);
        props.history.push(`/home`);
        console.log("go home");
      } else {
        setErrorMessage(resPayload.message);
      }
    }
  };

  return (
    <>
      <Form className="auth-form">
        <Form.Group className="mb-3" controlId="registrationUsername">
          <DebounceInput
            className="form-control"
            placeholder="Enter username"
            minLength={2}
            maxLength={10}
            debounceTimeout={300}
            onChange={handleUsernameChange}
          />

          <Form.Text style={{ color: "red" }}>
            {userEmptyError ? userEmptyError : ""}
          </Form.Text>
        </Form.Group>

        <Form.Group className="mb-3" controlId="registrationPassword">
          <Form.Control
            type="password"
            name="password"
            placeholder="Password"
            onChange={handlePasswordChange}
          />
          <Form.Text style={{ color: "red" }}>
            {passwordEmptyError ? passwordEmptyError : ""}
          </Form.Text>
        </Form.Group>

        <Form.Group className="mb-3" controlId="registrationErr">
          <Form.Text style={{ color: "red" }}>
            {registrationErrorMessage ? registrationErrorMessage : ""}
          </Form.Text>
        </Form.Group>
        <Button block variant="primary" onClick={registerUser}>
          Registration
        </Button>
      </Form>
    </>
  );
}

export default withRouter(Registration);
