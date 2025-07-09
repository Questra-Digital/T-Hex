import styles from "@/styles/components/PricingSection.module.scss";
import Image from "next/image";
import { dm_sans } from "../fonts/fonts";

// Constants
const FREE_TRIAL_FEATURES = [
  "1 Connected Agent",
  "Control Plane",
  "250 Tests Execution",
  "3 Users",
  "As much time as you need",
];

const COMMERCIAL_FEATURES = [
  "10 Connected Agents",
  "Advanced Control Plane",
  "Unlimited Test Execution",
];

const COMMERCIAL_ADDONS = [
  "Additonal Connect Agent ($500/month)",
  "Additional 5 Users($100/month)",
  "SS0 + Resource Groups($500/month)",
  "10+ Read-Only Users($50/month)",
];

const ICON_SIZE = 72;
const CHECK_ICON_SIZE = 26;

export default function PricingSection() {
  const renderFeatureList = (features: string[]) => (
    <ul className={styles.contentList}>
      {features.map((feature, index) => (
        <li key={index}>
          <Image
            src="/Icons/Check Circle.svg"
            width={CHECK_ICON_SIZE}
            height={CHECK_ICON_SIZE}
            alt="Check Circle"
          />
          <p>{feature}</p>
        </li>
      ))}
    </ul>
  );

  const renderFreeCard = () => (
    <div className={`${styles.card} ${styles.freeCard}`}>
      <div className={styles.cardInnerContainer}>
        <div className={styles.cardTitle}>
          <Image
            src="/Icons/FreeIcon.svg"
            width={ICON_SIZE}
            height={ICON_SIZE}
            alt="Free Icon"
            className={styles.icon}
          />
          <p className={styles.title}>Free Trial</p>
        </div>

        <div className={styles.cardContent}>
          <div className={styles.contentListContainer}>
            <p className={styles.contentHeading}>What's Included</p>
            {renderFeatureList(FREE_TRIAL_FEATURES)}
          </div>
        </div>

        <div className={styles.cardMessage}>
          <p>
            <span>Talk with us </span>
            to request additional capacity at any time
          </p>
        </div>

        <button className={styles.cardButton}>Get Started</button>
      </div>
    </div>
  );

  const renderCommercialCard = () => (
    <div className={`${styles.card} ${styles.commercialCard}`}>
      <div className={styles.cardInnerContainer}>
        <div className={styles.cardTitle}>
          <Image
            src="/Icons/Commercial Icon.svg"
            width={ICON_SIZE}
            height={ICON_SIZE}
            alt="Commercial Icon"
            className={styles.icon}
          />
          <p>Commercial Liscence</p>
        </div>

        <p className={styles.price}>
          <span>$700</span>/Month
        </p>

        <div className={styles.cardContent}>
          <div className={styles.contentListContainer}>
            <p className={styles.contentHeading}>What's Included</p>
            {renderFeatureList(COMMERCIAL_FEATURES)}
          </div>

          <div className={styles.contentListContainer}>
            <p className={styles.contentHeading}>Add-Ons:</p>
            {renderFeatureList(COMMERCIAL_ADDONS)}
          </div>
        </div>

        <div className={styles.cardMessage}>
          <p>
            <span>Contact us </span>
            for start-up and non-profit pricing
          </p>
        </div>

        <button className={styles.cardButton}>Contact Us</button>
      </div>
    </div>
  );

  return (
    <section className={`${styles.pricingSection} ${dm_sans.variable}`}>
      <h1 className={styles.pricingHeading}>
        T-Hex <span>Pricing</span>
      </h1>

      <div className={styles.cardsContainer}>
        {renderFreeCard()}
        {renderCommercialCard()}
      </div>
    </section>
  );
}
