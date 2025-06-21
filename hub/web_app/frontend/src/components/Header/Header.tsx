"use client";
import styles from "@/styles/components/Header/Header.module.scss";
import Link from "next/link";
import { montserrat, outfit } from "../../fonts/fonts";
import Image from "next/image";
import { useState, useEffect } from "react";
import MenuDrawer from "./MenuDrawer/MenuDrawer";

const links = [
  { name: "Home", href: "/landingPage" },
  { name: "Solutions", href: "/landingPage" },
  { name: "Services", href: "/landingPage" },
  { name: "Pricing", href: "/landingPage" },
  { name: "Technology Stack", href: "/landingPage" },
  { name: "Contact", href: "/landingPage" },
];

export default function Header() {
  const [isDrawerOpen, setIsDrawerOpen] = useState(false);

  useEffect(() => {
    const handleScreenResize = () => {
      if (window.innerWidth >= 768) {
        setIsDrawerOpen(false); //Stop rendering drawer if screen exceeds 768px
      }
    };

    // Add event listener for resize
    window.addEventListener("resize", handleScreenResize);

    // Check initial screen size on mount
    handleScreenResize();

    return () => {
      // Cleanup event listener
      window.removeEventListener("resize", handleScreenResize);
    };
  }, []);

  function onMenuPressed() {
    setIsDrawerOpen(!isDrawerOpen);
  }

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
          width={42}
          height={42}
          onClick={onMenuPressed}
        />

        {/* T-Hex Logo */}
        <h1 className={styles.logo}>T-Hex</h1>

        {/* Nav Bar */}
        <nav className={styles.navbar}>
          <ul className={styles.list}>
            {/*map each link to Next's Link*/}
            {links.map((link) => {
              return (
                <li key={link.name} className={styles.listItem}>
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
          <button className={styles.button}>Login</button>
          <button className={`${styles.getStarted} ${styles.button}`}>Get Started Free</button>
        </div>
      </header>
      
      {/* MenuDrawer for smaller screens */}
      <MenuDrawer links={links} isDrawerOpen={isDrawerOpen} />
    </>
  );
}
