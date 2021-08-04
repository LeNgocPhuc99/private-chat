import React, { useState } from "react";
import { withRouter } from "react-router";
import { DebounceInput } from "react-debounce-input";
import { Form, Button } from "react-bootstrap";

import { registerRequest } from "../../../services/api-service";
import { setItemToSS } from "../../../services/storage-service";

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
    if (resPayload.code === 200) {
      setItemToSS("userDetail", resPayload.response);
      props.history.push(`/home`);
      console.log("go home");
    } else {
      setErrorMessage(resPayload.message);
    }
  };

  return (
    <>
      <Form className="auth-form">
        <Form.Group className="mb-3" controlId="formUsername">
          <DebounceInput
            className="form-control"
            placeholder="Enter username"
            minLength={2}
            maxLength={10}
            debounceTimeout={300}
            onChange={handleUsernameChange}
          />
        </Form.Group>

        <Form.Group className="mb-3" controlId="formPassword">
          <Form.Control
            type="password"
            name="password"
            placeholder="Password"
            onChange={handlePasswordChange}
          />
        </Form.Group>
        <Button block variant="primary" onClick={registerUser}>
          Registration
        </Button>
      </Form>
    </>
  );
}

export default withRouter(Registration);
