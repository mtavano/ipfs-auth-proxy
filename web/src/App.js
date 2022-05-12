import './App.css';
import { BrowserRouter as Router, Navigate, Route, Routes, useNavigate } from 'react-router-dom'
import { Login, MainPage } from './Components'

function App() {
  return (
    <div className="App">
      <Router>
        <Routes>
          <Route exact={true} path='/' element={<Login />} />
          <Route path='/fleek' element={<MainPage />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
