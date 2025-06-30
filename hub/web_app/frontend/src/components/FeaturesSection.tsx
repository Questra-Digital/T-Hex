"use client";
import styles from "@/styles/components/FeaturesSection.module.scss";
import Image from "next/image";
import { useEffect, useState } from "react";

// Constants
const SELF_HOSTED_FEATURES = [
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

const SAAS_FEATURES = [
  {
    featureName: "Cloud Testing",
    description: "Scalable Cloud Infrastructure",
  },
  { featureName: "Analytics", description: "Advanced Test Analytics" },
  { featureName: "Integration", description: "CI/CD Pipeline Integration" },
  { featureName: "Collaboration", description: "Team Collaboration Tools" },
  { featureName: "Reporting", description: "Comprehensive Test Reports" },
];

const CATEGORIES = [
  { name: "Self-Hosted", description: "Explore our self-hosted features" },
  { name: "SaaS", description: "Explore our SaaS features" },
];

const BREAKPOINT_DESKTOP = 576;
const TRANSITION_DELAY = 300;
const CASCADE_DELAY = 100;
const FEATURE_ICON_SIZE = 34;

export default function FeaturesSection() {
  const [activeCategory, setActiveCategory] = useState(0);
  const [isTransitioning, setIsTransitioning] = useState(false);
  const [isDesktop, setIsDesktop] = useState(false);

  // Get current features based on active category
  const currentFeatures = activeCategory === 0 ? SELF_HOSTED_FEATURES : SAAS_FEATURES;

  // Handle category click
  const handleCategoryClick = (category: number) => {
    if (category !== activeCategory) {
      setIsTransitioning(true);

      setTimeout(() => {
        setActiveCategory(category);
        setIsTransitioning(false);
      }, TRANSITION_DELAY);
    }
  };

  // Handle screen size changes
  useEffect(() => {
    const handleResize = () => {
      setIsDesktop(window.innerWidth >= BREAKPOINT_DESKTOP);
    };

    handleResize();
    window.addEventListener("resize", handleResize);

    return () => window.removeEventListener("resize", handleResize);
  }, []);

  const renderCategoryButtons = () => (
    <div className={styles.categoriesContainer}>
      {CATEGORIES.map((category, index) => (
        <button
          key={category.name}
          className={`${styles.category} ${
            activeCategory === index ? styles.active : ""
          }`}
          onClick={() => handleCategoryClick(index)}
        >
          {category.name} <p>{category.description}</p>
        </button>
      ))}
    </div>
  );

  const renderFeatureCard = (feature: typeof SELF_HOSTED_FEATURES[0], index: number) => (
    <div
      className={`${styles.featureCard} ${
        !isTransitioning ? styles.slideIn : styles.slideOut
      }`}
      key={`${activeCategory}-${index}`}
      style={{
        animationDelay: `${index * CASCADE_DELAY}ms`,
      }}
    >
      {isDesktop ? (
        // Desktop structure: Icon and name in top row, description below
        <>
          <div className={styles.topRow}>
            <Image
              src="/Icons/FeatureIcon.svg"
              width={FEATURE_ICON_SIZE}
              height={FEATURE_ICON_SIZE}
              alt="Feature Icon"
              className={styles.imgIcon}
            />
            <h1 className={styles.featureName}>{feature.featureName}</h1>
          </div>
          <p className={styles.featureDescription}>{feature.description}</p>
        </>
      ) : (
        // Mobile structure: Icon and text side by side
        <>
          <Image
            src="/Icons/FeatureIcon.svg"
            width={FEATURE_ICON_SIZE}
            height={FEATURE_ICON_SIZE}
            alt="Feature Icon"
            className={styles.imgIcon}
          />
          <div className={styles.textContainer}>
            <h1 className={styles.featureName}>{feature.featureName}</h1>
            <p className={styles.featureDescription}>{feature.description}</p>
          </div>
        </>
      )}
    </div>
  );

  const renderFeaturesContainer = () => (
    <div className={styles.featuresContainer}>
      {currentFeatures.map(renderFeatureCard)}
    </div>
  );

  return (
    <section className={`${styles.features}`}>
      <h1 className={styles.featuresHeading}>
        Test Your <span>Websites</span> and <span>Mobile Apps</span>
      </h1>
      <div className={styles.mainContainer}>
        {renderCategoryButtons()}
        <span className={styles.border}></span>
        {renderFeaturesContainer()}
      </div>
    </section>
  );
}