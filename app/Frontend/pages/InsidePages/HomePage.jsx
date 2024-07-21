import React from "react";
import { Link } from "react-router-dom";
import Footer1 from "../Components/Footer1";


import styles from "./HomePage.module.css";

export default function HomePage() {
  const videoSrc = "/vid.mp4";
  return (
    <div>
      <main>
        <section className={styles.hero}>
          <div className={styles.heroVideo}>
            <video autoPlay muted loop className={styles.video}>
              <source src={videoSrc} type="video/mp4" />
            </video>
          </div>
          <div className={styles.heroContainer}>
          <h1 className={styles.heroTitle}>Web Test Hub</h1>
            <p className={styles.heroDescription} padding-top="50px" >
            WebTestHub is a cutting-edge test automation software designed 
            as an affordable, robust, and user-centric platform specifically
            tailored to streamline and enhance automated testing processes.
            WebTestHub focuses on revolutionizing automated testing by enabling
            multiple parallel test executions through the utilization of multiple
            containers.
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

          <div className={styles.automation}>
          <h1 className={styles.automationTitle}>Automation</h1>
          <div className={styles.image2}>
          <img src="/auto.png"  alt="Selenium" width={900} height={700} />
          
        </div>
       
        </div>
        <p className={styles.para}>
        Test automation is the process of using software tools to execute tests. These tests aim to replicate manual testing tasks, which would typically be performed by a human tester. Test automation involves creating test scripts that instruct the software on what actions to take and what results to expect.

There are a number of benefits to test automation, including:

Reduced costs: Automating tests can save time and money by reducing the need for manual testing.
Improved quality: Test automation can help to improve the quality of software by catching bugs that might be missed during manual testing.
Increased accuracy: Automated tests are typically more accurate than manual tests, as they are less prone to human error.
Improved productivity: Test automation can free up testers to focus on more complex tasks.
Faster development and delivery: Automated tests can be run quickly and frequently, which can help to speed up the development and delivery of software.
        </p>
        </section>
      </main>
      <Footer1/>
    </div>

  
  );
}
