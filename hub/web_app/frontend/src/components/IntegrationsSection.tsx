import styles from "@/styles/components/Integrations.module.scss";
import Image from "next/image";

const integrations = [
  {
    integrationName: "Selenium Test",
    integrationIcon: "/Icons/Integrations/Selenium.svg",
    description: "How to Create a Selenium Test",
  },
  {
    integrationName: "Gradle Test",
    integrationIcon: "/Icons/Integrations/Gradle.svg",
    description: "How to Create Gradle Test",
  },
  {
    integrationName: "JMeter Test",
    integrationIcon: "/Icons/Integrations/jmeter.png",
    description: "How to Create JMeter Test",
  },
  {
    integrationName: "KubePug Test",
    integrationIcon: "/Icons/Integrations/KubePug.png",
    description: "How to Create KubePug Test",
  },
  {
    integrationName: "Maven Test",
    integrationIcon: "/Icons/Integrations/Maven.svg",
    description: "How to Create Maven Test",
  },
  {
    integrationName: "Cypress Test",
    integrationIcon: "/Icons/Integrations/Cypress.svg",
    description: "How to Create a Cypress Test",
  },
];

// Helper function to convert integration name to CSS class name
const getIntegrationClass = (integrationName: string) => {
  return integrationName
    .toLowerCase()
    .replace(/\s+/g, "-") // Replace spaces with hyphens
    .replace(/[^a-z0-9-]/g, ""); // Remove special characters except hyphens
};

export default function IntegrationsSection() {
  return (
    <section className={styles.integration}>
      <h1 className={styles.integrationTitle}>
        <span>Integrations</span>
      </h1>
      <p className={styles.integrationSubtitle}>
        Explore our supported integrations and learn how to connect your
        favorite tools for seamless testing and automation.
      </p>
      <div className={styles.integrationsContainer}>
        {integrations.map((integration, index) => {
          const integrationClass = getIntegrationClass(
            integration.integrationName
          );
          return (
            <div key={index} className={`${styles.integrationCard} ${styles[integrationClass]}`}>
              <div className={styles.imageContainer}>
                <Image
                  src={integration.integrationIcon}
                  width={200}
                  height={200}
                  alt={`${integration.integrationName} icon`}
                  priority={index < 3}
                />
              </div>
              <p className={styles.integrationName}>
                {integration.integrationName}
              </p>
              <p className={styles.integrationDescription}>
                {integration.description}
              </p>
              <button>View Documentation</button>
            </div>
          );
        })}
      </div>
    </section>
  );
}
