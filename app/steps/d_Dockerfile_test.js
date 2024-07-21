// Import CodeceptJS and the necessary modules
const { I } = require('codeceptjs');

Feature('Home');

Scenario('Homepage scroll', async ({ I }) => {
  // Navigate to the homepage
  I.amOnPage('http://localhost:3000');

  // Wait for a certain amount of time (e.g., 2 seconds)
  I.executeScript(() => {
    window.scrollBy(0, 400); // Adjust the value as needed
  });
  I.wait(2);
  I.executeScript(() => {
    window.scrollBy(800, 0); 
  });
  I.wait(2);
  I.see(' WebTestHub');

});