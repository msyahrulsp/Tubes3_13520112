import {
  BrowserRouter as Router,
  Routes,
  Route
} from 'react-router-dom';

import './styles/base.scss';

import { Home } from "./Page/Home";
import { AddDNA } from "./Page/AddDNA";
import { CheckDNA } from "./Page/CheckDNA";

function App() {
  return (
    <Router>
      <div>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/add-dna" element={<AddDNA />} />
          <Route path="/check-dna" element={<CheckDNA />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
