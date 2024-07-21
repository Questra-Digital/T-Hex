import React, { useState } from 'react';
import style from './LoginScreen.module.css';
import { Link } from 'react-router-dom';
import axios from 'axios'; // Import Axios for making HTTP requests

const SignupScreen = () => {
    const [formData, setFormData] = useState({
        username: '',
        password: '',
        email: '',
        phoneNumber: ''
    });
    const [message, setMessage] = useState('');

    const handleChange = (e) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const response = await axios.post('/signup', formData);
            console.log(response.data);
            setMessage('User signed up successfully'); // Set message when signup is successful
            // Optionally, you can redirect the user to another page after successful signup
        } catch (error) {
            console.error('Signup failed:', error);
            setMessage('Signup failed. Please try again.'); // Set message when signup fails
            
        }
    };

    return (
        <div className={style.fullScreen} style={{ backgroundImage: `url('b.jpg')`, backgroundSize: 'cover' }}>
            {/* Your background elements */}
            <img src="logg.png" alt="Your Logo" className="logoImage" />
            <div className={style.loginContainer}>
                <div className={style.loginContent}>
                    <h1 className={style.welcomeText}>Welcome To Web Test Hub</h1>
                    {/* Signup form */}
                    <form className={style.loginForm} onSubmit={handleSubmit}>
                        <input
                            type="text"
                            name="username"
                            placeholder="Username"
                            className={style.inputField}
                            value={formData.username}
                            onChange={handleChange}
                            style={{ color: 'black' }} // Set text color to black
                        />
                        <input
                            type="password"
                            name="password"
                            placeholder="Password"
                            className={style.inputField}
                            value={formData.password}
                            onChange={handleChange}
                            style={{ color: 'black' }} // Set text color to black
                        />
                        <input
                            type="email"
                            name="email"
                            placeholder="Email"
                            className={style.inputField}
                            value={formData.email}
                            onChange={handleChange}
                            style={{ color: 'black' }} // Set text color to black
                        />
                        <input
                            type="tel" textcolor="black"
                            name="phoneNumber"
                            placeholder="Phone Number"
                            className={style.inputField}
                            value={formData.phoneNumber}
                            onChange={handleChange}
                            style={{ color: 'black' }} // Set text color to black
                            
                        />
                        <Link to="/SubscriptionPlans">
                        <button type="submit" className={style.loginButton}>Signup</button>
                        {/* Display message */}
                        {message && <p>{message}</p>}
                        </Link>
                    </form>
                    {/* Signup link */}
                    <Link to="/Login" className={style.signupLink}>Already have an account? Sign In here</Link>
                </div>
            </div>
        </div>
    );
};

export default SignupScreen;
