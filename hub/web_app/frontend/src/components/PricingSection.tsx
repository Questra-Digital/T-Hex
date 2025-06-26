import styles from "@/styles/components/PricingSection.module.scss";
import Image from "next/image";
import { dm_sans, outfit } from "@/font/fonts";

export default function PricingSection() {
  return (
    <section className={`${styles.pricingSection} ${dm_sans.variable}`}>
      <h1 className={outfit.variable}>
        T-Hex <span>Pricing</span>
      </h1>

      <div className={styles.cardsContainer}>
        {/* Free Trial Card */}
        <div className={`${styles.card} ${styles.freeCard}`}>
          <div className={styles.cardInnerContainer}>
            <div className={styles.cardTitle}>
              <Image
                src="/Icons/FreeIcon.svg"
                width={72}
                height={72}
                alt="Free Icon"
                className={styles.icon}
              />
              <p className={styles.title}>Free Trial</p>
            </div>

            <div className={styles.cardContent}>
              <div className={styles.contentListContainer}>
                <p className={styles.contentHeading}>What's Included</p>
                <ul className={styles.contentList}>
                  <li>
                    <Image
                      src="/Icons/Check Circle.svg"
                      width={26}
                      height={26}
                      alt="Check Circle"
                    />
                    <p>1 Connected Agent</p>
                  </li>
                  <li>
                    <Image
                      src="/Icons/Check Circle.svg"
                      width={26}
                      height={26}
                      alt="Check Circle"
                    />
                    <p>Control Plane</p>
                  </li>
                  <li>
                    <Image
                      src="/Icons/Check Circle.svg"
                      width={26}
                      height={26}
                      alt="Check Circle"
                    />
                    <p>250 Tests Execution</p>
                  </li>
                  <li>
                    <Image
                      src="/Icons/Check Circle.svg"
                      width={26}
                      height={26}
                      alt="Check Circle"
                    />
                    <p>3 Users</p>
                  </li>
                  <li>
                    <Image
                      src="/Icons/Check Circle.svg"
                      width={26}
                      height={26}
                      alt="Check Circle"
                    />
                    <p>As much time as you need</p>
                  </li>
                </ul>
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

        {/* Commercial Card */}
        <div className={`${styles.card} ${styles.commercialCard}`}>
          <div className={styles.cardInnerContainer}>
            <div className={styles.cardTitle}>
              <Image
                src="/Icons/Commercial Icon.svg"
                width={72}
                height={72}
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
                <ul className={styles.contentList}>
                  <li>
                    <Image
                      src="/Icons/Check Circle.svg"
                      width={26}
                      height={26}
                      alt="Check Circle"
                    />
                    <p>10 Connected Agents</p>
                  </li>
                  <li>
                    <Image
                      src="/Icons/Check Circle.svg"
                      width={26}
                      height={26}
                      alt="Check Circle"
                    />
                    <p>Advanced Control Plane</p>
                  </li>
                  <li>
                    <Image
                      src="/Icons/Check Circle.svg"
                      width={26}
                      height={26}
                      alt="Check Circle"
                    />
                    <p>Unlimited Test Execution</p>
                  </li>
                </ul>
              </div>

              <div className={styles.contentListContainer}>
                <p className={styles.contentHeading}>Add-Ons:</p>
                <ul className={styles.contentList}>
                  <li>
                    <Image
                      src="/Icons/Check Circle.svg"
                      width={26}
                      height={26}
                      alt="Check Circle"
                    />
                    <p>Additonal Connect Agent ($500/month)</p>
                  </li>
                  <li>
                    <Image
                      src="/Icons/Check Circle.svg"
                      width={26}
                      height={26}
                      alt="Check Circle"
                    />
                    <p>Additional 5 Users($100/month)</p>
                  </li>
                  <li>
                    <Image
                      src="/Icons/Check Circle.svg"
                      width={26}
                      height={26}
                      alt="Check Circle"
                    />
                    <p>SS0 + Resource Groups($500/month)</p>
                  </li>
                  <li>
                    <Image
                      src="/Icons/Check Circle.svg"
                      width={26}
                      height={26}
                      alt="Check Circle"
                    />
                    <p>10+ Read-Only Users($50/month)</p>
                  </li>
                </ul>
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
      </div>
    </section>
  );
}
