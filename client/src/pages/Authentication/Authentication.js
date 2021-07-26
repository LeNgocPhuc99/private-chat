import React, { useState } from "react";
import "./Authentication.css";

import Login from "./Login/Login";
import Registration from "./Registration/Registration";

function Authentication() {
  const [activeTab, setTabType] = useState("login");

  const changeTabType = (type) => {
    setTabType(type);
  };

  const getActiveTab = (type) => {
    return type === activeTab ? "active" : "";
  };

  return (
    <React.Fragment>
      <div className="app__authentication-container">
        <div className="authentication__tab-switcher">
          <button
            className={`${getActiveTab("login")} authentication__tab-button`}
            onClick={() => changeTabType("login")}
          >
            Login
          </button>

          <button
            className={`${getActiveTab(
              "registration"
            )} authentication__tab-button`}
            onClick={() => changeTabType("registration")}
          >
            Registration
          </button>
        </div>
        <div className="authentication__tab-viewer">
          {activeTab === "login" ? <Login /> : <Registration />}
        </div>
      </div>
    </React.Fragment>
  );
}

export default Authentication;

