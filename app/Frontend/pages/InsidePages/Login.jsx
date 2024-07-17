import React, { useState } from 'react';
import style from './LoginScreen.module.css';
import { Link } from 'react-router-dom';

const LoginScreen = () => {
    // State for login form
    const [loginEmail, setLoginEmail] = useState('');
    const [loginPassword, setLoginPassword] = useState('');

    // State for signup form
    const [signupUsername, setSignupUsername] = useState('');
    const [signupEmail, setSignupEmail] = useState('');
    const [signupPassword, setSignupPassword] = useState('');

    // Login form handlers
    const handleLoginEmailChange = (e) => {
        setLoginEmail(e.target.value);
    };

    const handleLoginPasswordChange = (e) => {
        setLoginPassword(e.target.value);
    };

    const handleLogin = (e) => {
        e.preventDefault();
        // Perform login action
        console.log("Login Email:", loginEmail);
        console.log("Login Password:", loginPassword);
        setLoginEmail('');
        setLoginPassword('');
    };

    // Signup form handlers
    const handleSignupUsernameChange = (e) => {
        setSignupUsername(e.target.value);
    };

    const handleSignupEmailChange = (e) => {
        setSignupEmail(e.target.value);
    };

    const handleSignupPasswordChange = (e) => {
        setSignupPassword(e.target.value);
    };

    const handleSignup = (e) => {
        e.preventDefault();
        // Perform signup action
        console.log("Signup Username:", signupUsername);
        console.log("Signup Email:", signupEmail);
        console.log("Signup Password:", signupPassword);
        setSignupUsername('');
        setSignupEmail('');
        setSignupPassword('');
    };

    return (
        
        <div className={style.fullScreen} style={{ backgroundImage: `url('b.jpg')`, backgroundSize: 'cover' }}>


            {/* Your background elements */}
            <img src="logg.png" alt="Your Logo" class="logoImage">
            </img>
            <div className={style.loginContainer}>
                
                <div className={style.loginContent}>
               
                    <h1 className={style.welcomeText}>Welcome Back</h1>
                    {/* Login form */}
              
                    <form className={style.loginForm} onSubmit={handleLogin}>
                        <input 
                            type="text" 
                            placeholder="Username" 
                            className={style.inputField} 
                            value={loginEmail} 
                            onChange={handleLoginEmailChange} 
                        />
                        <input 
                            type="password" 
                            placeholder="Password" 
                            className={style.inputField} 
                            value={loginPassword} 
                            onChange={handleLoginPasswordChange} 
                        />
                        <Link to="/CloneRepository">
                        <button type="submit" className={style.loginButton}>Login</button>
                        </Link>
                    </form>
                    {/* Signup link */}
                    <Link to="/Signup" className={style.signupLink}>Don't have an account? Sign up here</Link>
                </div>
            </div>
        </div>
    );
};

export default LoginScreen;
