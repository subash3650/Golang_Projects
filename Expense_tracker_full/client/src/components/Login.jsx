import React, { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
// import "./Login.css";
import axios from "axios";

const API_URL = import.meta.env.VITE_API_URL || "http://localhost:5000";
const Login = () => {
  const [form, setForm] = useState({ email: "", password: "" });
  const navigate = useNavigate();

  const handleChange = e => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async e => {
    e.preventDefault();
    if (!form.email || !form.password) return;
    try{
      const res = await axios.get(
        `${API_URL}/login`
      );
    }catch(err){
      console.error("Login error:", err);
      alert("Login failed. Please check your credentials.");
      return; 
    }
    navigate("/tracker");
  };

  return (
    <div className="auth-container">
      <form className="auth-form" onSubmit={handleSubmit}>
        <h2>Login</h2>
        <input
          type="email"
          name="email"
          placeholder="Email"
          value={form.email}
          onChange={handleChange}
          required
        />
        <input
          type="password"
          name="password"
          placeholder="Password"
          value={form.password}
          onChange={handleChange}
          required
        />
        <button type="submit">Login</button>
        <p>
          Don't have an account?{" "}
          <Link to="/signup" className="auth-link">Sign up</Link>
        </p>
      </form>
    </div>
  );
};

export default Login;