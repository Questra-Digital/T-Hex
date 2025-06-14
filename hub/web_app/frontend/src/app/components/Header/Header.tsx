"use client";
import styles from "./Header.module.scss";
import Link from "next/link";
import { montserrat, outfit } from "@/font/fonts";
import Image from "next/image";
import { useState, useEffect } from "react";

const links = [
  { name: "Home", href: "/landingPage" },
  { name: "Solutions", href: "/landingPage" },
  { name: "Services", href: "/landingPage" },
  { name: "Pricing", href: "/landingPage" },
  { name: "Technology Stack", href: "/landingPage" },
  { name: "Contact", href: "/landingPage" },
];

export default function Header() {
  const [isSmall, setIsSmall] = useState(false);
  const [isMedium, setIsMedium] = useState(false);
  const [isLarge, setIsLarge] = useState(false);
  const [screenSize, setScreenSize] = useState(false);

  return (
    <header
      className={`${styles.header} ${montserrat.variable} ${outfit.variable}`}
    >
      {/* Menu Icon */}
      <Image
        src="/Icons/Menu.svg"
        alt="Menu Icon"
        className={styles.menuIcon}
        width={58}
        height={58}
      />

      {/* T-Hex Logo */}
      <h1>T-Hex</h1>
      
      {/* Nav Bar */}
      <nav>
        <ul>
          {/*map each link to Next's Link*/}
          {links.map((link) => {
            return (
              <li key={link.name}>
                <Link href={link.href}>
                  {link.name}
                  <Image
                    height={32}
                    width={32}
                    src="/Icons/arrow_down.svg"
                    alt="Arrow Down Icon"
                    className={styles.arrowIcon}
                  />
                </Link>
              </li>
            );
          })}
        </ul>
      </nav>

      {/* Button Groups */}
      <div className={styles.buttonsGroup}>
        <button>Login</button>
        <button className={styles.getStarted}>Get Started Free</button>
      </div>
    </header>
  );
}
