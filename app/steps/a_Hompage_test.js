// Import CodeceptJS and the necessary modules
const { I } = require('codeceptjs');

Feature('search');

Scenario('test something on the homepage', async ({ I }) => {
  // Navigate to the page containing the CloneRepository component
  I.amOnPage('http://localhost:3000');
  I.wait(3);
  I.see('WebTestHub');
  I.click('Get Started');
  I.wait(2);
});