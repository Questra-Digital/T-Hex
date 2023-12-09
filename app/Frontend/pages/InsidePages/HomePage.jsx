import React from "react";
import { Link } from "react-router-dom";
import Footer1 from "../Components/Footer1";


import styles from "./HomePage.module.css";

export default function HomePage() {
  return (
    <div>
      <main>

        
        <section className={styles.hero}>
          <div className={styles.heroContainer}>
            <h1 className={styles.heroTitle}>Web Test Hub</h1>
            <p className={styles.heroDescription}>
              
WebTestHub is a cutting-edge test automation software designed as an affordable, robust, and user-centric platform specifically tailored to streamline and enhance automated testing processes. Born from the need to address limitations in existing solutions, particularly the constraints of the "Scalable Web Testing Platform," WebTestHub focuses on revolutionizing automated testing by enabling multiple parallel test executions through the utilization of multiple containers.
            </p>
            <Link to="/CloneRepository">
              <button className={styles.heroButton}>Get Started</button>
            </Link>
          </div>
        </section>

        <section className={styles.selenium}>
          <div className={styles.seleniumContainer}>
            <div className={styles.seleniumText}>
              <h2 className={styles.seleniumTitle}>Selenium</h2>
              <p className={styles.seleniumDescription}>
                Selenium is a powerful open-source framework for automating web
                browsers. It provides a suite of tools and libraries to interact
                with web elements, simulate user actions, and perform automated
                testing on different web browsers and platforms.
              </p>
            </div>
            <img src="/sel3.jpg" alt="Selenium" width={600} height={400} />
          </div>
        </section>
      </main>
      <Footer1/>
    </div>

  
  );
}
