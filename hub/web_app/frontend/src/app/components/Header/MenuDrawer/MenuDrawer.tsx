import styles from "./MenuDrawer.module.scss";
import Image from "next/image";
import { outfit } from "@/font/fonts";
import { useEffect, useState } from "react";

interface Link {
  name: string;
  href: string;
}

interface MenuDrawerProps {
  links: Link[];
  isDrawerOpen: boolean;
}

export default function Drawer({
  links,
  isDrawerOpen = false,
}: MenuDrawerProps) {
  const [shouldRender, setShouldRender] = useState(isDrawerOpen);

  useEffect(() => {
    if (!isDrawerOpen) {
      const timer = setTimeout(() => setShouldRender(false), 300); //wait for 3 milliseconds so that closing class takes effect
      return () => clearTimeout(timer); //CLeanup timer
    } else {
      setShouldRender(true);
    }
  }, [isDrawerOpen]);

  return (
    <>
      {shouldRender && (
        <div
          className={`${styles.drawerContainer} ${
            !isDrawerOpen ? styles.closing : ""
          }`}
        >
          {links.map((link) => (
            <div key={link.name} className={styles.linkCard}>
              <p className={`${styles.linkName} ${outfit.variable}`}>
                {link.name}
              </p>
              <Image
                height={34}
                width={34}
                src="/Icons/arrow_down.svg"
                alt="Arrow Down Icon"
                className={styles.arrowIcon}
              />
            </div>
          ))}

          <div className={`${styles.linkCard}`}>
            <button className={outfit.variable}>Login</button>
            <button className={`${styles.getStarted} ${outfit.variable}`}>Get Started Free</button>
          </div>
        </div>
      )}
    </>
  );
}
