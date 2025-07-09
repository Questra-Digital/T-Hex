"use client";
import styles from "@/styles/components/Header/Header.module.scss";
import Link from "next/link";
import { montserrat, outfit } from "../../fonts/fonts";
import Image from "next/image";
import { useState, useEffect } from "react";
import MenuDrawer from "./MenuDrawer/MenuDrawer";

// Constants
const NAVIGATION_LINKS = [
  { name: "Home", href: "/landingPage" },
  { name: "Solutions", href: "/landingPage" },
  { name: "Services", href: "/landingPage" },
  { name: "Pricing", href: "/landingPage" },
  { name: "Technology Stack", href: "/landingPage" },
  { name: "Contact", href: "/landingPage" },
];

const BREAKPOINT_MOBILE = 768;
const ICON_SIZE = 42;
const ARROW_ICON_SIZE = 32;

export default function Header() {
  const [isDrawerOpen, setIsDrawerOpen] = useState(false);

  useEffect(() => {
    const handleScreenResize = () => {
      if (window.innerWidth >= BREAKPOINT_MOBILE) {
        setIsDrawerOpen(false);
      }
    };

    window.addEventListener("resize", handleScreenResize);
    handleScreenResize(); // Check initial screen size

    return () => {
      window.removeEventListener("resize", handleScreenResize);
    };
  }, []);

  const handleMenuToggle = () => {
    setIsDrawerOpen(!isDrawerOpen);
  };

  const renderNavigationLinks = () => (
    <ul className={styles.list}>
      {NAVIGATION_LINKS.map((link) => (
        <li key={link.name} className={styles.listItem}>
          <Link href={link.href}>
            {link.name}
            <Image
              height={ARROW_ICON_SIZE}
              width={ARROW_ICON_SIZE}
              src="/Icons/arrow_down.svg"
              alt="Arrow Down Icon"
              className={styles.arrowIcon}
            />
          </Link>
        </li>
      ))}
    </ul>
  );

  const renderButtons = () => (
    <div className={styles.buttonsGroup}>
      <button className={styles.button}>Login</button>
      <button className={`${styles.getStarted} ${styles.button}`}>
        Get Started Free
      </button>
    </div>
  );

  return (
    <>
      <header
        className={`${styles.header} ${montserrat.variable} ${outfit.variable}`}
      >
        {/* Menu Icon */}
        <Image
          src={isDrawerOpen ? "/Icons/close.svg" : "/Icons/Menu.svg"}
          alt={isDrawerOpen ? "Close menu" : "Open menu"}
          className={styles.menuIcon}
          width={ICON_SIZE}
          height={ICON_SIZE}
          onClick={handleMenuToggle}
        />

        {/* T-Hex Logo */}
        <h1 className={styles.logo}>T-Hex</h1>

        {/* Navigation Bar */}
        <nav className={styles.navbar}>
          {renderNavigationLinks()}
        </nav>

        {/* Button Groups */}
        {renderButtons()}
      </header>
      
      {/* MenuDrawer for smaller screens */}
      <MenuDrawer links={NAVIGATION_LINKS} isDrawerOpen={isDrawerOpen} />
    </>
  );
}
