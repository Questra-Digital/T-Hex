import Header from "@/components/Header/Header";
import styles from "@/styles/pages/LandingPage.module.scss";
import { outfit } from "@/fonts/fonts";

export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <div className={`${styles.landingPage} ${outfit.variable}`}>
      <header>
        <Header />
      </header>
      {children}
    </div>
  );
}