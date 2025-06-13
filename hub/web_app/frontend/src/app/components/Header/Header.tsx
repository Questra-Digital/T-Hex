import styles from "./Header.module.scss";
import Link from "next/link";

const links =[
{name:'Home', href:'/landingPage'},
{name:'Solutions', href:'/landingPage'},
{name:'Services', href:'/landingPage'},
{name:'Pricing', href:'/landingPage'},
{name:'Technology Stack', href:'/landingPage'},
{name:'Contact', href:'/landingPage'}
];

export default function Header() {
  return (
    <header className={styles.header}>
      {/* T-Hex Logo */}
      <h1>T-Hex</h1>
      {/* Nav Bar */}
      <nav>
        <ul>
          {/*map each link to Next's Link*/}
          {links.map((link) => {
            return <li key={link.name} >
              <Link href={link.href}>
              {link.name}
              </Link>
            </li>
          })}
        </ul>
      </nav>
      <div className={styles.buttonsGroup}>
      <button >Login</button>
      <button className={styles.getStarted}>Get Started Free</button>
      </div>
      
    </header>
  );
}
