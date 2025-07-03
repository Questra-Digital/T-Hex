import styles from "@/styles/components/WhatItDoSection.module.scss";
import Image from "next/image";
import WorkflowImage from "./Diagrams/Workflow.png";

const featuresList = [
  { feature: "Execute Tests seamlessly with CLI." },
  { feature: "Integrate with Frameworks like Selenium." },
  { feature: "Deploy efficiently in CI/CD pipelines." },
  {
    feature:
      "Runs tests across multiple environments—local,  cloud, and CI/CD pipelines.",
  },
  { feature: "View real-time test results with detailed analytics." },
];
export default function WhatItDo() {
  return (
    <section className={styles.whatItDo}>
      <h1 className={styles.title}>
        Make Test Automation Easy with <span>T-Hex</span>
      </h1>
      <div className={styles.innerContainer}>
        <ul className={styles.list}>
          {featuresList.map((feature, index) => (
            <li className={styles.listItem} key={index}>
              <Image
                src="/Icons/Green Check Circle.svg"
                alt="Check Circle"
                width={32}
                height={32}
                className={styles.checkCircle}
              />
              <p>{feature.feature}</p>
            </li>
          ))}
        </ul>
        <Image
          src="/Diagrams/Workflow.png"
          alt="Workflow Diagram"
          width={521}
          height={609}
          className={styles.diagram}
        />
      </div>
      <button className={styles.learnMoreButton}>Learn More</button>
    </section>
  );
}
