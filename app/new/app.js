// src/App.js

import React, { useState } from 'react';
import './App.css';

const App = () => {
  const [name, setName] = useState('');
  const [password, setPassword] = useState('');
  const [user, setUser] = useState(null);

  const handleLogin = async () => {
    try {
      const response = await fetch('http://localhost:3001/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ name, password }),
      });

      if (!response.ok) {
        throw new Error('Login failed');
      }

      const userData = await response.json();
      setUser(userData);
    } catch (error) {
      console.error('Login error:', error);
    }
  };

  return (
    <div className="App">
      <h1>Login Page</h1>
      <label>
        Name:
        <input type="text" value={name} onChange={(e) => setName(e.target.value)} />
      </label>
      <br />
      <label>
        Password:
        <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} />
      </label>
      <br />
      <button onClick={handleLogin}>Login</button>

      {user && (
        <div>
          <h2>User Details:</h2>
          <p>Name: {user.name}</p>
          <p>Password: {user.password}</p>
        </div>
      )}
    </div>
  );
};

export default App;