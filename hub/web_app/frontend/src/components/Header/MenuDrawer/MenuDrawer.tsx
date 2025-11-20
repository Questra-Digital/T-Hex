import styles from "@/styles/components/Header/MenuDrawer.module.scss";
import Image from "next/image";
import { outfit } from "../../../fonts/fonts";
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { useAppSelector } from "@/store/hooks";
import { selectUserData } from "@/store/userData";
import { User } from "lucide-react";

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
  const router = useRouter();
  const userData = useAppSelector(selectUserData);
  
  // Check if user is logged in
  const isLoggedIn = userData.user_id && userData.user_id > 0;

  useEffect(() => {
    if (!isDrawerOpen) {
      const timer = setTimeout(() => setShouldRender(false), ANIMATION_DURATION);
      return () => clearTimeout(timer);
    } else {
      setShouldRender(true);
    }
  }, [isDrawerOpen]);

  const handleLoginClick = () => {
    router.push('/login');
  };

  const handleGetStartedClick = () => {
    router.push('/getting_started');
  };

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
      {isLoggedIn ? (
        <div className={styles.userIcon}>
          <User size={24} />
        </div>
      ) : (
        <button className={outfit.variable} onClick={handleLoginClick}>
          Login
        </button>
      )}
      <button className={`${styles.getStarted}`} onClick={handleGetStartedClick}>
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
