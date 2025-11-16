"use client";

import { useState } from "react";
import { motion, AnimatePresence } from "framer-motion";
import styles from "@/styles/pages/Login.module.scss";
import { outfit } from "@/fonts/fonts";
import Image from "next/image";
import { useSnackbar } from "@/contexts/SnackbarContext";
import { authService } from "@/services/authService";
import { setUserData } from "@/store/userData";
import { useDispatch } from "react-redux";
import { useRouter } from "next/navigation";

const featuresList = [
    {
        title: "5000+ Browsers, OS, & Real Devices",
        description: "Run manual and automated tests of web and mobile apps across 5000+ different browsers, real devices, and OS environments."
    },
    {
        title: "70% Faster Test Execution",
        description: "T-Hex's HyperExecute is up to 70% faster than any traditional cloud-based Selenium Grid"
    },
    {
        title: "AI-Native Test Intelligence",
        description: "Test smarter with AI-native test insights to help you anticipate and mitigate future issues before they take root."
    }
];

const COMPANY_LOGOS = [
    { src: "/Icons/Microsoft.svg", alt: "Microsoft", width: 100, height: 100 },
    { src: "/Icons/Vimeo.svg", alt: "Vimeo", width: 100, height: 100 },
    { src: "/Icons/Nvidia.svg", alt: "Nvidia", width: 100, height: 100 },
    { src: "/Icons/Telstra.svg", alt: "Telstra", width: 100, height: 100 },
    { src: "/Icons/Rubrik.svg", alt: "Rubrik", width: 100, height: 100 },
];

const LOGO_DISPLAY_SIZE = { width: 120, height: 80 };

export default function LoginPage() {
    const [email, setEmail] = useState("");
    const [otp, setOtp] = useState("");
    const [emailError, setEmailError] = useState("");
    const [otpError, setOtpError] = useState("");
    const [isLoading, setIsLoading] = useState(false);
    const [showOTPForm, setShowOTPForm] = useState(false);
    const [userEmail, setUserEmail] = useState("");
    const { showSnackbar } = useSnackbar();
    const dispatch = useDispatch();
    const router = useRouter();
    const handleEmailSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        
        // Validate email
        const emailValidation = authService.validateEmail(email);
        if (!emailValidation.isValid) {
            setEmailError(emailValidation.error || "");
            return;
        }
        
        setEmailError("");
        setIsLoading(true);

        try {
            const response = await authService.signup(email);
            
            if (response.success) {
                setUserEmail(email);
                setShowOTPForm(true);
                showSnackbar("OTP sent to your email address", "success");
            } else {
                showSnackbar(response.message || "Failed to send OTP", "error");
            }
        } catch (error) {
            showSnackbar("An error occurred. Please try again.", "error");
        } finally {
            setIsLoading(false);
        }
    };

    const handleOTPSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        
        // Validate OTP
        const otpValidation = authService.validateOTP(otp);
        if (!otpValidation.isValid) {
            setOtpError(otpValidation.error || "");
            return;
        }
        
        setOtpError("");
        setIsLoading(true);

        try {
            const response = await authService.verifyOTP(userEmail, otp);
            
            if (response.success) {

                const userId = response.data?.user_id || 0
                dispatch(setUserData({ user_id: userId as number }));

                showSnackbar("Login successful!", "success");
                
                router.push('/getting_started');
                console.log("Login successful:", response.data);
            } else {
                showSnackbar(response.message || "Invalid OTP", "error");
            }
        } catch (error) {
            showSnackbar("An error occurred. Please try again.", "error");
        } finally {
            setIsLoading(false);
        }
    };

    const handleBackToEmail = () => {
        setShowOTPForm(false);
        setOtp("");
        setOtpError("");
        setEmail("");
        setEmailError("");
    };

    return (
        <div className={`${styles.loginPage} ${outfit.variable}`}>
            {/* Left Column - Info Section */}
            <div className={styles.infoSection}>
                <p className={styles.title}>T-Hex</p>
                <p className={styles.subtitle}>With your new T-Hex Account, you get:</p>
                <div className={styles.featuresList}>
                    {featuresList.map((feature) => (
                        <div className={styles.featureItem} key={feature.title}>
                            <div className={styles.checkmark}>
                                <Image src="/Icons/login/tick.svg" alt="Tick Mark" width={20} height={20} />
                            </div>
                            <div className={styles.featureContent}>
                                <h2 className={styles.featureTitle}>{feature.title}</h2>
                                <p className={styles.featureDescription}>{feature.description}</p>
                            </div>
                        </div>
                    ))}
                </div>
                <div className={styles.companyLogos}>
                    <p className={styles.companyLogosTitle}>Trusted by users globally</p>
                    <div className={styles.companyIcons}>
                        {COMPANY_LOGOS.map((logo) => {
                            const classKey = logo.alt.toLowerCase();
                            return (
                                <div key={classKey} className={styles.companyIcon}>
                                    <Image
                                        src={logo.src}
                                        alt={logo.alt}
                                        width={LOGO_DISPLAY_SIZE.width}
                                        height={LOGO_DISPLAY_SIZE.height}
                                        className={styles[classKey]}
                                    />
                                </div>
                            );
                        })}
                    </div>
                </div>
            </div>

            {/* Right Column - Form Section */}
            <div className={styles.formSection}>
                <h2 className={styles.formTitle}>
                    {showOTPForm ? "Verify Your Email" : "Get Started for Free"}
                </h2>
                
                {!showOTPForm && (
                    <>
                        {/* Social Login Buttons */}
                        <div className={styles.socialButtons}>
                            <button className={styles.socialButton}>
                                <Image src="/Icons/login/Google.svg" alt="Google" width={20} height={20} />
                                Sign Up with Google
                            </button>
                            <button className={styles.socialButton}>
                                <Image src="/Icons/login/Github.svg" alt="GitHub" width={20} height={20} />
                                Sign Up with GitHub
                            </button>
                        </div>

                        {/* Divider */}
                        <div className={styles.divider}>
                            <span className={styles.dividerText}>OR</span>
                        </div>
                    </>
                )}

                {/* Email Form */}
                <AnimatePresence mode="wait">
                    {!showOTPForm ? (
                        <motion.form
                            key="email-form"
                            initial={{ opacity: 0, x: 20 }}
                            animate={{ opacity: 1, x: 0 }}
                            exit={{ opacity: 0, x: -20 }}
                            transition={{ duration: 0.3 }}
                            className={styles.signupForm}
                            onSubmit={handleEmailSubmit}
                        >
                            <div className={styles.inputGroup}>
                                <label htmlFor="email" className={styles.inputLabel}>Email Address</label>
                                <input 
                                    type="email" 
                                    id="email" 
                                    className={`${styles.inputField} ${emailError ? styles.inputError : ''}`}
                                    placeholder="Enter your email address"
                                    value={email}
                                    onChange={(e) => {
                                        setEmail(e.target.value);
                                        if (emailError) setEmailError("");
                                    }}
                                    disabled={isLoading}
                                />
                                {emailError && <span className={styles.errorMessage}>{emailError}</span>}
                            </div>
                            <button 
                                type="submit" 
                                className={styles.signupButton}
                                disabled={isLoading}
                            >
                                {isLoading ? (
                                    <span className={styles.loadingSpinner}>Sending...</span>
                                ) : (
                                    "Send OTP"
                                )}
                            </button>
                        </motion.form>
                    ) : (
                        <motion.form
                            key="otp-form"
                            initial={{ opacity: 0, x: 20 }}
                            animate={{ opacity: 1, x: 0 }}
                            exit={{ opacity: 0, x: -20 }}
                            transition={{ duration: 0.3 }}
                            className={styles.signupForm}
                            onSubmit={handleOTPSubmit}
                        >
                            <div className={styles.otpInfo}>
                                <p className={styles.otpInfoText}>
                                    We've sent a 6-digit code to <strong>{userEmail}</strong>
                                </p>
                                <button 
                                    type="button" 
                                    className={styles.backButton}
                                    onClick={handleBackToEmail}
                                    disabled={isLoading}
                                >
                                    ← Change email
                                </button>
                            </div>
                            
                            <div className={styles.inputGroup}>
                                <label htmlFor="otp" className={styles.inputLabel}>Enter OTP</label>
                                <input 
                                    type="text" 
                                    id="otp" 
                                    className={`${styles.inputField} ${styles.otpInput} ${otpError ? styles.inputError : ''}`}
                                    placeholder="000000"
                                    value={otp}
                                    onChange={(e) => {
                                        const value = e.target.value.replace(/\D/g, '').slice(0, 6);
                                        setOtp(value);
                                        if (otpError) setOtpError("");
                                    }}
                                    disabled={isLoading}
                                    maxLength={6}
                                />
                                {otpError && <span className={styles.errorMessage}>{otpError}</span>}
                            </div>
                            
                            <button 
                                type="submit" 
                                className={styles.signupButton}
                                disabled={isLoading || otp.length !== 6}
                            >
                                {isLoading ? (
                                    <span className={styles.loadingSpinner}>Verifying...</span>
                                ) : (
                                    "Verify OTP"
                                )}
                            </button>
                        </motion.form>
                    )}
                </AnimatePresence>

                {/* Terms and Policies */}
                <p className={styles.termsText}>
                    By clicking Sign Up or registering in through a third party you accept the T-Hex{" "}
                    <a href="#" className={styles.link}>Terms of Service</a> and acknowledge the{" "}
                    <a href="#" className={styles.link}>Privacy Policy</a> and{" "}
                    <a href="#" className={styles.link}>Cookie Policy</a>
                </p>

                {/* Login Link */}
                <p className={styles.loginPrompt}>
                    Already Have An Account? <a href="#" className={styles.link}>Sign In</a>
                </p>
            </div>
        </div>
    );
}