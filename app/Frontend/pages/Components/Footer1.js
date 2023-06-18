import React from 'react'
import * as SoIcons from '@fortawesome/free-brands-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import styles from './Footer1.module.css'
import { Link } from 'react-router-dom'


function Footer() {
    return (
        <>
            <div className={styles.fBody}>
                <footer className={styles.footer}>
                    <div className={styles.fContainer}>
                        <div className={styles.fRow}>
                            <div className={styles.fCol}>
                                <h4>Features</h4>
                                <ul>

                                <li><Link to="/CloneRepository" >Clone Git Repository</Link></li>
                                <li><Link to="/CustomizeSettings" >Configure Settings</Link></li>
                                <li><Link to="/CreateDockerFile" >Write Docker File</Link></li>
                                <li><Link to="/RunTest" >start test</Link></li>
                               </ul>
                            </div>
                            <div className={styles.fCol}>
                                <h4>Results</h4>
                                <ul>
                                <li><Link to="/Results" >Test Result</Link></li>
                                <li><Link to="/Videos" >Test Videos</Link></li>
                                <li><Link to="/Logs" >Test Logs</Link></li>
                                <li><Link to="/Screenshots" >Test Screenshots</Link></li>
                                </ul>
                            </div>
                            <div className={styles.fCol}>
                                <h4>Get help</h4>
                                <ul>
                                <li><Link to="/ContactUs" >Contact Us</Link></li>
                                <li><Link to="/ContactUs" >About Us</Link></li>                                
                                </ul>
                            </div>
                            <div className={styles.fCol}>
                                <h4>follow us</h4>
                                <div className={styles.fSocialLinks}>
                                    <a href="https://www.facebook.com/"><FontAwesomeIcon icon={SoIcons.faFacebook} /></a>
                                    <a href="https://twitter.com/"><FontAwesomeIcon icon={SoIcons.faTwitter} /></a>
                                    <a href="https://www.instagram.com/"><FontAwesomeIcon icon={SoIcons.faInstagram} /></a>
                                    <a href="https://pk.linkedin.com/"><FontAwesomeIcon icon={SoIcons.faLinkedinIn} /></a>
                                </div>
                            </div>
                        </div>
                    </div>
                </footer>
            </div>
        </>
    )
}

export default Footer