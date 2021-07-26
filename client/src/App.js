import './App.css';

import {
  BrowserRouter as Router,
  Route,
  Switch
} from "react-router-dom";

import Authentication from './pages/Authentication/Authentication';

function App() {
  return (
    <Router>
      <Switch>
        <Route path="/" exact component={Authentication} />
      </Switch>
    </Router>
  );
}

export default App;
