// server.js

const express = require('express');
const mongoose = require('mongoose');
const redis = require('redis');
const bodyParser = require('body-parser');

const app = express();
const port = 3001;

// Connect to MongoDB (Make sure MongoDB is running)
mongoose.connect('mongodb://localhost:27017/login-app', { useNewUrlParser: true, useUnifiedTopology: true });
const db = mongoose.connection;

db.on('error', console.error.bind(console, 'MongoDB connection error:'));
db.once('open', () => console.log('Connected to MongoDB'));

// Connect to Redis (Make sure Redis is running)
const redisClient = redis.createClient();
redisClient.on('connect', () => console.log('Connected to Redis'));
redisClient.on('error', err => console.error('Redis error:', err));

// Middleware to parse JSON
app.use(bodyParser.json());

// MongoDB Schema
const userSchema = new mongoose.Schema({
  name: String,
  password: String,
});

const User = mongoose.model('User', userSchema);

// Express endpoint to handle login
app.post('/login', async (req, res) => {
  const { name, password } = req.body;

  // Check if user exists in Redis cache
  redisClient.get(name, async (err, cachedUser) => {
    if (err) {
      console.error('Redis error:', err);
      return res.status(500).send('Internal Server Error');
    }

    if (cachedUser) {
      // If user exists in cache, return the cached data
      console.log('User found in Redis cache');
      return res.json(JSON.parse(cachedUser));
    } else {
      // If user not in cache, check MongoDB
      try {
        const user = await User.findOne({ name, password });
        if (user) {
          // If user found in MongoDB, store in Redis cache and return the data
          redisClient.setex(name, 3600, JSON.stringify(user)); // Cache for 1 hour
          console.log('User stored in Redis cache');
          return res.json(user);
        } else {
          // If user not found in MongoDB
          return res.status(404).send('User not found');
        }
      } catch (error) {
        console.error('MongoDB error:', error);
        return res.status(500).send('Internal Server Error');
      }
    }
  });
});

// Start the server
app.listen(port, () => {
    console.log( 'Server is running on http://localhost:${port}');
  });