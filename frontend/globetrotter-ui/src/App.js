import React from "react";
import { BrowserRouter as Router, Routes, Route, useLocation } from "react-router-dom";
import Dashboard from "./views/Dashboard";
import Login from "./views/Login";
import ResponseSelection from "./views/ResponseSelection";
import Score from "./views/Score";
import Header from "./views/Header";
import "./assets/css/app.css";
import FetchUserDetails from "./services/fetchUserDetails";

function App() {
  return (
    <Router>
      <AppContent />
    </Router>
  );
}

function AppContent() {
  const location = useLocation();
  const isLoggedIn = localStorage.getItem("userId");

  return (
    <>
      <FetchUserDetails/>
      {isLoggedIn && location.pathname !== "/" && <Header />}
      <Routes>
        <Route index path="/" element={<Login />} />
        <Route path="/dashboard" element={<Dashboard />} />
        <Route path="/responses/:cardId" element={<ResponseSelection />} />
        <Route path="/score" element={<Score />} />
      </Routes>
    </>
  );
}

export default App;
