import React, { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
// import "./Login.css";
import axios from "axios";
const API_URL = import.meta.env.VITE_API_URL || "http://localhost:5000";

const Signup = () => {
  const [form, setForm] = useState({ email: "", password: "", confirm: "" });
  const navigate = useNavigate();

  const handleChange = e => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async e => {
    e.preventDefault();
    try{
      const res = await axios.post(`${API_URL}/signup`, {
        email: form.email,
        password: form.password,
        confirm_password: form.confirm,
      })
    }catch(err){
      console.error("Signup error:", err);
      alert("Signup failed. Please try again.");
      return; 
    }
    navigate("/login");
  };

  return (
    <div className="auth-container">
      <form className="auth-form" onSubmit={handleSubmit}>
        <h2>Sign Up</h2>
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
        <input
          type="password"
          name="confirm"
          placeholder="Confirm Password"
          value={form.confirm}
          onChange={handleChange}
          required
        />
        <button type="submit">Sign Up</button>
        <p>
          Already have an account?{" "}
          <Link to="/login" className="auth-link">Log in</Link>
        </p>
      </form>
    </div>
  );
};

export default Signup;