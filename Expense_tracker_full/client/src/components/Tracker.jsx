import React, {useState, useEffect} from "react";
import axios from "axios";
import "./Tracker.css";

const API_URL = import.meta.env.VITE_API_URL || "http://localhost:5000";

const Tracker = () =>{
    const [expenses, setExpenses] = useState([]);
  const [form, setForm] = useState({
    title: "",
    amount: "",
    category: "",
    description: "",
  });
  const [editId, setEditId] = useState(null);
  const [categories, setCategories] = useState([]);
  const [filterCategory, setFilterCategory] = useState("");

  const fetchExpenses = async (category = "") => {
    try {
      const res = await axios.get(
        `${API_URL}/expense${category ? `?category=${encodeURIComponent(category)}` : ""}`
      );
      setExpenses(Array.isArray(res.data) ? res.data : []);
    } catch (e) {
      setExpenses([]);
      console.error("Fetch error:", e);
    }
  };

  const fetchCategories = async () => {
    try {
      const res = await axios.get(`${API_URL}/categories`);
      setCategories(Array.isArray(res.data) ? res.data : []);
    } catch (e) {
      setCategories([]);
    }
  };

  useEffect(() => {
    fetchExpenses(filterCategory);
    fetchCategories();
  }, [filterCategory]);

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (!form.title || !form.amount) return;
    try {
      if (editId) {
        await axios.put(`${API_URL}/expense/${editId}`, {
          ...form,
          amount: parseFloat(form.amount),
        });
        setEditId(null);
      } else {
        await axios.post(`${API_URL}/expense`, {
          ...form,
          amount: parseFloat(form.amount),
        });
      }
      setForm({ title: "", amount: "", category: "", description: "" });
      fetchExpenses(filterCategory);
      fetchCategories();
    } catch (e) {
      alert("Error: " + (e.response?.data?.error || e.message));
    }
  };

  const handleEdit = (exp) => {
    setEditId(exp.id || exp._id);
    setForm({
      title: exp.title,
      amount: exp.amount,
      category: exp.category,
      description: exp.description,
    });
  };

  const handleDelete = async (id) => {
    if (!window.confirm("Delete this expense?")) return;
    try {
      await axios.delete(`${API_URL}/expense/${id}`);
      fetchExpenses(filterCategory);
      fetchCategories();
    } catch (e) {
      alert("Delete failed");
    }
  };

  return (
    <div className="app-container">
      <h1>Expense Tracker</h1>
      <form onSubmit={handleSubmit} style={{ marginBottom: 24 }}>
        <input
          name="title"
          placeholder="Title"
          value={form.title}
          onChange={handleChange}
          required
        />
        <input
          name="amount"
          type="number"
          step="0.01"
          placeholder="Amount"
          value={form.amount}
          onChange={handleChange}
          required
        />
        <input
          list="category-list"
          name="category"
          placeholder="Category"
          value={form.category}
          onChange={handleChange}
          required
        />
        <datalist id="category-list">
          {categories.map((cat) => (
            <option value={cat} key={cat} />
          ))}
        </datalist>
        <textarea
          name="description"
          placeholder="Description"
          value={form.description}
          onChange={handleChange}
          rows={2}
        />
        <button type="submit">{editId ? "Update" : "Add"}</button>
        {editId && (
          <button
            type="button"
            onClick={() => {
              setEditId(null);
              setForm({ title: "", amount: "", category: "", description: "" });
            }}
          >
            Cancel
          </button>
        )}
      </form>
      <div style={{ marginBottom: 24 }}>
        <label style={{ marginRight: 8, fontWeight: "bold" }}>Filter by Category:</label>
        <select
          value={filterCategory}
          onChange={(e) => setFilterCategory(e.target.value)}
        >
          <option value="">All</option>
          {categories.map((cat) => (
            <option value={cat} key={cat}>
              {cat}
            </option>
          ))}
        </select>
      </div>
      <ul>
        {(Array.isArray(expenses) ? expenses : []).map((exp) => (
          <li key={exp.id || exp._id}>
            <strong>{exp.title}</strong> — ₹{exp.amount} <br />
            <span>Category: {exp.category || "N/A"}</span>
            <span>{exp.description}</span>
            <span>
              Date:{" "}
              {exp.date
                ? new Date(exp.date).toLocaleString()
                : "N/A"}
            </span>
            <div>
              <button onClick={() => handleEdit(exp)}>Edit</button>
              <button onClick={() => handleDelete(exp.id || exp._id)}>
                Delete
              </button>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default Tracker;