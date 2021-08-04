import React from "react";
import { Tabs, Tab } from "react-bootstrap";
import "./Authentication.css";

import Login from "./Login/Login";
import Registration from "./Registration/Registration";

function Authentication() {
  
  return (
    <div className="container">
      <div className="authentication-screen">
        <Tabs variant="pills" defaultActiveKey="login">
          <Tab eventKey="login" title="Login">
            <Login />
          </Tab>
          <Tab eventKey="registration" title="Registration">
            <Registration />
          </Tab>
        </Tabs>
      </div>
    </div>
  );
}

export default Authentication;
