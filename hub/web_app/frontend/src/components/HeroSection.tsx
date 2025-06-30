import styles from "@/styles/components/HeroSection.module.scss";
import { outfit } from "../fonts/fonts";
import Image from "next/image";

// Constants
const COMPANY_LOGOS = [
  { src: "/Icons/Microsoft.svg", alt: "Microsoft", width: 100, height: 100 },
  { src: "/Icons/Airbus.svg", alt: "Airbus", width: 100, height: 100 },
  { src: "/Icons/Nvidia.svg", alt: "Nvidia", width: 100, height: 100 },
  { src: "/Icons/Telstra.svg", alt: "Telstra", width: 100, height: 100 },
  { src: "/Icons/Rubrik.svg", alt: "Rubrik", width: 100, height: 100 },
];

const GOOGLE_ICON_SIZE = 55;
const LOGO_DISPLAY_SIZE = { width: 120, height: 80 };
const DASHBOARD_SIZE = { width: 865.5, height: 577 };
const PLAY_CIRCLE_SIZE = 48;

export default function HeroSection() {
  const renderCompanyLogos = () => (
    <div className={styles.usersIcons}>
      {COMPANY_LOGOS.map((logo) => {
        const classKey = logo.alt.toLowerCase();
        return (
          <div key={classKey} className={styles.iconCell}>
            <Image
              src={logo.src}
              alt={logo.alt}
              width={LOGO_DISPLAY_SIZE.width}
              height={LOGO_DISPLAY_SIZE.height}
              className={styles[classKey]}
            />
          </div>
        );
      })}
    </div>
  );

  const renderButtons = () => (
    <div className={styles.buttonsGroup}>
      <button className={styles.googleButton}>
        <Image
          className={styles.googleIcon}
          src="/Icons/Google Icon.svg"
          alt="Google Icon"
          width={GOOGLE_ICON_SIZE}
          height={GOOGLE_ICON_SIZE}
        />
        Start Free with Google
      </button>
      <button>Start Free with Email</button>
    </div>
  );

  const renderTrustedSection = () => (
    <div className={styles.trustedGlobally}>
      <p>Trusted by users globally</p>
      {renderCompanyLogos()}
    </div>
  );

  const renderImageSection = () => (
    <div className={styles.imageContainer}>
      <span className={styles.backgroundRectangle}></span>
      <Image
        src="/Icons/Thex_side_image.png"
        alt="Thex Dashboard Image"
        width={DASHBOARD_SIZE.width}
        height={DASHBOARD_SIZE.height}
        className={styles.dashboard}
      />
      <Image
        src="/Icons/Play circle.svg"
        alt="Play Circle Icon"
        width={PLAY_CIRCLE_SIZE}
        height={PLAY_CIRCLE_SIZE}
        className={styles.playCircle}
      />
    </div>
  );

  return (
    <section className={`${styles.heroSection} ${outfit.variable}`}>
      <h1 className={styles.heroHeading}>
        Power Your Software Testing with<span> AI and Cloud</span>
      </h1>
      <span className={styles.breakLine}></span>
      <p className={styles.bodyText}>
        Test Intelligently and ship faster. Deliver unparalled digital
        experiences for real world enterprises
      </p>
      {renderButtons()}
      {renderTrustedSection()}
      {renderImageSection()}
    </section>
  );
}
