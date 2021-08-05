import React, { useState } from "react";
import { withRouter } from "react-router";
import { Form, Button } from "react-bootstrap";

import { loginRequest } from "../../../services/api-service";
import { setItemToSS } from "../../../services/storage-service";

function Login(props) {
  const [loginErrorMessage, setErrorMessage] = useState(null);
  const [userEmptyError, setUsernameError] = useState(null);
  const [passwordEmptyError, setPasswordError] = useState(null);
  const [username, setUsername] = useState(null);
  const [password, setPassword] = useState(null);

  const handleUsernameChange = (event) => {
    setUsername(event.target.value);
  };

  const handlePasswordChange = (event) => {
    setPassword(event.target.value);
  };

  const loginUser = async () => {
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
      const resPayload = await loginRequest(username, password);
      if (resPayload.code === 200) {
        setItemToSS("userDetail", resPayload.response);
        console.log("go home");
        props.history.push(`/home`);
      } else {
        setErrorMessage(resPayload.message);
      }
    }
  };

  return (
    <>
      <Form className="auth-form">
        <Form.Group className="mb-3" controlId="loginUsername">
          <Form.Control
            type="text"
            name="username"
            placeholder="Enter username"
            onChange={handleUsernameChange}
          />
          <Form.Text style={{ color: "red" }}>
            {userEmptyError ? userEmptyError : ""}
          </Form.Text>
        </Form.Group>

        <Form.Group className="mb-3" controlId="loginPassword">
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

        <Form.Group className="mb-3" controlId="loginError">
          <Form.Text style={{ color: "red" }}>
            {loginErrorMessage ? loginErrorMessage : ""}
          </Form.Text>
        </Form.Group>
        <Button block variant="primary" onClick={loginUser}>
          Login
        </Button>
      </Form>
    </>
  );
}

export default withRouter(Login);
