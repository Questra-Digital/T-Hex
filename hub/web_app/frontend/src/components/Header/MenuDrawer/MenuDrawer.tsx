import styles from "@/styles/components/Header/MenuDrawer.module.scss";
import Image from "next/image";
import { outfit } from "../../../fonts/fonts";
import { useEffect, useState } from "react";

interface Link {
  name: string;
  href: string;
}

interface MenuDrawerProps {
  links: Link[];
  isDrawerOpen: boolean;
}

// Constants
const ANIMATION_DURATION = 300; // milliseconds
const ARROW_ICON_SIZE = 34;

export default function MenuDrawer({
  links,
  isDrawerOpen = false,
}: MenuDrawerProps) {
  const [shouldRender, setShouldRender] = useState(isDrawerOpen);

  useEffect(() => {
    if (!isDrawerOpen) {
      const timer = setTimeout(() => setShouldRender(false), ANIMATION_DURATION);
      return () => clearTimeout(timer);
    } else {
      setShouldRender(true);
    }
  }, [isDrawerOpen]);

  const renderNavigationLinks = () => (
    <>
      {links.map((link) => (
        <div key={link.name} className={styles.linkCard}>
          <p className={`${styles.linkName} ${outfit.variable}`}>
            {link.name}
          </p>
          <Image
            height={ARROW_ICON_SIZE}
            width={ARROW_ICON_SIZE}
            src="/Icons/arrow_down.svg"
            alt="Arrow Down Icon"
            className={styles.arrowIcon}
          />
        </div>
      ))}
    </>
  );

  const renderButtons = () => (
    <div className={styles.linkCard}>
      <button className={outfit.variable}>Login</button>
      <button className={`${styles.getStarted}`}>
        Get Started Free
      </button>
    </div>
  );

  if (!shouldRender) {
    return null;
  }

  return (
    <div
      className={`${styles.drawerContainer}  ${outfit.variable} ${
        !isDrawerOpen ? styles.closing : ""
      }`}
    >
      {renderNavigationLinks()}
      {renderButtons()}
    </div>
  );
}
