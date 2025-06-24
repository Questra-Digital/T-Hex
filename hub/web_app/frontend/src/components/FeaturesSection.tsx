"use client";
import styles from "@/styles/components/FeaturesSection.module.scss";
import Image from "next/image";
import { outfit } from "@/font/fonts";
import { useEffect, useState } from "react";

const selfHostedFeatures = [
  { featureName: "Live", description: "Manual Cross Browser Testing" },
  { featureName: "Automate", description: "Browser Automation Grid" },
  { featureName: "Percy", description: "Automated Visual Testing" },
  { featureName: "App Live", description: "Manual Real Device Testing" },
  { featureName: "App Automate", description: "Real Device Automation Cloud" },
  {
    featureName: "Test Management",
    description: "Unity & Track All Test Cases",
  },
];

const saasFeatures = [
  {
    featureName: "Cloud Testing",
    description: "Scalable Cloud Infrastructure",
  },
  { featureName: "Analytics", description: "Advanced Test Analytics" },
  { featureName: "Integration", description: "CI/CD Pipeline Integration" },
  { featureName: "Collaboration", description: "Team Collaboration Tools" },
  { featureName: "Reporting", description: "Comprehensive Test Reports" },
];

export default function FeaturesSection() {
  const [activeCategory, setActiveCategory] = useState(0);
  const [isTransitioning, setIsTransitioning] = useState(false);
  const [isDesktop, setIsDesktop] = useState(false);

  // Get current features based on active category
  const currentFeatures =
    activeCategory === 0 ? selfHostedFeatures : saasFeatures;

  // Handle category click
  const handleCategoryClick = (category: number) => {
    if (category !== activeCategory) {
      setIsTransitioning(true);

      // Start transition after a brief delay
      setTimeout(() => {
        setActiveCategory(category);
        setIsTransitioning(false);
      }, 300);
    }
  };

  //useEffect to handle screen size
  useEffect(() => {
    const handleResize = () => {
      setIsDesktop(window.innerWidth >= 576);
    };

    handleResize();

    //Add Event Listener
    window.addEventListener("resize", handleResize);

    // Cleanup
    return () => window.removeEventListener("resize", handleResize);
  });

  return (
    <div className={`${styles.features} ${outfit.variable}`}>
      <h1 className={styles.featuresHeading}>
        Test Your <span>Websites</span> and <span>Mobile Apps</span>
      </h1>
      <div className={styles.mainContainer}>
        <div className={styles.categoriesContainer}>
          <button
            className={`${styles.category} ${
              activeCategory === 0 ? styles.active : ""
            }`}
            onClick={() => handleCategoryClick(0)}
          >
            Self-Hosted <p>Explore our self-hosted features</p>
          </button>
          <button
            className={`${styles.category} ${
              activeCategory === 1 ? styles.active : ""
            }`}
            onClick={() => handleCategoryClick(1)}
          >
            SaaS<p>Explore our SaaS features</p>
          </button>
        </div>
        <span className={styles.border}></span>
        <div className={styles.featuresContainer}>
          {currentFeatures.map((feature, index) => (
            <div
              className={`${styles.featureCard} ${
                !isTransitioning ? styles.slideIn : styles.slideOut
              }`}
              key={`${activeCategory}-${index}`}
              style={{
                animationDelay: `${index * 100}ms`, // Cascading delay
              }}
            >
              {isDesktop ? (
                // Desktop structure: Icon and name in top row, description below
                <>
                  <div className={styles.topRow}>
                    <Image
                      src="/Icons/FeatureIcon.svg"
                      width={34}
                      height={34}
                      alt="Feature Icon"
                      className={styles.imgIcon}
                    />
                    <h1 className={styles.featureName}>
                      {feature.featureName}
                    </h1>
                  </div>
                  <p className={styles.featureDescription}>
                    {feature.description}
                  </p>
                </>
              ) : (
                <>
                  <Image
                    src="/Icons/FeatureIcon.svg"
                    width={34}
                    height={34}
                    alt="Feature Icon"
                    className={styles.imgIcon}
                  />
                  <div className={styles.textContainer}>
                    <h1 className={styles.featureName}>
                      {feature.featureName}
                    </h1>
                    <p className={styles.featureDescription}>
                      {feature.description}
                    </p>
                  </div>
                </>
              )}
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
