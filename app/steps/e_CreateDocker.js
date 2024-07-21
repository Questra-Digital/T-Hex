// TODO: test it again
// Import CodeceptJS and the necessary modules
const { I } = require('codeceptjs');

Feature('Run Test');

Scenario('Docker', async ({ I }) => {
  // Navigate to the page containing the CloneRepository component
  I.amOnPage('http://localhost:3000');
  I.wait(3);
  I.see('WebTestHub');
  I.click('Get Started');
  I.click('Create Docker file ');
  I.click('Confirm');
  I.wait(20);
});
