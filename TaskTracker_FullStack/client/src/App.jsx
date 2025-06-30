import React, { useEffect, useState } from "react";
import './App.css';
import axios from "axios";

const API_URL = import.meta.env.VITE_API_URL || "http://localhost:8080";

export default function App() {
  const [tasks, setTasks] = useState([]);
  const [title, setTitle] = useState("");
  const [desc, setDesc] = useState("");

  const fetchTasks = async () => {
    try {
      const res = await axios.get(`${API_URL}/tasks`);
      setTasks(res.data);
    } catch (e) {
      console.error("Fetch error:", e);
    }
  };

  useEffect(() => { fetchTasks() }, []);

  const createTask = async () => {
    if (!title.trim()) return;
    try {
      await axios.post(`${API_URL}/tasks`, { title, description: desc });
      setTitle(""); setDesc("");
      fetchTasks();
    } catch (e) {
      console.error("Create error:", e);
    }
  };

  const toggleComplete = async (id, completed) => {
    try {
      const task = tasks.find(t => t.id === id);
      await axios.put(`${API_URL}/tasks/${id}`, {
        title: task.title,
        description: task.description,
        completed: !completed,
      });
      fetchTasks();
    } catch (e) {
      console.error("Toggle error:", e);
    }
  };

  const deleteTask = async id => {
    try {
      await axios.delete(`${API_URL}/tasks/${id}`);
      fetchTasks();
    } catch (e) {
      console.error("Delete error:", e);
    }
  };

  return (
    <div style={{ padding:24 }}>
      <h1>Tasks</h1>
      <input placeholder="Title" value={title}
        onChange={e=>setTitle(e.target.value)} />
      <input placeholder="Description" value={desc}
        onChange={e=>setDesc(e.target.value)} />
      <button onClick={createTask}>Add</button>

      {Array.isArray(tasks) ? (
        <ul>
          {tasks.map(t=>(
            <li key={t.id}>
              <strong>{t.title}</strong> — {t.description}<br/>
              <span>Status: {t.completed ? "✅ Completed" : "❌ Not completed"}</span><br/>
              <span>Created: {new Date(t.createdAt).toLocaleString()}</span><br/>
              <button onClick={()=>toggleComplete(t.id, t.completed)}>
                {t.completed ? "Mark as Incomplete":"Mark as Complete"}
              </button>
              <button onClick={()=>deleteTask(t.id)}>Delete</button>
            </li>
          ))}
        </ul>
      ) : (
        <p>Loading…</p>
      )}
    </div>
  )
}
