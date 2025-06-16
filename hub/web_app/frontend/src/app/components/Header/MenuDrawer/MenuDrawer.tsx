import styles from "./MenuDrawer.module.scss";
import Image from "next/image";
import { outfit } from "@/font/fonts";

interface Link {
  name: string;
  href: string;
}

interface Links {
  links: Link[];
}

export default function Drawer({ links }: Links) {
  return (
    <div className={styles.drawerContainer}> 
      {links.map((link) => (
        <div key={link.name} className={styles.linkCard}> 
          <p className={`${styles.linkName} ${outfit.variable}`}>{link.name}</p> 
          <Image 
            height={34}
            width={34}
            src="/Icons/arrow_down.svg"
            alt="Arrow Down Icon"
            className={styles.arrowIcon}
          />
        </div>
      ))}
    </div>
  );
}
