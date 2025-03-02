import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import URLS from "../urls";
import { getRequestToAPI } from "../services/api";
import store from "../store/store"
import { updateUser } from "../store/actions";

const Login = () => {
  const [userId, setUserId] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const handleLogin = async () => {
    if (!userId.trim()) {
      setError("User ID is required.");
      return;
    }

    setLoading(true);
    setError("");

    try {
      const url = `${URLS.login}?userId=${userId}`;
      const response = await getRequestToAPI(url);

      if (response.status === 200 && response.data) {
        const userData = response.data.data; 
        localStorage.setItem("userId",userData.userId)
        store.dispatch(updateUser(userData))
        navigate("/dashboard");
      }
    } catch (err) {
      setError("Failed to fetch user details. Please try again.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="container">
      <h2>Login</h2>
      {error && <p className="error">{error}</p>}
      <input
        type="text"
        placeholder="Enter your ID"
        value={userId}
        onChange={(e) => setUserId(e.target.value)}
      />
      <button onClick={handleLogin} disabled={loading}>
        {loading ? "Logging in..." : "Enter Game"}
      </button>
    </div>
  );
};

export default Login;
