import React from 'react';
import { Link } from 'react-router-dom'; 
import styles from './SubscriptionPlans.module.css';

function SubscriptionPlans() {
  return (
    <div style={{ maxWidth: '960px', margin: '0 auto' }}>
      <h1 style={{ textAlign: 'center', margin: '20px 0' }}>Choose Your Plan</h1>
      <div className={styles.container}>
        <div className={styles.plan}>
          <div className={styles.title}>Free Plan</div>
          <div className={styles.description}>Access to basic features</div>
          <div className={styles.description}>5 test Cases per day</div>
          <div className={styles.description}>No Parallel Testing</div>
          <div className={styles.price}>$0/month</div>
          <Link to="/payment">
                        <button type="submit">Subscribe</button>
                        </Link>
        </div>
        <div className={styles.plan}>
          <div className={styles.title}>Individual Plan</div>
          <div className={styles.description}>Full access for one user</div>
          <div className={styles.description}>Mutilple Parallel Testing</div>
          <div className={styles.description}>Unlimited Number of Test</div>
          <div className={styles.price}>$20/month</div>
          <Link to="/payment">
                        <button type="submit">Subscribe</button>
                        </Link>
        </div>
        <div className={styles.plan}>
          <div className={styles.title}>Team Plan</div>
          <div className={styles.description}>Full access for the team</div>
          <div className={styles.description}>Mutilple Parallel Testing</div>
          <div className={styles.description}>Unlimited Number of Test</div>
          <div className={styles.price}>$50/month</div>
          <Link to="/payment">
                        <button type="submit">Subscribe</button>
                        </Link>
        </div>
      </div>
    </div>
  );
}

export default SubscriptionPlans;
