// Import CodeceptJS and the necessary modules
const { I } = require('codeceptjs');

Feature('search');

Scenario('test something on the homepage', async ({ I }) => {
  // Navigate to the homepage
  I.amOnPage('http://localhost:3000');

  // Wait for a certain amount of time (e.g., 2 seconds)
  I.wait(2);

  // Assert that the "Web Test Hub" title is present on the page
  I.see('Web Test Hub');

  // Perform additional assertions or interactions as needed

  // Example: Assert that the Selenium section is present on the page
  I.see('Selenium');
  I.see('Selenium is a powerful open-source framework');

  // Click the "Get Started" button
  I.click('Get Started');

  // Wait for a certain amount of time (e.g., 3 seconds) after clicking the button
  I.wait(3);

});
