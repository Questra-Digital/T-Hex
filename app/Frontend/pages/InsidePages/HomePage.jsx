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
            <h1 className={styles.heroTitle}>Scalable Web Testing Platform</h1>
            <p className={styles.heroDescription}>
              Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed ut
              sapien nulla. Vestibulum eu mauris velit. Phasellus nec justo
              urna. Morbi eleifend elit quis ligula facilisis, vitae iaculis
              dolor dignissim. Integer quis dolor vel odio consequat mattis et
              et diam.
            </p>
            <Link to="/CloneRepository">
              <button className={styles.heroButton}>Get Started</button>
            </Link>
          </div>
        </section>

        <section className={styles.features}>
          <div className={styles.featureContainer}>
            <div className={styles.featureItem}>
              <h2 className={styles.featureTitle}>Cross-web Testing</h2>
              <img
                src="/cwt.webp"
                alt="Cross Web Testing"
                width={400}
                height={60}
              />
              <span> -----------------------------------------</span>
              <p className={styles.featureDescription}>
                Test your web application across multiple browsers, platforms,
                and devices to ensure consistent functionality and performance.
              </p>
            </div>
            <div className={styles.featureItem}>
              <h2 className={styles.featureTitle}>Screenshots</h2>
              <img src="/ss.png" alt="Screenshots" width={400} height={60} />
              <span> -----------------------------------------</span>
              <p className={styles.featureDescription}>
                Capture screenshots of your web application on different
                browsers and devices for visual comparison.
              </p>
            </div>
            <div className={styles.featureItem}>
              <h2 className={styles.featureTitle}>Videos of Tests Ran</h2>
              <img src="/vid.png" alt="Videos" width={400} height={60} />
              <span> -----------------------------------------</span>
              <p className={styles.featureDescription}>
                Record videos of your test runs to review user interactions,
                identify UI issues, and improve your web application.
              </p>
            </div>
            <div className={styles.featureItem}>
              <h2 className={styles.featureTitle}>Logs</h2>
              <img src="/logs.jpg" alt="Logs" width={300} height={30} />
              <span> -----------------------------------------</span>
              <p className={styles.featureDescription}>
                Access detailed logs of your test runs to analyze errors,
                diagnose issues, and optimize your web application.
              </p>
            </div>
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
