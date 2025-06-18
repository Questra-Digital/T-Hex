import styles from "./HeroSection.module.scss";
import { outfit } from "@/font/fonts";
import Image from "next/image";

const logos = [
  { src: "/Icons/Microsoft.svg", alt: "Microsoft", width: 174, height: 116 },
  { src: "/Icons/Vimeo.svg", alt: "Vimeo", width: 103, height: 62 },
  { src: "/Icons/Nvidia.svg", alt: "Nvidia", width: 156, height: 156 },
  { src: "/Icons/Telstra.svg", alt: "Telstra", width: 134, height: 101 },
  { src: "/Icons/Rubrik.svg", alt: "Rubrik", width: 108, height: 31 },
];

export default function HeroSection() {
  return (
    <section className={`${styles.heroSection} ${outfit.variable}`}>
      <div className={styles.leftSection}>
        <h1>
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
              const classKey = logo.alt.toLowerCase(); // e.g. "Microsoft" → "microsoft"
              return (
                <div key={classKey} className={styles.iconCell}>
                  <Image
                    src={logo.src}
                    alt={logo.alt}
                    width={logo.width}
                    height={logo.height}
                    className={styles[classKey]}
                  />
                </div>
              );
            })}
          </div>
        </div>
      </div>
      <div className={styles.rightSection}>
      <div className={styles.imageContainer}>
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
      </div>
    </section>
  );
}
