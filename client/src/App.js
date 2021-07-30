import "./App.css";

import { BrowserRouter as Router, Route, Switch } from "react-router-dom";

import Authentication from "./pages/Authentication/Authentication";
import Home from "./pages/Home/Home";

function App() {
  return (
    <Router>
      <Switch>
        <Route path="/" exact component={Authentication} />
        <Route path="/home/" component={Home} />
      </Switch>
    </Router>
  );
}

export default App;
