import styles from '@/styles/components/Footer.module.scss';
import Link from 'next/link';
import { FaFacebook, FaTwitter, FaInstagram, FaLinkedin } from 'react-icons/fa';

export default function Footer() {
    return (
        <footer className={styles.footer}>
            <div className={styles.footerContent}>
                {/* Company Information */}
                <div className={styles.companyInfo}>
                    <div className={styles.logo}>
                        <h3>T-Hex</h3>
                        <p>Empowering businesses with scalable testing platform</p>
                    </div>
                    <div className={styles.socialLinks}>
                        <h4>Follow Us</h4>
                        <div className={styles.socialIcons}>
                            <Link href="https://facebook.com" aria-label="Facebook" target="_blank" rel="noopener noreferrer">
                                <FaFacebook />
                            </Link>
                            <Link href="https://twitter.com" aria-label="Twitter" target="_blank" rel="noopener noreferrer">
                                <FaTwitter />
                            </Link>
                            <Link href="https://instagram.com" aria-label="Instagram" target="_blank" rel="noopener noreferrer">
                                <FaInstagram />
                            </Link>
                            <Link href="https://linkedin.com" aria-label="LinkedIn" target="_blank" rel="noopener noreferrer">
                                <FaLinkedin />
                            </Link>
                        </div>
                    </div>
                </div>

                {/* Contact Information */}
                <div className={styles.contactInfo}>
                    <h4>Contact Us</h4>
                    <div className={styles.contact}>
                        <h5>Address</h5>
                        <p>852-B Milaad St, Block B Faisal Town, Lahore, Pakistan</p>
                    </div>
                    <div className={styles.contact}>
                        <h5>Phone</h5>
                        <a href="tel:+9242111128128">(042) 111 128 128</a>
                    </div>
                    <div className={styles.contact}>
                        <h5>Email</h5>
                        <a href="mailto:info@t-hex.com">info@t-hex.com</a>
                    </div>
                </div>

                {/* Quick Links */}
                <div className={styles.quickLinks}>
                    <h4>Quick Links</h4>
                    <nav>
                        <Link href="/landingPage">Home</Link>
                        <Link href="/landingPage">About Us</Link>
                        <Link href="/landingPage">Our Services</Link>
                        <Link href="/landingPage">Technology Stack</Link>
                        <Link href="/landingPage">Portfolio</Link>
                        <Link href="/landingPage">Careers</Link>
                        <Link href="/landingPage">Blog</Link>
                        <Link href="/landingPage">Contact</Link>
                    </nav>
                </div>
            </div>

            {/* Bottom Section */}
            <div className={styles.footerBottom}>
                <div className={styles.legalLinks}>
                    <Link href="/landingPage">Privacy Policy</Link>
                    <Link href="/landingPage">Terms of Service</Link>
                    <Link href="/landingPage">Cookie Policy</Link>
                    <Link href="/landingPage">Sitemap</Link>
                </div>
                <div className={styles.copyright}>
                    <p>&copy; {new Date().getFullYear()} T-Hex. All rights reserved.</p>
                </div>
            </div>
        </footer>
    );
}
