import styles from "@/styles/components/HeroSection.module.scss";
import { outfit } from "@/font/fonts";
import Image from "next/image";

const logos = [
  { src: "/Icons/Microsoft.svg", alt: "Microsoft", width: 100, height: 100 },
  { src: "/Icons/Airbus.svg", alt: "Airbus", width: 100, height: 100 },
  { src: "/Icons/Nvidia.svg", alt: "Nvidia", width: 100, height: 100 },
  { src: "/Icons/Telstra.svg", alt: "Telstra", width: 100, height: 100 },
  { src: "/Icons/Rubrik.svg", alt: "Rubrik", width: 100, height: 100 },
];

export default function HeroSection() {
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
      <div className={styles.buttonsGroup}>
        <button className={styles.googleButton}>
          <Image
            className={styles.googleIcon}
            src="/Icons/Google Icon.svg"
            alt="Google Icon"
            width={55}
            height={55}
          />
          Start Free with Google
        </button>
        <button>Start Free with Email</button>
      </div>
      <div className={styles.trustedGlobally}>
        <p>Trusted by users globally</p>
        <div className={styles.usersIcons}>
          {logos.map((logo) => {
            const classKey = logo.alt.toLowerCase();
            return (
              <div key={classKey} className={styles.iconCell}>
                <Image
                  src={logo.src}
                  alt={logo.alt}
                  width={120}
                  height={80}
                  className={styles[classKey]}
                />
              </div>
            );
          })}
        </div>
      </div>

      <div className={styles.imageContainer}>
        <span className={styles.backgroundRectangle}></span>
        <Image
          src="/Icons/Thex_side_image.png"
          alt="Thex Dashboard Image"
          width={865.5}
          height={577}
          className={styles.dashboard}
        />
        <Image
          src="/Icons/Play circle.svg"
          alt="Play Circle Icon"
          width={48}
          height={48}
          className={styles.playCircle}
        />
      </div>
    </section>
  );
}
